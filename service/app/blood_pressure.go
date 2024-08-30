package appservice

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
)

type BloodPressureService struct{}

// CreateBloodPressure is a method to create blood pressure
// @Description: 创建血压数据
// @Param: pressure *appmodel.BloodPressure
// @return: error
func (b *BloodPressureService) CreateBloodPressure(pressure *appmodel.BloodPressure) error {
	return global.FitnessDb.Create(pressure).Error
}

// GetAllBloodPressureOfUser is a method to get all blood pressure of user
// @Description: 获取用户所有血压数据
// @Param: uid uint
// @return: []appmodel.BloodPressure
func (b *BloodPressureService) GetAllBloodPressureOfUser(uid uint) ([]appmodel.BloodPressure, error) {
	var pressures []appmodel.BloodPressure
	err := global.FitnessDb.Where("uid = ?", uid).Find(&pressures).Error
	return pressures, err
}

// GetLatestBloodPressureOfUser is a method to get latest blood pressure of user
// @Description: 获取用户最新血压数据
// @Param: uid uint
// @return: appmodel.BloodPressure
func (b *BloodPressureService) GetLatestBloodPressureOfUser(uid uint) (appmodel.BloodPressure, error) {
	var pressure appmodel.BloodPressure
	err := global.FitnessDb.Where("uid = ?", uid).Order("updated_at desc").First(&pressure).Error
	return pressure, err
}

// DeleteBloodPressure is a method to delete blood pressure
// @Description: 删除血压数据
// @Param: id uint, uid uint
// @return: error
func (b *BloodPressureService) DeleteBloodPressure(id uint, uid uint) error {
	return global.FitnessDb.Where("id = ?", id).Where("uid = ?", uid).Delete(&appmodel.BloodPressure{}).Error
}
