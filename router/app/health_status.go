package approuter

import "github.com/gin-gonic/gin"

type HealthStatusRouter struct{}

func (r HealthStatusRouter) InitHealthStatusRouter(Router *gin.RouterGroup) {
	healthStatusRouter := Router.Group("health_status")
	{
		healthStatusRouter.PUT("put_health_status", healthStatusApi.PutHealthStatus)
		healthStatusRouter.GET("get_health_status", healthStatusApi.GetHealthStatus)
	}

}
