package apprequest

import (
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"github.com/lazybearlee/yuedong-fitness/model/common/request"
)

// SearchExercisePlanParams 查询运动计划参数
type SearchExercisePlanParams struct {
	request.PageInfo
	appmodel.ExercisePlan
	CheckComplete bool   `json:"check_complete"` // 是否检查完成
	Order         string `json:"order"`          // 排序字段
	Desc          bool   `json:"desc"`           // 是否倒序
}

// DeleteExercisePlansParams 删除运动计划参数
type DeleteExercisePlansParams struct {
	IDs []uint `json:"ids"`
}
