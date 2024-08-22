package core

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	"testing"
)

func TestGormDBInit(t *testing.T) {
	global.FITNESS_VIPER = ViperInit()
	GormDBInit()
}
