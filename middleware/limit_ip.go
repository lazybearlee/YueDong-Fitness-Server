package middleware

import (
	"context"
	"errors"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type LimitConfig struct {
	// GenerationKey 根据业务生成key 下面CheckOrMark查询生成
	GenerationKey func(c *gin.Context) string
	// 检查函数,用户可修改具体逻辑,更加灵活
	CheckOrMark func(key string, expire int, limit int) error
	// Expire key 过期时间
	Expire int
	// Limit 周期时间
	Limit int
}

func (l LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			response.ErrorWithMessage(err.Error(), c)
			return
		} else {
			c.Next()
		}
	}
}

// DefaultGenerationKey 默认生成key
func DefaultGenerationKey(c *gin.Context) string {
	return "FitnessLimit" + c.ClientIP()
}

func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
	// 判断是否开启redis
	if global.FitnessRedis == nil {
		return err
	}
	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
		global.FitnessLog.Error("limit", zap.Error(err))
	}
	return err
}

func CheckOrMarkUsingLocalCache(key string, expire int, limit int) error {
	return SetLimitWithTimeUsingLocalCache(key, limit, time.Duration(expire)*time.Second)
}

func DefaultLimit() gin.HandlerFunc {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   DefaultCheckOrMark,
		Expire:        global.FitnessConfig.System.LimitTimeIP,
		Limit:         global.FitnessConfig.System.LimitCountIP,
	}.LimitWithTime()
}

func LimitWithTimeUsingLocalCache() gin.HandlerFunc {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   CheckOrMarkUsingLocalCache,
		Expire:        global.FitnessConfig.System.LimitTimeIP,
		Limit:         global.FitnessConfig.System.LimitCountIP,
	}.LimitWithTime()
}

// SetLimitWithTimeUsingLocalCache 设置访问次数
func SetLimitWithTimeUsingLocalCache(key string, limit int, expiration time.Duration) error {
	count, ok := global.FitnessCache.Get(key)
	if count == 0 || !ok {
		global.FitnessCache.Set(key, 1, expiration)
	} else {
		if times, ok := count.(int); !ok {
			return errors.New("类型转换失败")
		} else {
			if times >= limit {
				return errors.New("请求太过频繁，请稍后再试")
			} else {
				global.FitnessCache.Increment(key, 1)
			}
		}
	}
	return nil
}

// SetLimitWithTime 设置访问次数
func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	count, err := global.FitnessRedis.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		pipe := global.FitnessRedis.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err = pipe.Exec(context.Background())
		return err
	} else {
		// 次数
		if times, err := global.FitnessRedis.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := global.FitnessRedis.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("请求太过频繁，请稍后再试")
				} else {
					return errors.New("请求太过频繁, 请 " + t.String() + " 秒后尝试")
				}
			} else {
				return global.FitnessRedis.Incr(context.Background(), key).Err()
			}
		}
	}
}
