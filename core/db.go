package core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lazybearlee/yuedong-fitness/config"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/system"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func GormDBInit() {
	c := global.FITNESS_CONFIG.MySQL
	if c.Dbname == "" {
		c.Dbname = "fitness"
	}
	// 构建mysql连接配置
	mysqlConfig := mysql.Config{
		DSN:                       c.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	// 连接mysql
	db, err := gorm.Open(mysql.New(mysqlConfig), NewDBConfig(c))
	if err != nil {
		panic(err)
	}
	// 设置表引擎、最大空闲连接数、最大连接数
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	global.FITNESS_DB = db
	// 初始化表
	DBTablesInit()
}

// NewDBConfig 初始化数据库配置
func NewDBConfig(c config.Mysql) *gorm.Config {
	generalDB := c.GeneralDB
	return &gorm.Config{
		Logger: logger.New(NewWriter(generalDB, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      generalDB.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Prefix,
			SingularTable: c.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
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

// Writer 格式化打印日志
type Writer struct {
	config config.GeneralDB
	writer logger.Writer
}

// NewWriter 初始化Writer
func NewWriter(config config.GeneralDB, writer logger.Writer) *Writer {
	return &Writer{config: config, writer: writer}
}

// Printf 格式化打印日志
func (c *Writer) Printf(message string, data ...interface{}) {
	if c.config.LogZap {
		switch c.config.LogLevel() {
		case logger.Silent:
			zap.L().Debug(fmt.Sprintf(message, data...))
		case logger.Error:
			zap.L().Error(fmt.Sprintf(message, data...))
		case logger.Warn:
			zap.L().Warn(fmt.Sprintf(message, data...))
		case logger.Info:
			zap.L().Info(fmt.Sprintf(message, data...))
		default:
			zap.L().Info(fmt.Sprintf(message, data...))
		}
		return
	}
	c.writer.Printf(message, data...)
}
