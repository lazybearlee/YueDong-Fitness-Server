package app

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"go.uber.org/zap"
)

type ExerciseRecordInitializer struct{}

func (s *ExerciseRecordInitializer) Name() string {
	return appmodel.ExerciseRecord{}.TableName()
}

func (s *ExerciseRecordInitializer) MigrateTable() error {
	err := global.FitnessDb.AutoMigrate(&appmodel.ExerciseRecord{})
	if err != nil {
		global.FitnessLog.Error("exercise plan初始化失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("exercise plan数据库迁移成功")
		return nil
	}
}

func (s *ExerciseRecordInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&appmodel.ExerciseRecord{})
}

func (s *ExerciseRecordInitializer) InitializeData() error {
	// 无需初始化数据
	return nil
}

func (s *ExerciseRecordInitializer) DataInitialized() bool {
	// 无需初始化数据
	return true
}
