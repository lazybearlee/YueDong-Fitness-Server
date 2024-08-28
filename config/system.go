package config

type System struct {
	DbType        string `mapstructure:"mysql-type" json:"mysql-type" yaml:"mysql-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`       // Oss类型
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`
	UseHttps      bool   `mapstructure:"use-https" json:"use-https" yaml:"use-https"`
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`    // 多点登录拦截
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                   // 使用redis
	UseMongo      bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`                   // 使用mongo
	MysqlInitData bool   `mapstructure:"mysql-init-data" json:"mysql-init-data" yaml:"mysql-init-data"` // 是否初始化数据
	AdminPassword string `mapstructure:"admin-password" json:"admin-password" yaml:"admin-password"`    // 管理员密码
}
