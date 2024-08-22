package core

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/service/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"os"
)

func JWTInit() {
	// parse jwt duration
	dr, err := utils.ParseDuration(global.FITNESS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		global.FITNESS_LOG.Error("jwt parse duration failed", zap.Error(err))
		os.Exit(0)
	}
	// parse jwt buffer time
	_, err = utils.ParseDuration(global.FITNESS_CONFIG.JWT.BufferTime)
	if err != nil {
		global.FITNESS_LOG.Error("jwt parse duration failed", zap.Error(err))
		os.Exit(0)
	}
	// set jwt cache expire time
	global.FITNESS_CACHE = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
	global.FITNESS_LOG.Info("jwt parse duration success")
	// load jwt data from database
	if global.FITNESS_DB != nil {
		// load jwt data from database
		system.LoadJwtBlackList()
	}
}
