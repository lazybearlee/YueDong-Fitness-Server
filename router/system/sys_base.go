package sysrouter

import "github.com/gin-gonic/gin"

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("verification_code", baseApi.VerificationCode)
		baseRouter.POST("register_with_code", baseApi.RegisterWithCode)
	}
}
