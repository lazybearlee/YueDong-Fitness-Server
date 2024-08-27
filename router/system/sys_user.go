package sysrouter

import "github.com/gin-gonic/gin"

type SysUserRouter struct{}

func (s *SysUserRouter) InitSysUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("user")
	{
		sysUserRouter.POST("register", SysUserApi.AdminRegister)
	}
}
