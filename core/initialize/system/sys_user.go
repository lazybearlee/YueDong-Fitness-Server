package system

import (
	"github.com/gofrs/uuid/v5"
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserInitializer struct{}

func (u *UserInitializer) Name() string {
	return sysmodel.SysUser{}.TableName()
}

func (u *UserInitializer) MigrateTable() error {
	// 如果数据库连接未建立，报错
	if global.FITNESS_DB == nil {
		global.FITNESS_LOG.Error("用户初始化失败，db未初始化")
		return errors.New("用户初始化失败，db未初始化")
	}
	// 自动迁移sys_user表
	err := global.FITNESS_DB.AutoMigrate(&sysmodel.SysUser{})
	if err != nil {
		global.FITNESS_LOG.Error("用户初始化失败", zap.Error(err))
		return err
	} else {
		global.FITNESS_LOG.Info("用户数据库迁移成功")
		return nil
	}
}

func (u *UserInitializer) TableCreated() bool {
	if global.FITNESS_DB == nil {
		global.FITNESS_LOG.Error("用户查询表是否建立失败，db未初始化")
		return false
	}
	return global.FITNESS_DB.Migrator().HasTable(&sysmodel.SysUser{})
}

func (u *UserInitializer) InitializeData() error {
	if global.FITNESS_DB == nil {
		global.FITNESS_LOG.Error("用户初始化数据失败，db未初始化")
		return errors.New("用户初始化数据失败，db未初始化")
	}
	// 使用AdminPassword注册一个超级管理员一个普通用户
	password := utils.CryptWithBcrypt(global.FITNESS_CONFIG.System.AdminPassword)

	entities := []sysmodel.SysUser{
		{
			UUID:        uuid.Must(uuid.NewV4()),
			Username:    "admin",
			Password:    password,
			NickName:    "lee",
			HeaderImg:   "https://nimg.ws.126.net/?url=http%3A%2F%2Fspider.ws.126.net%2Fb8e2028eb012fdeed94e007ea8974b07.jpeg&thumbnail=660x2147483647&quality=80&type=jpg",
			AuthorityId: global.AdminSuper,
			Phone:       "17777777777",
			Email:       "333333333@qq.com",
		},
		{
			UUID:        uuid.Must(uuid.NewV4()),
			Username:    "a303176530",
			Password:    password,
			NickName:    "用户1",
			HeaderImg:   "https://nimg.ws.126.net/?url=http%3A%2F%2Fspider.ws.126.net%2Fb8e2028eb012fdeed94e007ea8974b07.jpeg&thumbnail=660x2147483647&quality=80&type=jpg",
			AuthorityId: global.CommonUser,
			Phone:       "17777777777",
			Email:       "333333333@qq.com"},
	}

	// 插入数据
	if err := global.FITNESS_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, sysmodel.SysUser{}.TableName()+"表数据初始化失败!")
	}

	// 插入关联数据
	if err := global.FITNESS_DB.Model(&entities[0]).Association("Authorities").Replace(authorityEntities); err != nil {
		return err
	}
	if err := global.FITNESS_DB.Model(&entities[1]).Association("Authorities").Replace(authorityEntities[:1]); err != nil {
		return err
	}
	return nil
}

// DataInitialized 数据是否已插入
func (u *UserInitializer) DataInitialized() bool {
	if global.FITNESS_DB == nil {
		global.FITNESS_LOG.Error("用户查询数据是否建立失败，db未初始化")
		return false
	}
	// 查询sys_user表是否有数据 admin
	var user sysmodel.SysUser
	if errors.Is(global.FITNESS_DB.Where("username = ?", "admin").Preload("Authorities").First(&user).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return len(user.Authorities) > 0 && user.Authorities[0].AuthorityId == global.AdminSuper
}
