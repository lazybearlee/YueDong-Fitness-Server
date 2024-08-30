package appmodel

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
)

type BloodPressure struct {
	global.BaseModel
	UserId            uint `json:"userId" gorm:"comment:用户ID"`
	*sysmodel.SysUser `json:"-" gorm:"foreignKey:UserId;references:ID"`
	Systolic          int `json:"systolic" gorm:"comment:收缩压"`
	Diastolic         int `json:"diastolic" gorm:"comment:舒张压"`
}

func (b BloodPressure) TableName() string {
	return "blood_pressures"
}
