package router

import (
	approuter "github.com/lazybearlee/yuedong-fitness/router/app"
	sysrouter "github.com/lazybearlee/yuedong-fitness/router/system"
)

type Group struct {
	System sysrouter.RouterGroup
	App    approuter.RouterGroup
}

var MainRouterGroup = new(Group)
