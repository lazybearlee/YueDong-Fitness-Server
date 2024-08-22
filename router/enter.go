package router

import (
	"github.com/lazybearlee/yuedong-fitness/router/app"
	"github.com/lazybearlee/yuedong-fitness/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	App    app.RouterGroup
}

var routerGroup = new(RouterGroup)
