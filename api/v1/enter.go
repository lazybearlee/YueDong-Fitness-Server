package v1

import (
	appapi "github.com/lazybearlee/yuedong-fitness/api/v1/app"
	sysapi "github.com/lazybearlee/yuedong-fitness/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup sysapi.ApiGroup
	AppApiGroup    appapi.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
