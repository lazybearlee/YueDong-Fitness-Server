package appmodel

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
)

type BloodPressure struct {
	global.BaseModel
	UID       uint              `json:"-" gorm:"not null;index;comment:用户ID"`  // 用户ID，外键
	SysUser   *sysmodel.SysUser `json:"-" gorm:"foreignKey:UID;references:ID"` // 关联SysUser表
	Systolic  int               `json:"systolic" gorm:"comment:收缩压"`
	Diastolic int               `json:"diastolic" gorm:"comment:舒张压"`
}

func (b BloodPressure) TableName() string {
	return "blood_pressures"
}
