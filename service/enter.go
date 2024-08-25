package service

import (
	appservice "github.com/lazybearlee/yuedong-fitness/service/app"
	sysservice "github.com/lazybearlee/yuedong-fitness/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup sysservice.ServiceGroup
	AppServiceGroup    appservice.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
