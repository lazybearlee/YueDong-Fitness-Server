package config

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	MySQL   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Qiniu   Qiniu   `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Cors    CORS    `mapstructure:"cors" json:"cors" yaml:"cors"`
}
