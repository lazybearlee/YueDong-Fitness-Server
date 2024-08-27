package app

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"go.uber.org/zap"
)

type PlanStageInitializer struct{}

func (s *PlanStageInitializer) Name() string {
	return appmodel.PlanStage{}.TableName()
}

func (s *PlanStageInitializer) MigrateTable() error {
	err := global.FitnessDb.AutoMigrate(&appmodel.PlanStage{})
	if err != nil {
		global.FitnessLog.Error("exercise plan初始化失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("exercise plan数据库迁移成功")
		return nil
	}
}

func (s *PlanStageInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&appmodel.PlanStage{})
}

func (s *PlanStageInitializer) InitializeData() error {
	// 无需初始化数据
	return nil
}

func (s *PlanStageInitializer) DataInitialized() bool {
	// 无需初始化数据
	return true
}
