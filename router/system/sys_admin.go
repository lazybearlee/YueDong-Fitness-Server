package sysrouter

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (*AdminRouter) InitAdminRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	return Router.Group("admin") // TODO: .Use(middleware.OperationRecord()) 使用操作记录中间件
}
