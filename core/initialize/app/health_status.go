package app

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"go.uber.org/zap"
)

type HealthStatusInitializer struct{}

func (s *HealthStatusInitializer) Name() string {
	return appmodel.HealthStatus{}.TableName()
}

func (s *HealthStatusInitializer) MigrateTable() error {
	err := global.FitnessDb.AutoMigrate(&appmodel.HealthStatus{})
	if err != nil {
		global.FitnessLog.Error("exercise plan初始化失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("exercise plan数据库迁移成功")
		return nil
	}
}

func (s *HealthStatusInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&appmodel.HealthStatus{})
}

func (s *HealthStatusInitializer) InitializeData() error {
	// 无需初始化数据
	return nil
}

func (s *HealthStatusInitializer) DataInitialized() bool {
	// 无需初始化数据
	return true
}
