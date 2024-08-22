package system

import (
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"gorm.io/gorm"
)

// UserService
// @description: 用户服务
// 提供登陆、注册、获取用户信息、修改用户信息、删除用户、批量删除用户、修改密码、重置密码等功能
type UserService struct{}

var User = new(UserService)

// UserRegister
// @description: 用户注册
// @param: user model.SysUser
// @return: userInter model.SysUser, err error
func (userService *UserService) UserRegister(user system.SysUser) (userInter system.SysUser, err error) {
	var u system.User
	// 查询用户名是否注册
	if !errors.Is(global.FITNESS_DB.Where("username = ?", user).First(&u).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	// 附加uuid 密码hash加密 注册
	user.Password = utils.CryptWithBcrypt(user.Password)
	user.UUID = uuid.Must(uuid.NewV4())
	err = global.FITNESS_DB.Create(&user).Error // 创建用户
	return user, err
}

// UserLogin
