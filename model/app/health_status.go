package appmodel

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	"time"
)

type HealthStatus struct {
	global.BaseModel
	UID            uint              `json:"-" gorm:"not null;index;comment:用户ID"`     // 用户ID，外键
	SysUser        *sysmodel.SysUser `json:"-" gorm:"foreignKey:UID;references:ID"`    // 关联SysUser表
	Date           time.Time         `json:"date" gorm:"not null;comment:记录日期"`        // 记录日期
	Height         float64           `json:"height" gorm:"comment:身高(单位: 厘米)"`         // 身高 (单位: 厘米)
	Weight         float64           `json:"weight" gorm:"comment:体重(单位: 公斤)"`         // 体重 (单位: 公斤)
	BMI            float64           `json:"bmi" gorm:"comment:BMI"`                   // BMI
	CaloriesBurned float64           `json:"caloriesBurned" gorm:"comment:消耗的卡路里"`     // 消耗的卡路里
	ExerciseTime   int               `json:"exerciseTime" gorm:"comment:锻炼时长(单位: 分钟)"` // 锻炼时长 (单位: 分钟)
	StepsCount     int               `json:"stepsCount" gorm:"comment:步数"`             // 步数
}

func (h HealthStatus) TableName() string {
	return "health_statuses"
}
