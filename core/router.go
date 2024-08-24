package core

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/docs"
	"github.com/lazybearlee/yuedong-fitness/middleware"
	"github.com/lazybearlee/yuedong-fitness/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化总路由
func InitRouter() *gin.Engine {
	Router := gin.New()
	// 全局故障恢复与处理
	Router.Use(gin.Recovery())

	// 如果是调试模式则打印日志
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	// 注册路由分组，分别是系统路由和APP路由
	systemRouter := router.MainRouterGroup.System
	appRouter := router.MainRouterGroup.App
	// 注册swagger
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 注册公共路由和私人路由
	PublicRouter := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateRouter := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	// 对于公共路由，不需要进行JWT验证
	PublicRouter.Use(middleware.JWTAuth())

	// TODO: 采用中间件进行跨域处理/HTTPS处理

	return router
}
