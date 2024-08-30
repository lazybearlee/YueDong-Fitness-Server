package appservice

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
)

type HeartRateService struct{}

func (h *HeartRateService) CreateHeartRate(rate *appmodel.HeartRate) error {
	return global.FitnessDb.Create(rate).Error
}

// GetAllHeartRateOfUser is a method to get all heart rate of user
// @Description: 获取用户所有心率数据
// @Param: uid uint
// @return: []appmodel.HeartRate
func (h *HeartRateService) GetAllHeartRateOfUser(uid uint) ([]appmodel.HeartRate, error) {
	var rates []appmodel.HeartRate
	err := global.FitnessDb.Where("uid = ?", uid).Find(&rates).Error
	return rates, err
}

// GetLatestHeartRateOfUser is a method to get latest heart rate of user
// @Description: 获取用户最新心率数据
// @Param: uid uint
// @return: appmodel.HeartRate
func (h *HeartRateService) GetLatestHeartRateOfUser(uid uint) (appmodel.HeartRate, error) {
	var rate appmodel.HeartRate
	err := global.FitnessDb.Where("uid = ?", uid).Order("updated_at desc").First(&rate).Error
	return rate, err
}

// DeleteHeartRate is a method to delete heart rate
// @Description: 删除心率数据
// @Param: id uint, uid uint
// @return: error
func (h *HeartRateService) DeleteHeartRate(id uint, uid uint) error {
	return global.FitnessDb.Where("id = ?", id).Where("uid = ?", uid).Delete(&appmodel.HeartRate{}).Error
}
