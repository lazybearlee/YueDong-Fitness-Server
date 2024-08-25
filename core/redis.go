package core

import (
	"context"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"os"
)

func RedisInit() {
	// 初始化Redis
	// 判断是否使用Redis
	if global.FitnessConfig.System.UseRedis || global.FitnessConfig.System.UseMultipoint {
		// 使用Redis
		config := global.FitnessConfig.Redis
		client := redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
		})
		// 测试Redis连接
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			global.FitnessLog.Error("Redis connect ping failed", zap.Error(err))
			os.Exit(0)
		} else {
			global.FitnessLog.Info("Redis connect ping success")
			global.FitnessRedis = client
		}
	} else {
		// 不使用Redis
		global.FitnessLog.Info("Redis is not used")
	}
}
