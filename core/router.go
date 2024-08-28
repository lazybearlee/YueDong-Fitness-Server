package core

import (
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/docs"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/middleware"
	"github.com/lazybearlee/yuedong-fitness/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
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
	// 注册文件系统
	Router.StaticFS(global.FitnessConfig.Local.StorePath, fileSystem{http.Dir(global.FitnessConfig.Local.StorePath)})
	if global.FitnessConfig.System.UseHttps {
		Router.Use(middleware.LoadTls())
	}
	// 注册swagger
	docs.SwaggerInfo.BasePath = global.FitnessConfig.System.RouterPrefix
	Router.GET(global.FitnessConfig.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 注册公共路由和私人路由
	PublicRouter := Router.Group(global.FitnessConfig.System.RouterPrefix)
	PrivateRouter := Router.Group(global.FitnessConfig.System.RouterPrefix)
	// 对于公共路由，不需要进行JWT验证
	PublicRouter.Use(middleware.LimitWithTimeUsingLocalCache()) // 限流
	PrivateRouter.Use(middleware.LimitWithTimeUsingLocalCache()).Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	// TODO: 采用中间件进行跨域处理/HTTPS处理

	{
		// 注册系统路由
		systemRouter.InitBaseRouter(PublicRouter)                  // 注册基础路由
		adminRouter := systemRouter.InitAdminRouter(PrivateRouter) // 注册管理员路由
		systemRouter.InitSysUserRouter(adminRouter)                // 注册用户路由
	}
	{
		appRouter.InitRankRouter(PublicRouter)            // 注册排行榜路由
		appRouter.InitUserRouter(PrivateRouter)           // 注册用户路由
		appRouter.InitExerciseRecordRouter(PrivateRouter) // 注册运动记录路由
		appRouter.InitHealthStatusRouter(PrivateRouter)   // 注册健康状态路由
		appRouter.InitExercisePlanRouter(PrivateRouter)   // 注册训练计划路由
	}

	global.FitnessRouters = Router.Routes()

	return Router
}
