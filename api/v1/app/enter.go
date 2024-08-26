package appapi

import (
	"github.com/lazybearlee/yuedong-fitness/service"
)

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
