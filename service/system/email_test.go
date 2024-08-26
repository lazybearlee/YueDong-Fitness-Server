package sysservice

import (
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/spf13/viper"
	"testing"
)

func TestEmail(t *testing.T) {
	// 设置测试配置
	config := global.ConfigDefaultFile
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	_ = v.ReadInConfig()
	v.WatchConfig()
	if err := v.Unmarshal(&global.FitnessConfig); err != nil {
		panic(fmt.Errorf("unmarshal config error: %s \n", err))
	}
	// 初始化 EmailService
	emailService := EmailService{}
	// 测试发送邮件
	if err := emailService.SendValidatorMessage("Lz1958455046@outlook.com"); err != nil {
		t.Error(err)
	} else {
		t.Log("Email sent successfully")
	}
}
