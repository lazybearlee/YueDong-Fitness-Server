package system

import (
	"errors"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lazybearlee/yuedong-fitness/core"
	"github.com/lazybearlee/yuedong-fitness/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CasbinInitializer struct{}

// Name 初始化器名称
func (s *CasbinInitializer) Name() string {
	return adapter.CasbinRule{}.TableName()
}

// MigrateTable 初始化表
func (s *CasbinInitializer) MigrateTable() error {
	// 如果数据库连接未建立，报错
	if global.FITNESS_DB == nil {
		global.FITNESS_LOG.Error("casbin初始化失败，db未初始化")
		return errors.New("casbin初始化失败，db未初始化")
	}
	// 自动迁移casbin规则
	err := global.FITNESS_DB.AutoMigrate(&adapter.CasbinRule{})
	if err != nil {
		global.FITNESS_LOG.Error("casbin初始化失败", zap.Error(err))
		return err
	} else {
		global.FITNESS_LOG.Info("casbin数据库迁移成功")
		return nil
	}
}

// TableCreated 表是否已创建
func (s *CasbinInitializer) TableCreated() bool {
	if global.FITNESS_DB == nil {
		global.FITNESS_LOG.Error("casbin查询表是否建立失败，db未初始化")
		return false
	}
	return global.FITNESS_DB.Migrator().HasTable(&adapter.CasbinRule{})
}

// InitializeData 初始化数据
func (s *CasbinInitializer) InitializeData() error {
	if global.FITNESS_DB == nil {
		global.FITNESS_LOG.Error("casbin初始化数据失败，db未初始化")
		return errors.New("casbin初始化数据失败，db未初始化")
	}
	// 初始化casbin规则
	rules := []adapter.CasbinRule{
		// 超级管理员注册管理员角色
		{Ptype: "p", V0: core.AdminSuperStr, V1: "/user/admin_register", V2: "POST"},
	}
	// 创建casbin规则
	err := global.FITNESS_DB.Create(&rules).Error
	if err != nil {
		global.FITNESS_LOG.Error("casbin初始化数据失败", zap.Error(err))
		return err
	} else {
		global.FITNESS_LOG.Info("casbin初始化数据成功")
		return nil
	}
}

// DataInitialized 数据是否已插入
func (s *CasbinInitializer) DataInitialized() bool {
	// 检查其中一个规则是否存在
	if global.FITNESS_DB == nil {
		global.FITNESS_LOG.Error("casbin查询数据是否初始化失败，db未初始化")
		return false
	}
	if errors.Is(global.FITNESS_DB.Where(adapter.CasbinRule{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
