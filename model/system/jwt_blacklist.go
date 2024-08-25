package sysmodel

import "github.com/lazybearlee/yuedong-fitness/global"

type JwtBlacklist struct {
	global.BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (j JwtBlacklist) TableName() string {
	return "jwt_blacklists"
}
