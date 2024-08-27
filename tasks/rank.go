package tasks

import (
	"errors"
	"github.com/lazybearlee/yuedong-fitness/global"
	appservice "github.com/lazybearlee/yuedong-fitness/service/app"
	"gorm.io/gorm"
	"time"
)

var rankService = new(appservice.RankService)

// RankOfStep
// @Description: 获取用户步数排名
func RankOfStep(db *gorm.DB) error {
	if db == nil {
		return errors.New("mysql Cannot be empty")
	}
	// 当前时间的前一天
	date := time.Now().AddDate(0, 0, -1)
	rank, err := rankService.GetStepRank(date, 50)
	if err != nil {
		return err
	}
	// 更新用户步数排名到Cache，以便后续查询，过期时间为一天
	global.FitnessCache.Set(global.StepRankYesterday, rank, time.Hour*24)
	return nil
}
