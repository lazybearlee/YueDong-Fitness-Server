package core

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	"testing"
)

func TestJWTInit(t *testing.T) {
	global.FITNESS_VIPER = ViperInit()   // 初始化Viper, 读取配置文件
	global.FITNESS_LOG = ZapLoggerInit() // 初始化Zap日志
	GormDBInit()                         // 初始化Gorm数据库
	JWTInit()                            // 初始化JWT
}
