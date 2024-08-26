package approuter

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.GET("get_user_info", userApi.GetUserInfo)
		userRouter.POST("update_user_info", userApi.UpdateUserInfo)
		userRouter.POST("update_user_password", userApi.UpdateUserPassword)
	}
}
