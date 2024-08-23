package core

import "github.com/gin-gonic/gin"

// InitRouter 初始化总路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 全局故障恢复与处理
	router.Use(gin.Recovery())

	// 如果是调试模式则打印日志
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	// 注册路由分组，分别是系统路由和APP路由

	// TODO: 采用中间件进行跨域处理/HTTPS处理

	return router
}
