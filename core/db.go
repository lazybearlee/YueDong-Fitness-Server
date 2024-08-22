package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/system"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GormDBInit() {
	config := global.FITNESS_CONFIG.MySQL
	if config.Dbname == "" {
		config.Dbname = "fitness"
	}
	// 构建mysql连接配置
	mysqlConfig := mysql.Config{
		DSN:                       config.Dsn(), // DSN data source name
		DefaultStringSize:         191,          // string 类型字段的默认长度
		SkipInitializeWithVersion: false,        // 根据版本自动配置
	}
	// 连接mysql
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 设置表引擎、最大空闲连接数、最大连接数
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	global.FITNESS_DB = db
	// 初始化表
	DBTablesInit()
}

func DBTablesInit() {
	db := global.FITNESS_DB
	// 注册系统表
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysAuthority{},
		system.JwtBlacklist{},
	)
	if err != nil {
		global.FITNESS_LOG.Error("register sys table failed", zap.Error(err))
		os.Exit(0)
	}
	// 注册APP表
	err = db.AutoMigrate(
	// TODO: add app tables
	)
	if err != nil {
		global.FITNESS_LOG.Error("register app table failed", zap.Error(err))
		os.Exit(0)
	}
	global.FITNESS_LOG.Info("register table success")
}
