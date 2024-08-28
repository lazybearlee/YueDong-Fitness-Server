package system

import (
	"errors"
	adapter "github.com/casbin/gorm-adapter/v3"
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
	// 自动迁移casbin规则
	err := global.FitnessDb.AutoMigrate(&adapter.CasbinRule{})
	if err != nil {
		global.FitnessLog.Error("casbin初始化失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("casbin数据库迁移成功")
		return nil
	}
}

// TableCreated 表是否已创建
func (s *CasbinInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&adapter.CasbinRule{})
}

// InitializeData 初始化数据
func (s *CasbinInitializer) InitializeData() error {
	// 初始化casbin规则
	rules := []adapter.CasbinRule{
		// 超级管理员注册管理员角色
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/user/get_user_info", V2: "GET"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/user/update_user_info", V2: "POST"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/user/update_user_password", V2: "POST"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/record/insert_exercise_record", V2: "POST"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/record/update_exercise_record", V2: "PUT"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/record/delete_exercise_record", V2: "DELETE"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/record/delete_exercise_records", V2: "DELETE"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/record/get_exercise_record", V2: "GET"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/record/get_exercise_record_list", V2: "GET"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/record/get_all_exercise_record_of_user", V2: "GET"},
		// rank
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/rank/get_rank_list", V2: "GET"},
		// health_status
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/health_status/get_health_status", V2: "GET"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/health_status/put_health_status", V2: "PUT"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/health_status/get_health_status_list", V2: "GET"},
		// exercise_plan
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/plan/get_all_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/plan/get_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/plan/get_started_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/plan/get_uncompleted_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/plan/create_exercise_plan", V2: "POST"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/plan/update_exercise_plan", V2: "PUT"},
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/plan/delete_exercise_plans", V2: "DELETE"},
		// admin
		{Ptype: "p", V0: global.AdminSuperStr, V1: "/admin/user/register", V2: "POST"},
		// 管理员注册用户角色
		{Ptype: "p", V0: global.AdminUserStr, V1: "/user/get_user_info", V2: "GET"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/user/update_user_info", V2: "POST"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/user/update_user_password", V2: "POST"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/record/insert_exercise_record", V2: "POST"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/record/update_exercise_record", V2: "PUT"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/record/delete_exercise_record", V2: "DELETE"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/record/delete_exercise_records", V2: "DELETE"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/record/get_exercise_record", V2: "GET"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/record/get_exercise_record_list", V2: "GET"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/record/get_all_exercise_record_of_user", V2: "GET"},
		// rank
		{Ptype: "p", V0: global.AdminUserStr, V1: "/rank/get_rank_list", V2: "GET"},
		// health_status
		{Ptype: "p", V0: global.AdminUserStr, V1: "/health_status/get_health_status", V2: "GET"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/health_status/put_health_status", V2: "PUT"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/health_status/get_health_status_list", V2: "GET"},
		// exercise_plan
		{Ptype: "p", V0: global.AdminUserStr, V1: "/plan/get_all_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/plan/get_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/plan/get_started_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/plan/get_uncompleted_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/plan/create_exercise_plan", V2: "POST"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/plan/update_exercise_plan", V2: "PUT"},
		{Ptype: "p", V0: global.AdminUserStr, V1: "/plan/delete_exercise_plans", V2: "DELETE"},
		// 普通用户
		// user路由
		{Ptype: "p", V0: global.CommonUserStr, V1: "/user/get_user_info", V2: "GET"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/user/update_user_info", V2: "POST"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/user/update_user_password", V2: "POST"},
		// record
		{Ptype: "p", V0: global.CommonUserStr, V1: "/record/insert_exercise_record", V2: "POST"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/record/update_exercise_record", V2: "PUT"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/record/delete_exercise_record", V2: "DELETE"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/record/delete_exercise_records", V2: "DELETE"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/record/get_exercise_record", V2: "GET"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/record/get_exercise_record_list", V2: "GET"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/record/get_all_exercise_record_of_user", V2: "GET"},
		// rank
		{Ptype: "p", V0: global.CommonUserStr, V1: "/rank/get_rank_list", V2: "GET"},
		// health_status
		{Ptype: "p", V0: global.CommonUserStr, V1: "/health_status/get_health_status", V2: "GET"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/health_status/put_health_status", V2: "PUT"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/health_status/get_health_status_list", V2: "GET"},
		// exercise_plan
		{Ptype: "p", V0: global.CommonUserStr, V1: "/plan/get_all_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/plan/get_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/plan/get_started_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/plan/get_uncompleted_exercise_plans", V2: "GET"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/plan/create_exercise_plan", V2: "POST"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/plan/update_exercise_plan", V2: "PUT"},
		{Ptype: "p", V0: global.CommonUserStr, V1: "/plan/delete_exercise_plans", V2: "DELETE"},
	}
	// 创建casbin规则
	err := global.FitnessDb.Create(&rules).Error
	if err != nil {
		global.FitnessLog.Error("casbin初始化数据失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("casbin初始化数据成功")
		return nil
	}
}

// DataInitialized 数据是否已插入
func (s *CasbinInitializer) DataInitialized() bool {
	// 检查其中一个规则是否存在
	if errors.Is(global.FitnessDb.Where(adapter.CasbinRule{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
