package sysservice

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lazybearlee/yuedong-fitness/global"
	sysrequest "github.com/lazybearlee/yuedong-fitness/model/system/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type CasbinService struct{}

var (
	CasbinServiceApp     = new(CasbinService)
	once                 sync.Once
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
)

// Casbin 初始化casbin，采用类似于单例的方式，在第一次调用时初始化
// 采用casbin的RBAC模型，将权限存储在数据库中
func (casbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.FitnessDb)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}

/** ------------------- 以下方法不要求传入gorm.DB，使用syncedCachedEnforcer ------------------- **/

// ClearCasbin 清除角色权限
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// UpdateCasbin 更新casbin策略
func (casbinService *CasbinService) UpdateCasbin(req *sysrequest.CasbinInReceive) error {
	// 首先清除策略
	authorityID := strconv.Itoa(int(req.AuthorityId))
	casbinService.ClearCasbin(0, authorityID)
	var rules [][]string
	deduplicateMap := make(map[string]struct{}) // 用于去重
	for _, v := range req.CasbinInfos {
		// 去重
		if _, ok := deduplicateMap[v.Path+v.Method]; ok {
			continue
		}
		deduplicateMap[v.Path+v.Method] = struct{}{}
		rules = append(rules, []string{authorityID, v.Path, v.Method})
	}
	success, _ := casbinService.Casbin().AddPolicies(rules)
	if !success {
		return errors.New("存在重复数据，请重新编辑")
	} else {
		return nil
	}
}

// UpdateCasbinApi 更新api策略
func (casbinService *CasbinService) UpdateCasbinApi(old *sysrequest.CasbinApiInfo, new *sysrequest.CasbinApiInfo) error {
	err := global.FitnessDb.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", old.Path, old.Method).Updates(map[string]interface{}{
		"v1": new.Path,
		"v2": new.Method,
	}).Error
	e := casbinService.Casbin()
	err = e.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

// GetPolicyPathByAuthorityId 获取权限列表
func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId uint) (pathMaps []sysrequest.CasbinApiInfo) {
	e := casbinService.Casbin()
	id := strconv.Itoa(int(authorityId))
	policy := e.GetFilteredPolicy(0, id) // 获取策略
	for _, v := range policy {
		pathMaps = append(pathMaps, sysrequest.CasbinApiInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return
}

// FreshCasbinPolicy 更新casbin策略
func (casbinService *CasbinService) FreshCasbinPolicy() error {
	return casbinService.Casbin().LoadPolicy()
}

/** ------------------- 以下方法要求传入gorm.DB，作为事务处理的一部分 ------------------- **/

// RemovePolicyByAuthorityId 根据角色ID删除权限
func (casbinService *CasbinService) RemovePolicyByAuthorityId(db *gorm.DB, authorityId string) error {
	// 删除权限
	return db.Delete(&gormadapter.CasbinRule{}, "v0 = ?", authorityId).Error
}

// AddPolicies 添加权限
func (casbinService *CasbinService) AddPolicies(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
}

// SyncPolicy 同步权限
func (casbinService *CasbinService) SyncPolicy(db *gorm.DB, authorityId string, rules [][]string) error {
	err := casbinService.RemovePolicyByAuthorityId(db, authorityId)
	if err != nil {
		return err
	}
	return casbinService.AddPolicies(db, rules)
}
