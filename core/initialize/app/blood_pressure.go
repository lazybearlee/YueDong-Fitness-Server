package app

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"go.uber.org/zap"
)

type BloodPressureInitializer struct{}

func (b *BloodPressureInitializer) Name() string {
	return appmodel.BloodPressure{}.TableName()
}

func (b *BloodPressureInitializer) MigrateTable() error {
	err := global.FitnessDb.AutoMigrate(&appmodel.BloodPressure{})
	if err != nil {
		global.FitnessLog.Error("blood pressure初始化失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("blood pressure数据库迁移成功")
		return nil
	}
}

func (b *BloodPressureInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&appmodel.BloodPressure{})
}

func (b *BloodPressureInitializer) InitializeData() error {
	// 插入血压数据
	//bloodPressures := []appmodel.BloodPressure{
	//	{
	//		UID: 3,
	//		Date: time.Now(),
	//	},
	//}
	return nil
}

func (b *BloodPressureInitializer) DataInitialized() bool {
	// 无需初始化数据
	return true
}
