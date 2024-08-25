package core

import (
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/global"
	"testing"
)

// test basic viper usage

// TestBasicViperInit test basic viper init
func TestBasicViperInit(t *testing.T) {
	ViperInit()
	fmt.Println(global.FitnessViper)
	fmt.Println(global.FitnessConfig)
}
