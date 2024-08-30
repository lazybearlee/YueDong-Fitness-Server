package appmodel

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
)

type HeartRate struct {
	global.BaseModel
	UserId            uint `json:"userId" gorm:"comment:用户ID"`
	*sysmodel.SysUser `json:"-" gorm:"foreignKey:UserId;references:ID"`
	TPS               int `json:"tps" gorm:"comment:心率"`
}

func (h HeartRate) TableName() string {
	return "heart_rates"
}
