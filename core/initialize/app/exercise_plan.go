package app

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"go.uber.org/zap"
)

type ExercisePlanInitializer struct{}

func (s *ExercisePlanInitializer) Name() string {
	return appmodel.ExercisePlan{}.TableName()
}

func (s *ExercisePlanInitializer) MigrateTable() error {
	err := global.FitnessDb.AutoMigrate(&appmodel.ExercisePlan{})
	if err != nil {
		global.FitnessLog.Error("exercise plan初始化失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("exercise plan数据库迁移成功")
		return nil
	}
}

func (s *ExercisePlanInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&appmodel.ExercisePlan{})
}

func (s *ExercisePlanInitializer) InitializeData() error {
	// 无需初始化数据
	return nil
}

func (s *ExercisePlanInitializer) DataInitialized() bool {
	// 无需初始化数据
	return true
}
