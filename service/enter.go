package service

import (
	"github.com/lazybearlee/yuedong-fitness/service/app"
	"github.com/lazybearlee/yuedong-fitness/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	AppServiceGroup    app.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
