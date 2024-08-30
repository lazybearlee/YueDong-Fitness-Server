package approuter

import "github.com/gin-gonic/gin"

type HeartRateRouter struct{}

func (h HeartRateRouter) InitHeartRateRouter(Router *gin.RouterGroup) {
	heartRateRouter := Router.Group("heart_rate")
	{
		heartRateRouter.POST("create_heart_rate", heartRateApi.CreateHeartRate)
		heartRateRouter.GET("get_all_heart_rate_of_user", heartRateApi.GetAllHeartRateOfUser)
		heartRateRouter.GET("get_latest_heart_rate_of_user", heartRateApi.GetLatestHeartRateOfUser)
		heartRateRouter.DELETE("delete_heart_rate", heartRateApi.DeleteHeartRate)
	}
}
