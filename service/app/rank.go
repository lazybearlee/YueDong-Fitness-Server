package appservice

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"time"
)

type RankService struct{}

// GetStepRank is a method to get step rank
// @Description: 获取某一天的步数排行榜
// @Return: []appmodel.UserStepRank
func (r *RankService) GetStepRank(date time.Time, limit int) ([]appmodel.UserStepRank, error) {
	var stepRanks []appmodel.UserStepRank

	// 查询 HealthStatus 表，按步数降序排列，并限制结果数量
	if err := global.FitnessDb.Table("health_statuses").
		Select("health_statuses.uid as user_id, sys_users.header_img as header_img, sys_users.nick_name as nickname, health_statuses.steps_count as step").
		Joins("left join sys_users on sys_users.id = health_statuses.uid").
		Where("DATE(health_statuses.date) = ?", date.Format("2006-01-02")).
		Order("health_statuses.steps_count desc").
		Limit(limit).
		Scan(&stepRanks).Error; err != nil {
		return nil, err
	}

	// 添加排名信息
	for i := range stepRanks {
		stepRanks[i].Rank = uint(i + 1)
	}

	return stepRanks, nil
}

// GetDistanceRank is a method to get distance rank
// @Description: 获取今日的距离排行榜
// @Return: []appmodel.UserDistanceRank
func (r *RankService) GetDistanceRank(limit int) ([]appmodel.UserDistanceRank, error) {
	var distanceRanks []appmodel.UserDistanceRank

	// 查询 HealthStatus 表，按距离降序排列，并限制结果数量
	if err := global.FitnessDb.Table("health_statuses").
		Select("health_statuses.uid as user_id, sys_users.header_img as header_img, sys_users.nick_name as nickname, health_statuses.distance as distance").
		Joins("left join sys_users on sys_users.id = health_statuses.uid").
		Where("health_statuses.date >= ?", time.Now().Add(-24*time.Hour)).
		Where("health_statuses.date <= ?", time.Now()).
		Order("health_statuses.distance desc").
		Limit(limit).
		Scan(&distanceRanks).Error; err != nil {
		return nil, err
	}

	// 添加排名信息
	for i := range distanceRanks {
		distanceRanks[i].Rank = uint(i + 1)
	}

	return distanceRanks, nil
}
