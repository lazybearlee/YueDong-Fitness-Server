package sysservice

import (
	"context"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"go.uber.org/zap"
)

// LoadJwtBlackList 载入jwt黑名单缓存
func LoadJwtBlackList() {
	var jwtBlackList []string
	err := global.FitnessDb.Model(&sysmodel.JwtBlacklist{}).Select("jwt").Find(&jwtBlackList).Error
	if err != nil {
		global.FitnessLog.Error("LoadJwtBlackList error", zap.Error(err))
		return
	}
	// 将jwt黑名单载入缓存
	for _, v := range jwtBlackList {
		global.FitnessCache.SetDefault(v, struct{}{})
	}
}

// JwtService
// JWT service interface 提供了一些JWT相关的方法\n
// 1. SetInBlacklist 拉黑jwt（需要注意缓存与数据库的同步问题）
// 2. IsInBlacklist 判断JWT是否在黑名单内部
// 3. GetRedisJWT 从redis取jwt
// 4. SetRedisJWT jwt存入redis并设置过期时间
type JwtService struct{}

var JwtServiceApp = new(JwtService)

// SetInBlacklist 拉黑jwt
func (jwtService *JwtService) SetInBlacklist(blacklist sysmodel.JwtBlacklist) (err error) {
	err = global.FitnessDb.Create(&blacklist).Error
	if err != nil {
		return
	}
	global.FitnessCache.SetDefault(blacklist.Jwt, struct{}{})
	return
}

// IsInBlacklist 判断JWT是否在黑名单内部
func (jwtService *JwtService) IsInBlacklist(jwt string) bool {
	_, ok := global.FitnessCache.Get(jwt)
	return ok
	// Do we need to query the database?
}

// GetRedisJWT 从redis取jwt
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.FitnessRedis.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// SetRedisJWT jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt, userName string) (err error) {
	// 首先解析jwt的过期时间
	dr, err := utils.ParseDuration(global.FitnessConfig.JWT.ExpiresTime)
	if err != nil {
		global.FitnessLog.Error("jwt parse duration failed", zap.Error(err))
		return err
	}
	// 将jwt存入redis并设置过期时间
	timer := dr
	err = global.FitnessRedis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
