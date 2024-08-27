package core

import (
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/tasks"
	"github.com/robfig/cron/v3"
)

// TimerInit 用于启动定时任务
// 1. 定时清理DB中的垃圾数据
// ...
func TimerInit() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.FitnessTimer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := tasks.DBGc(global.FitnessDb) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}
		// 定时任务：获取用户步数排名
		_, err = global.FitnessTimer.AddTaskByFunc("RankOfStep", "10 0 * * *", func() {
			// 每天凌晨0点10分执行
			err := tasks.RankOfStep(global.FitnessDb)
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "获取用户步数排名", option...)
	}()
}
