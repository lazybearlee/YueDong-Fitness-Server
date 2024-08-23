package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/service"
)

// 拿到jwt服务
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

// JWTAuth 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取出token

	}
}
