package global

import (
	"github.com/gin-gonic/gin"
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
	FitnessDb     *gorm.DB
	FitnessRedis  redis.UniversalClient
	FitnessConfig config.Server
	FitnessViper  *viper.Viper
	FitnessLog    *zap.Logger
	//FITNESS_TIMER  timer.Timer = timer.NewTimerTask()
	FitnessCc      = &singleflight.Group{} // 并发控制
	FitnessCache   local_cache.Cache
	FitnessRouters gin.RoutesInfo
	LOCK           sync.RWMutex
)
