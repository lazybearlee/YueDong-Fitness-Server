package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"github.com/lazybearlee/yuedong-fitness/service"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"strconv"
	"strings"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetJWTClaims(c)
		// 获取请求的PATH
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, global.FitnessConfig.System.RouterPrefix)
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))
		e := casbinService.Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			response.ErrorWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
