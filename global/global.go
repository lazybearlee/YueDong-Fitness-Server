package global

import (
	"github.com/lazybearlee/yuedong-fitness/config"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"
)

var (
	FITNESS_DB     *gorm.DB
	FITNESS_REDIS  redis.UniversalClient
	FITNESS_CONFIG config.Server
	FITNESS_VIPER  *viper.Viper
	FITNESS_LOG    *zap.Logger
	//FITNESS_TIMER  timer.Timer = timer.NewTimerTask()
	FITNESS_CC    = &singleflight.Group{} // 并发控制
	FITNESS_CACHE local_cache.Cache
	LOCK          sync.RWMutex
)
