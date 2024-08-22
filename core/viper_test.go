package core

import (
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/global"
	"testing"
)

// test basic viper usage

// TestBasicViperInit test basic viper init
func TestBasicViperInit(t *testing.T) {
	global.FITNESS_VIPER = ViperInit()
	fmt.Println(global.FITNESS_VIPER)
	fmt.Println(global.FITNESS_CONFIG)
}
