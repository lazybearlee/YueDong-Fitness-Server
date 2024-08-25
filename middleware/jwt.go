package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/service"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// 拿到jwt服务
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

// JWTAuth 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取出token
		token := utils.GetToken(c)
		// 如果token为空，返回未登录或非法访问
		if token == "" {
			response.NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}
		// 如果token在黑名单中，返回您的帐户异地登陆或令牌失效
		if jwtService.IsInBlacklist(token) {
			response.NoAuth("您的帐户异地登陆或令牌失效", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		// 解析token
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			// 返回错误信息
			response.NoAuth(err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		// 将claims放入上下文
		c.Set("claims", claims)
		// 判断是否即将过期，如果即将过期，刷新token
		if j.NeedRefreshToken(claims) {
			dr, _ := utils.ParseDuration(global.FitnessConfig.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims) // 通过旧token创建新token
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
			// 如果使用多点登录，需要将旧token加入黑名单
			if global.FitnessConfig.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.FitnessLog.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.SetInBlacklist(sysmodel.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Next()

		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}
