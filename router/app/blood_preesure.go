package approuter

import "github.com/gin-gonic/gin"

type BloodPressureRouter struct{}

func (b BloodPressureRouter) InitBloodPressureRouter(Router *gin.RouterGroup) {
	bloodPressureRouter := Router.Group("blood_pressure")
	{
		bloodPressureRouter.POST("create_blood_pressure", bloodPressureApi.CreateBloodPressure)
		bloodPressureRouter.GET("get_all_blood_pressure_of_user", bloodPressureApi.GetAllBloodPressureOfUser)
		bloodPressureRouter.GET("get_latest_blood_pressure_of_user", bloodPressureApi.GetLatestBloodPressureOfUser)
		bloodPressureRouter.DELETE("delete_blood_pressure", bloodPressureApi.DeleteBloodPressure)
	}
}
