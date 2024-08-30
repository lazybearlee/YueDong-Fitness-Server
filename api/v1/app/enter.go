package appapi

import (
	"github.com/lazybearlee/yuedong-fitness/service"
)

type ApiGroup struct {
	UserApi
	RecordApi
	RankApi
	HealthStatusApi
	ExercisePlanApi
	HeartRateApi
	BloodPressureApi
}

var (
	userService           = service.ServiceGroupApp.SystemServiceGroup.UserService
	exerciseRecordService = service.ServiceGroupApp.AppServiceGroup.ExerciseRecordService
	rankService           = service.ServiceGroupApp.AppServiceGroup.RankService
	healthStatusService   = service.ServiceGroupApp.AppServiceGroup.HealthStatusService
	exercisePlanService   = service.ServiceGroupApp.AppServiceGroup.ExercisePlanService
	fileService           = service.ServiceGroupApp.SystemServiceGroup.FileService
	heartRateService      = service.ServiceGroupApp.AppServiceGroup.HeartRateService
	bloodPressureService  = service.ServiceGroupApp.AppServiceGroup.BloodPressureService
)
