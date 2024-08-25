package sysrequest

import "github.com/lazybearlee/yuedong-fitness/model/system"

// RegisterReq Register request struct
type RegisterReq struct {
	Username     string `json:"username" form:"username" binding:"required"`
	Password     string `json:"password" form:"password" binding:"required"`
	NickName     string `json:"nickName" form:"nickName" binding:"required"`
	HeaderImg    string `json:"headerImg" form:"headerImg"`
	AuthorityId  uint   `json:"authorityId" form:"authorityId" binding:"required"`
	Enable       int    `json:"enable" form:"enable"`
	AuthorityIds []uint `json:"authorityIds" form:"authorityIds"`
	Phone        string `json:"phone" form:"phone"`
	Email        string `json:"email" form:"email"`
}

// LoginReq User login structure
type LoginReq struct {
	Username  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Captcha   string `json:"captcha" form:"captcha"`
	CaptchaId string `json:"captchaId" form:"captchaId"`
}

// ChangePasswordReq Modify password structure，从 JWT 中提取 user id，避免越权
type ChangePasswordReq struct {
	ID          uint   `json:"-"`
	Password    string `json:"password" form:"password"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

// SetUserAuthReq Modify user's auth structure，不需要用户ID，因为从 JWT 中提取 user id
type SetUserAuthReq struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"`
}

// SetUserAuthoritiesReq Modify user's auth structure
type SetUserAuthoritiesReq struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds" form:"authorityIds"`
}

// ChangeUserInfoReq Modify user's info structure
type ChangeUserInfoReq struct {
	ID           uint                    `gorm:"primarykey;" binding:"required"`                                                       // 主键ID
	NickName     string                  `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string                  `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint                  `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string                  `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string                  `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string                  `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Enable       int                     `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities  []sysmodel.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
