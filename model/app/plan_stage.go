package appmodel

import "github.com/lazybearlee/yuedong-fitness/global"

// PlanStage 计划阶段
type PlanStage struct {
	global.BaseModel
	PlanID      uint   `json:"planID" gorm:"not null;index;comment:计划ID"`   // 计划ID，外键
	Title       string `json:"title" gorm:"not null;comment:阶段标题"`          // 阶段标题
	Description string `json:"description" gorm:"type:text;comment:阶段描述"`   // 阶段描述
	StartDate   string `json:"startDate" gorm:"not null;comment:阶段开始日期"`    // 阶段开始日期
	EndDate     string `json:"endDate" gorm:"not null;comment:阶段结束日期"`      // 阶段结束日期
	Completed   bool   `json:"completed" gorm:"default:false;comment:是否完成"` // 是否完成
}

func (p PlanStage) TableName() string {
	return "plan_stages"
}
