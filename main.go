package main

import "github.com/lazybearlee/yuedong-fitness/core"

func main() {
	core.Init()      // 初始化数据库、日志、定时任务等
	core.RunServer() // 启动服务器
}
