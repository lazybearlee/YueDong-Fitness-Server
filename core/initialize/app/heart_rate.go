package app

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"go.uber.org/zap"
)

type HeartRateInitializer struct{}

func (h *HeartRateInitializer) Name() string {
	return appmodel.HeartRate{}.TableName()
}

func (h *HeartRateInitializer) MigrateTable() error {
	err := global.FitnessDb.AutoMigrate(&appmodel.HeartRate{})
	if err != nil {
		global.FitnessLog.Error("heart rate初始化失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("heart rate数据库迁移成功")
		return nil
	}
}

func (h *HeartRateInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&appmodel.HeartRate{})
}

func (h *HeartRateInitializer) InitializeData() error {
	// 插入心率数据
	//heartRates := []appmodel.HeartRate{
	//	{
	//		UID: 3,
	//		Date: time.Now(),
	//	},
	//}
	return nil
}

func (h *HeartRateInitializer) DataInitialized() bool {
	// 无需初始化数据
	return true
}
