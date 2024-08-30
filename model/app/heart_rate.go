package appmodel

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
)

type HeartRate struct {
	global.BaseModel
	UID     uint              `json:"-" gorm:"not null;index;comment:用户ID"`  // 用户ID，外键
	SysUser *sysmodel.SysUser `json:"-" gorm:"foreignKey:UID;references:ID"` // 关联SysUser表
	TPS     int               `json:"tps" gorm:"comment:心率"`
}

func (h HeartRate) TableName() string {
	return "heart_rates"
}
