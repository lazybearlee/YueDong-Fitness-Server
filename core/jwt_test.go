package core

import (
	"testing"
)

func TestJWTInit(t *testing.T) {
	ViperInit()     // 初始化Viper, 读取配置文件
	ZapLoggerInit() // 初始化Zap日志
	GormDBInit()    // 初始化Gorm数据库
	JWTInit()       // 初始化JWT
}
