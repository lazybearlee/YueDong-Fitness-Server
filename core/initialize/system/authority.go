package system

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"go.uber.org/zap"
)

var authorityEntities = []sysmodel.SysAuthority{
	{AuthorityId: global.AdminSuper, AuthorityName: global.AdminSuperStr, ParentId: utils.Pointer[uint](0), DefaultRouter: global.DefaultRouter},
	{AuthorityId: global.AdminUser, AuthorityName: global.AdminUserStr, ParentId: utils.Pointer[uint](0), DefaultRouter: global.DefaultRouter},
	{AuthorityId: global.CommonUser, AuthorityName: global.CommonUserStr, ParentId: utils.Pointer[uint](0), DefaultRouter: global.DefaultRouter},
}

type AuthorityInitializer struct{}

// Name 初始化器名称
func (s *AuthorityInitializer) Name() string {
	return sysmodel.SysAuthority{}.TableName()
}

// MigrateTable 初始化表
func (s *AuthorityInitializer) MigrateTable() error {
	err := global.FitnessDb.AutoMigrate(&sysmodel.SysAuthority{})
	if err != nil {
		global.FitnessLog.Error("authority初始化失败", zap.Error(err))
		return err
	} else {
		global.FitnessLog.Info("authority数据库迁移成功")
		return nil
	}
}

// TableCreated 表是否已创建
func (s *AuthorityInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&sysmodel.SysAuthority{})
}

// InitializeData 初始化数据
func (s *AuthorityInitializer) InitializeData() error {
	err := global.FitnessDb.Create(&authorityEntities).Error
	if err != nil {
		global.FitnessLog.Error("authority初始化数据失败", zap.Error(err))
		return err
	}
	// 设置数据权限
	if err := global.FitnessDb.Model(&authorityEntities[0]).Association("DataAuthorityId").Replace(
		[]*sysmodel.SysAuthority{
			{AuthorityId: global.AdminSuper},
			{AuthorityId: global.AdminUser},
			{AuthorityId: global.CommonUser},
		}); err != nil {
		global.FitnessLog.Error("authority初始化数据失败", zap.Error(err))
		return err
	}
	if err := global.FitnessDb.Model(&authorityEntities[1]).Association("DataAuthorityId").Replace(
		[]*sysmodel.SysAuthority{
			{AuthorityId: global.AdminUser},
			{AuthorityId: global.CommonUser},
		}); err != nil {
		global.FitnessLog.Error("authority初始化数据失败", zap.Error(err))
		return err
	}
	return nil
}

// DataInitialized 数据是否已插入
func (s *AuthorityInitializer) DataInitialized() bool {
	return global.FitnessDb.Where("authority_id = ?", global.AdminSuper).First(&sysmodel.SysAuthority{}).Error == nil
}
