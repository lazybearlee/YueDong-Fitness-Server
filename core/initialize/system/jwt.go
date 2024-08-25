package system

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
)

type JwtBlacklistInitializer struct{}

var JwtBlacklist = new(JwtBlacklistInitializer)

// Name 初始化器名称
func (s *JwtBlacklistInitializer) Name() string {
	return sysmodel.JwtBlacklist{}.TableName()
}

// MigrateTable 初始化表
func (s *JwtBlacklistInitializer) MigrateTable() error {
	return global.FitnessDb.AutoMigrate(&sysmodel.JwtBlacklist{})
}

// TableCreated 表是否已创建
func (s *JwtBlacklistInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&sysmodel.JwtBlacklist{})
}

// InitializeData 初始化数据
func (s *JwtBlacklistInitializer) InitializeData() error {
	return nil // 无需初始化数据
}

// DataInitialized 数据是否已插入
func (s *JwtBlacklistInitializer) DataInitialized() bool {
	return true // 无需初始化数据
}
