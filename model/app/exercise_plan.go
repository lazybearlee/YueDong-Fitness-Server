package appmodel

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	"time"
)

// ExercisePlan 运动计划
type ExercisePlan struct {
	global.BaseModel
	UID          uint             `json:"-" gorm:"not null;index;comment:用户ID"` // 用户ID，外键，json忽略（从jwt中获取）
	SysUser      sysmodel.SysUser `gorm:"foreignKey:UID;references:ID"`         // 关联SysUser表
	Title        string           `json:"title" gorm:"not null;comment:计划标题"`
	Description  string           `json:"description" gorm:"type:text;comment:计划描述"`
	StartDate    time.Time        `json:"startDate" gorm:"not null;comment:计划开始日期"`
	EndDate      time.Time        `json:"endDate" gorm:"not null;comment:计划结束日期"`
	CurrentStage string           `json:"currentStage" gorm:"type:varchar(100);comment:当前阶段"`
	TotalStages  int              `json:"totalStages" gorm:"comment:阶段总数"`
	Stages       []PlanStage      `json:"stages" gorm:"foreignKey:PlanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Completed    bool             `json:"completed" gorm:"default:false;comment:是否完成"`
}

func (e ExercisePlan) TableName() string {
	return "exercise_plans"
}
