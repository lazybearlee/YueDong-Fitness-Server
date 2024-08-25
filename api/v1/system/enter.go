package sysapi

import "github.com/lazybearlee/yuedong-fitness/service"

type ApiGroup struct {
	BaseApi
}

var (
	jwtService       = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService    = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
)
