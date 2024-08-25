package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"github.com/spf13/viper"
	"log"
	"os"
)

// this file is used to init viper

func ViperInit() {
	var config string
	// check the cmd params
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	if config == "" {
		// if the cmd params is empty, use the default config file
		if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" {
			switch gin.Mode() {
			case gin.DebugMode:
				config = global.ConfigDefaultFile
			case gin.ReleaseMode:
				config = global.ConfigReleaseFile
			case gin.TestMode:
				config = global.ConfigTestFile
			}
		} else {
			config = configEnv
		}
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	// first, check the config file exists
	if !utils.FileExists(config) {
		panic(fmt.Errorf("viper's config file [%s] not found\n", config))
	}

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()

	// when the config file is changed, print the log
	v.OnConfigChange(func(e fsnotify.Event) {
		// print the log
		log.Println("config file changed:", e.Name)
		// read the new config
		if err = v.Unmarshal(&global.FitnessConfig); err != nil {
			panic(fmt.Errorf("unmarshal config error: %s \n", err))
		}
	})
	// unmarshal the config
	if err = v.Unmarshal(&global.FitnessConfig); err != nil {
		panic(fmt.Errorf("unmarshal config error: %s \n", err))
	}

	global.FitnessViper = v
}
