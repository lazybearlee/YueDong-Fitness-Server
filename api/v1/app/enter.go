package appapi

import (
	"github.com/lazybearlee/yuedong-fitness/service"
)

type ApiGroup struct {
	UserApi
	RecordApi
}

var (
	userService           = service.ServiceGroupApp.SystemServiceGroup.UserService
	exerciseRecordService = service.ServiceGroupApp.AppServiceGroup.ExerciseRecordService
)
