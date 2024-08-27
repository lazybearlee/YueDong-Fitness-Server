package apprequest

import (
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"github.com/lazybearlee/yuedong-fitness/model/common/request"
)

type SearchExerciseRecordParams struct {
	request.PageInfo
	appmodel.ExerciseRecord
	Order string `json:"order"` // 排序字段
	Desc  bool   `json:"desc"`  // 是否倒序
}

type DeleteExerciseRecordsParams struct {
	IDs []uint `json:"ids"`
}
