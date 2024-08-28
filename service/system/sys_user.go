package sysservice

import (
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"gorm.io/gorm"
	"time"
)

// UserService
// @description: 用户服务
// 提供登陆、注册、获取用户信息、修改用户信息、删除用户、批量删除用户、修改密码、重置密码等功能
type UserService struct{}

// UserRegister
// @description: 用户注册
// @param: user model.SysUser
// @return: model.SysUser, error
func (userService *UserService) UserRegister(user sysmodel.SysUser) (sysmodel.SysUser, error) {
	var u sysmodel.SysUser
	// 查询用户名是否注册
	err := global.FitnessDb.Where("username = ?", user.Username).First(&u).Error
	switch err {
	case nil:
		return u, errors.New("用户名已注册")
	case gorm.ErrRecordNotFound:
		break
	default:
		return u, err
	}
	// 附加uuid 密码hash加密 注册
	user.Password = utils.CryptWithBcrypt(user.Password)
	user.UUID = uuid.Must(uuid.NewV4())
	err = global.FitnessDb.Create(&user).Error // 创建用户
	return user, err
}

// UserLogin
// @description: 用户登录
// @param: user model.SysUser
// @return: *model.SysUser, error
func (userService *UserService) UserLogin(user sysmodel.SysUser) (*sysmodel.SysUser, error) {
	// 首先判断数据库是否初始化
	if nil == global.FitnessDb {
		return nil, errors.New("mysql not init")
	}
	var u sysmodel.SysUser
	// 查询用户，预加载权限
	err := global.FitnessDb.Where("username = ?", user.Username).Preload("Authorities").Preload("Authority").First(&u).Error
	if err == nil {
		// 检查密码
		if ok := utils.CryptCheckWithBcrypt(user.Password, u.Password); !ok {
			return nil, errors.New("密码错误")
		}
		// TODO: 查询用户权限，设置默认路由
	}
	return &u, err
}

// UserChangePassword
// @description: 修改用户密码
// @param: user model.SysUser, newPassword string
// @return: *model.SysUser, error
func (userService *UserService) UserChangePassword(user sysmodel.SysUser, newPassword string) (*sysmodel.SysUser, error) {
	// 首先判断用户是否存在以及密码是否正确
	var u sysmodel.SysUser
	if err := global.FitnessDb.Where("id = ?", user.ID).First(&u).Error; err != nil {
		return nil, err
	}
	if ok := utils.CryptCheckWithBcrypt(user.Password, u.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	// 修改密码
	u.Password = utils.CryptWithBcrypt(newPassword)
	err := global.FitnessDb.Save(&u).Error
	return &u, err
}

// UserResetPassword
// @description: 重置用户密码(默认123456)
// @param: userID int
// @return: error
func (userService *UserService) UserResetPassword(userId int) error {
	// 修改密码
	return global.FitnessDb.Model(&sysmodel.SysUser{}).Where("id = ?", userId).Update("password", utils.CryptWithBcrypt("123456")).Error
}

// UserSetAuthority
// TODO: 用户设置权限

// UserDelete
// @description: 删除用户
// @param: userID int
// @return: error
func (userService *UserService) UserDelete(userID int) error {
	return global.FitnessDb.Transaction(func(tx *gorm.DB) error {
		// 删除用户
		if err := tx.Delete(&sysmodel.SysUser{}, userID).Error; err != nil {
			return err
		}
		// 删除用户角色
		if err := tx.Delete(&[]sysmodel.SysUserAuthority{}, "user_id = ?", userID).Error; err != nil {
			return err
		}
		return nil
	})
}

// UserSetInfo
// @description: 设置用户信息(仅昵称、电话、邮箱等)
// @param: user model.SysUser
// @return: error
func (userService *UserService) UserSetInfo(user sysmodel.SysUser) error {
	return global.FitnessDb.Model(&sysmodel.SysUser{}).
		Select("updated_at", "nick_name", "gender", "phone", "email", "enable").
		Where("id=?", user.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  user.NickName,
			"gender":     user.Gender,
			"phone":      user.Phone,
			"email":      user.Email,
			"enable":     user.Enable,
		}).Error
}

// UserSetAvatar
// @description: 设置用户头像
// @param: url string, uid uint
// @return: error
func (userService *UserService) UserSetAvatar(url string, uid uint) error {
	return global.FitnessDb.Model(&sysmodel.SysUser{}).Where("id = ?", uid).Update("header_img", url).Error
}

// UserGetInfoWithUUID
// @description: 通过UUID获取用户信息
// @param: uuid uuid.UUID
// @return: *model.SysUser, error
func (userService *UserService) UserGetInfoWithUUID(uuid uuid.UUID) (*sysmodel.SysUser, error) {
	var u sysmodel.SysUser
	err := global.FitnessDb.Where("uuid = ?", uuid).First(&u).Error
	if err != nil {
		return &u, err
	}
	// TODO: 查询用户权限，设置默认路由
	return &u, err
}

// UserGetInfoWithID
// @description: 通过ID获取用户信息
// @param: id int
// @return: *model.SysUser, error
func (userService *UserService) UserGetInfoWithID(id uint) (*sysmodel.SysUser, error) {
	var u sysmodel.SysUser
	err := global.FitnessDb.Where("id = ?", id).First(&u).Error
	if err != nil {
		return &u, err
	}
	// TODO: 查询用户权限，设置默认路由
	return &u, err
}
