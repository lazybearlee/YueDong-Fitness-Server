package core

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysservice "github.com/lazybearlee/yuedong-fitness/service/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"os"
)

func JWTInit() {
	// parse jwt duration
	dr, err := utils.ParseDuration(global.FitnessConfig.JWT.ExpiresTime)
	if err != nil {
		global.FitnessLog.Error("jwt parse duration failed", zap.Error(err))
		os.Exit(0)
	}
	// parse jwt buffer time
	_, err = utils.ParseDuration(global.FitnessConfig.JWT.BufferTime)
	if err != nil {
		global.FitnessLog.Error("jwt parse duration failed", zap.Error(err))
		os.Exit(0)
	}
	// set jwt cache expire time
	global.FitnessCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
	global.FitnessLog.Info("jwt parse duration success")
	// load jwt data from database
	if global.FitnessDb != nil {
		// load jwt data from database
		sysservice.LoadJwtBlackList()
	}
}
