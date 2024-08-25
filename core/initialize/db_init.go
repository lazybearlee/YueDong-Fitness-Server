package initialize

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	"go.uber.org/zap"
)

// InitDB 初始化数据库
func InitDB() {
	// 首先注册所有的初始化器
	LoadInitializers()
	// 初始化表
	InitTables()
	// 初始化数据
	InitData()
}

// InitTables 初始化表
func InitTables() {
	// 对每个初始化器进行表初始化
	for _, initializer := range initializers {
		if !initializer.TableCreated() {
			err := initializer.MigrateTable()
			if err != nil {
				global.FitnessLog.Error("初始化表失败", zap.Error(err))
				panic(err)
			}
		}
	}
}

// InitData 初始化数据
func InitData() {
	// 如果配置中要求初始化数据，则对每个初始化器进行数据初始化
	if global.FitnessConfig.System.MysqlInitData {
		for _, initializer := range initializers {
			if !initializer.DataInitialized() {
				global.FitnessLog.Info("初始化数据", zap.String("name", initializer.Name()))
				err := initializer.InitializeData()
				if err != nil {
					global.FitnessLog.Error("初始化数据失败", zap.Error(err))
					panic(err)
				}
			}
		}
	}
}
