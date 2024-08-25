package sysservice

import (
	"errors"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/common/request"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	sysrequest "github.com/lazybearlee/yuedong-fitness/model/system/request"
	"gorm.io/gorm"
	"strconv"
)

var (
	ErrAuthorityExist  = errors.New("该角色已存在")
	ErrorAuthorityGet  = errors.New("获取角色失败")
	ErrDeleteUserUsing = errors.New("删除失败，有用户正在使用该角色")
)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

// CreateAuthority 创建一个权限
func (authorityService *AuthorityService) CreateAuthority(authority sysmodel.SysAuthority) (auth sysmodel.SysAuthority, err error) {
	// 首先查询是否存在相同的角色id
	if err = global.FitnessDb.Where("authority_id = ?", authority.AuthorityId).First(&sysmodel.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return authority, ErrAuthorityExist
	}
	// 开启事务
	// 做以下几件事，创建权限、设置casbin权限
	e := global.FitnessDb.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&authority).Error; err != nil {
			return err
		}
		// 设置casbin权限
		casbinInfos := getDefaultCasbinInfo()
		authorityId := strconv.Itoa(int(authority.AuthorityId))
		var rules [][]string
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules)
	})
	return authority, e
}

// UpdateAuthority 更新一个权限
func (authorityService *AuthorityService) UpdateAuthority(authority sysmodel.SysAuthority) (auth sysmodel.SysAuthority, err error) {
	// 先查询是否存在该权限
	var old sysmodel.SysAuthority
	if err = global.FitnessDb.Where("authority_id = ?", authority.AuthorityId).First(&old).Error; err != nil {
		return authority, ErrorAuthorityGet
	}
	// 更新，这里old是查询到的权限，由于其自带id，所以不需要再次传入id
	err = global.FitnessDb.Model(&old).Updates(&auth).Error
	return auth, err
}

// DeleteAuthority 删除一个权限
func (authorityService *AuthorityService) DeleteAuthority(auth *sysmodel.SysAuthority) (err error) {
	// 首先检查是否存在该权限
	if err = global.FitnessDb.Debug().Preload("Users").First(&auth).Error; err != nil {
		return ErrorAuthorityGet
	}
	// 如果有用户使用该权限，不允许删除
	if len(auth.Users) > 0 {
		return ErrDeleteUserUsing
	}
	if !errors.Is(global.FitnessDb.Where("authority_id = ?", auth.AuthorityId).First(&sysmodel.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.FitnessDb.Where("parent_id = ?", auth.AuthorityId).First(&sysmodel.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}

	// 开启事务，需要依次完成以下操作
	// 1. 删除权限
	// 2. 删除相关DataAuthority
	// 3. 删除用户与权限关联
	// 4. 删除Casbin权限
	return global.FitnessDb.Transaction(func(tx *gorm.DB) (err error) {
		// 删除权限，不过是软删除，也就是不会真正删除，只是将deleted_at字段设置为当前时间
		if err = tx.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth).Error; err != nil {
			return
		}
		// 删除DataAuthority
		if len(auth.DataAuthorityId) > 0 {
			if err = tx.Model(auth).Association("DataAuthorityId").Delete(auth.DataAuthorityId); err != nil {
				return
			}
		}
		// 删除用户与权限关联
		if err = tx.Delete(&sysmodel.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error; err != nil {
			return
		}
		// 删除Casbin权限
		if err = CasbinServiceApp.RemovePolicyByAuthorityId(tx, strconv.Itoa(int(auth.AuthorityId))); err != nil {
			return
		}
		return
	})
}

// getDefaultCasbinInfo 获取默认的casbin信息
func getDefaultCasbinInfo() []sysrequest.CasbinApiInfo {
	return []sysrequest.CasbinApiInfo{
		/** base */
		{Path: "/base/login", Method: "POST"},
		/** user */
		{Path: "/user/changePassword", Method: "POST"},
		{Path: "/user/getUserInfo", Method: "POST"},
		{Path: "/user/setUserInfo", Method: "POST"},
	}
}

// GetChildrenAuthority 获取子权限
func (authorityService *AuthorityService) findChildrenAuthority(authority *sysmodel.SysAuthority) (err error) {
	err = global.FitnessDb.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

// GetAuthorityInfoList 获取权限列表
func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.FitnessDb.Model(&sysmodel.SysAuthority{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var authority []sysmodel.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authority).Error
	for k := range authority {
		err = authorityService.findChildrenAuthority(&authority[k])
	}
	return authority, total, err
}

// GetAuthorityInfoByID 获取权限信息
func (authorityService *AuthorityService) GetAuthorityInfoByID(auth sysmodel.SysAuthority) (sa sysmodel.SysAuthority, err error) {
	err = global.FitnessDb.Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return
}

// SetDataAuthority 设置数据权限
func (authorityService *AuthorityService) SetDataAuthority(auth sysmodel.SysAuthority) error {
	var s sysmodel.SysAuthority
	global.FitnessDb.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.FitnessDb.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}
