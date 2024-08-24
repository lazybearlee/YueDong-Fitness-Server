package oss

import "github.com/lazybearlee/yuedong-fitness/global"

type OSS struct{}

func NewOSS() *OSS {
	switch global.FITNESS_CONFIG.System.OssType {
	case "local":
	}
}
