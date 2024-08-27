package approuter

import v1 "github.com/lazybearlee/yuedong-fitness/api/v1"

type RouterGroup struct {
	UserRouter
	ExerciseRecordRouter
}

var (
	userApi           = v1.ApiGroupApp.AppApiGroup.UserApi
	exerciseRecordApi = v1.ApiGroupApp.AppApiGroup.RecordApi
)
