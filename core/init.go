package core

func Init() {
	// 必选项
	ViperInit()     // 初始化Viper, 读取配置文件
	ZapLoggerInit() // 初始化Zap日志
	GormDBInit()    // 初始化Gorm数据库
	JWTInit()       // 初始化JWT
	TimerInit()     // 初始化定时任务
	RedisInit()     // 初始化Redis
}
