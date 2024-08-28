package middleware

import (
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/global"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// LoadTls 加载tls，重定向到https
func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     global.FitnessConfig.System.Addr + ":443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			// 如果出现错误，请不要继续
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
