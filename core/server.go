package core

import (
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/global"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// RunServer 启动服务器
func RunServer() {
	// 初始化路由
	router := InitRouter()
	// TODO: 使用更高效的库配置http server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.FITNESS_CONFIG.System.Addr),
		Handler:        router,
		ReadTimeout:    5 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	global.FITNESS_LOG.Info("server run success on ", zap.String("address", fmt.Sprintf(":%d", global.FITNESS_CONFIG.System.Addr)))
	global.FITNESS_LOG.Error(server.ListenAndServe().Error())
}
