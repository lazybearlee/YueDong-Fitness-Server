package approuter

import v1 "github.com/lazybearlee/yuedong-fitness/api/v1"

type RouterGroup struct {
	UserRouter
	ExerciseRecordRouter
	RankRouter
	HealthStatusRouter
	ExercisePlanRouter
}

var (
	userApi           = v1.ApiGroupApp.AppApiGroup.UserApi
	exerciseRecordApi = v1.ApiGroupApp.AppApiGroup.RecordApi
	rankApi           = v1.ApiGroupApp.AppApiGroup.RankApi
	healthStatusApi   = v1.ApiGroupApp.AppApiGroup.HealthStatusApi
	exercisePlanApi   = v1.ApiGroupApp.AppApiGroup.ExercisePlanApi
)
