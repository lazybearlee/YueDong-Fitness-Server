package appservice

import (
	"errors"
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	apprequest "github.com/lazybearlee/yuedong-fitness/model/app/request"
	"gorm.io/gorm"
)

var (
	healthOrderMap = map[string]bool{
		"id":       true,
		"uid":      true,
		"date":     true,
		"height":   true,
		"weight":   true,
		"bmi":      true,
		"distance": true,
		"steps":    true,
	}
)

// HealthStatusService is a service for health status
// @Description: 用户健康状态服务
// 提供用户健康状态的增删改查
type HealthStatusService struct{}

// InsertHealthStatus is a method to insert health status
// @Description: 插入用户健康状态
// @Param: healthStatus appmodel.HealthStatus
// @return: error
func (h *HealthStatusService) InsertHealthStatus(healthStatus appmodel.HealthStatus) error {
	// 首先查询是否存在该记录（即用户id和日期相同）
	var record appmodel.HealthStatus
	err := global.FitnessDb.Where("uid = ? AND date = ?", healthStatus.UID, healthStatus.Date).First(&record).Error
	if err == nil {
		// 如果存在相同记录，不插入
		return errors.New("存在相同记录")
	}
	// 插入数据
	err = global.FitnessDb.Create(&healthStatus).Error
	return err
}

// UpdateHealthStatus is a method to update health status
// @Description: 更新用户健康状态
// @Param: healthStatus appmodel.HealthStatus
// @return: error
func (h *HealthStatusService) UpdateHealthStatus(healthStatus appmodel.HealthStatus) error {
	// 首先查询是否存在该记录
	var record appmodel.HealthStatus
	// 要保证更新的记录是用户自己的记录
	err := global.FitnessDb.Where("id = ?", healthStatus.ID).Where("uid = ?", healthStatus.UID).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在该记录
		return errors.New("不存在该记录")
	}
	// 更新数据
	err = global.FitnessDb.Save(&healthStatus).Error
	return err
}

// PutHealthStatus is a method to insert or update health status
// @Description: 插入用户健康状态/更新用户健康状态(如果存在相同记录则更新，否则插入)
// @Param: healthStatus appmodel.HealthStatus
// @return: error
func (h *HealthStatusService) PutHealthStatus(healthStatus appmodel.HealthStatus) error {
	// 首先查询是否存在该记录（即用户id和日期相同）
	var record appmodel.HealthStatus
	err := global.FitnessDb.Where("uid = ? AND date = ?", healthStatus.UID, healthStatus.Date).First(&record).Error
	switch {
	case err == nil:
		// 如果存在相同记录，更新数据
		err = global.FitnessDb.Model(&record).Updates(&healthStatus).Error
	case errors.Is(err, gorm.ErrRecordNotFound):
		// 如果不存在相同记录，插入数据
		err = global.FitnessDb.Create(&healthStatus).Error
	}
	return err

}

// DeleteHealthStatus is a method to delete health status
// @Description: 删除用户健康状态
// @Param: id uint, uid uint
// @return: error
func (h *HealthStatusService) DeleteHealthStatus(id uint, uid uint) error {
	// 首先查询是否存在该记录
	var record appmodel.HealthStatus
	err := global.FitnessDb.First(&record, id).Where("uid = ?", uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在该记录
		return errors.New("不存在该记录")
	}
	// 删除数据
	err = global.FitnessDb.Delete(&record).Error
	return err
}

// GetLatestHealthStatus is a method to get the latest health status of a user
// @Description: 根据用户id获取用户最近的健康状态
// @Param: uid uint
// @return: appmodel.HealthStatus, error
func (h *HealthStatusService) GetLatestHealthStatus(uid uint) (appmodel.HealthStatus, error) {
	var healthStatus appmodel.HealthStatus
	err := global.FitnessDb.Where("uid = ?", uid).Order("date desc").First(&healthStatus).Error
	return healthStatus, err
}

// GetHealthStatusList is a method to get health status list of a user
// @Description: 根据用户id获取用户的健康状态列表
// @Param: apprequest.GetHealthStatusListReq, uid uint
// @return: []appmodel.HealthStatus, int64, error
func (h *HealthStatusService) GetHealthStatusList(req apprequest.GetHealthStatusListReq, uid uint) ([]appmodel.HealthStatus, int64, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.FitnessDb.Model(&appmodel.HealthStatus{})
	var statuses []appmodel.HealthStatus

	if uid != 0 {
		db = db.Where("uid = ?", uid)
	}

	// 查询指定日期范围内的健康状态
	if !req.StartTime.IsZero() && !req.EndTime.IsZero() {
		db = db.Where("date between ? and ?", req.StartTime, req.EndTime)
	} else if !req.StartTime.IsZero() {
		db = db.Where("date >= ?", req.StartTime)
	} else if !req.EndTime.IsZero() {
		db = db.Where("date <= ?", req.EndTime)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return statuses, total, err
	}

	db.Limit(limit).Offset(offset)

	orderStr := "date desc"
	if req.Order != "" && healthOrderMap[req.Order] {
		orderStr = req.Order
		if req.Desc {
			orderStr += " desc"
		}
	}

	err = db.Order(orderStr).Find(&statuses).Error

	return statuses, total, err
}
