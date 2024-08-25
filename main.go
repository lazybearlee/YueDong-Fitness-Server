package main

import "github.com/lazybearlee/yuedong-fitness/core"

// 运动健身服务器，启动！
// @title YueDong-Fitness-API
// @version 0.1
// @description 悦动健身API说明
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	core.Init()      // 初始化数据库、日志、定时任务等
	core.RunServer() // 启动服务器
}
