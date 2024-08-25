package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lazybearlee/yuedong-fitness/config"
	"github.com/lazybearlee/yuedong-fitness/core/initialize"
	"github.com/lazybearlee/yuedong-fitness/global"
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
	c := global.FitnessConfig.MySQL
	CreateDBIfNotExist(&c)
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
	global.FitnessDb = db
	// 初始化表与数据
	initialize.InitDB()
}

// CreateDBIfNotExist 创建数据库
func CreateDBIfNotExist(c *config.Mysql) {
	if c.Dbname == "" {
		c.Dbname = "fitness"
	}
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Dbname)
	if c.Path == "" {
		c.Path = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", c.Username, c.Password, c.Path, c.Port)
	// 使用sql连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("数据库连接失败", zap.Error(err))
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			zap.L().Error("数据库连接关闭失败", zap.Error(err))
			panic(err)
		}
	}()
	// Ping()检查数据库连接是否正常
	err = db.Ping()
	if err != nil {
		zap.L().Error("数据库连接失败", zap.Error(err))
		panic(err)
	}
	// 创建数据库
	_, err = db.Exec(createSql)
	if err != nil {
		zap.L().Error("数据库创建失败", zap.Error(err))
		panic(err)
	}

	global.FitnessLog.Info("已确认数据库建立正常")
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
