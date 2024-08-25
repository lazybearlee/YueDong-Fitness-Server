package sysrouter

import v1 "github.com/lazybearlee/yuedong-fitness/api/v1"

type RouterGroup struct {
	BaseRouter
}

var (
	baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
)
