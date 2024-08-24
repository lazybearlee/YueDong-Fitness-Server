package sysmodel

import (
	"github.com/gofrs/uuid/v5"
	"github.com/lazybearlee/yuedong-fitness/global"
)

// Login interface
// used to only get the necessary information of the user
type Login interface {
	GetUsername() string
	GetNickname() string
	GetUUID() uuid.UUID
	GetUserId() uint
	GetAuthorityId() uint
	GetUserInfo() any
}

var _ Login = new(SysUser)

type SysUser struct {
	global.BaseModel
	UUID        uuid.UUID       `json:"uuid" gorm:"index;comment:用户UUID"`                                                                                                                                                       // 用户UUID
	Username    string          `json:"userName" gorm:"index;comment:用户登录名"`                                                                                                                                                    // 用户登录名
	Password    string          `json:"-"  gorm:"comment:用户登录密码"`                                                                                                                                                               // 用户登录密码
	NickName    string          `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                                                                                                              // 用户昵称
	SideMode    string          `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                                                                                                                            // 用户侧边主题
	HeaderImg   string          `json:"headerImg" gorm:"default:https://nimg.ws.126.net/?url=http%3A%2F%2Fspider.ws.126.net%2Fb8e2028eb012fdeed94e007ea8974b07.jpeg&thumbnail=660x2147483647&quality=80&type=jpg;comment:用户头像"` // 用户头像
	BaseColor   string          `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                                                                                                                             // 基础颜色
	AuthorityId uint            `json:"authorityId" gorm:"index;default:888;comment:用户角色ID"`                                                                                                                                    // 用户角色ID
	Authority   SysAuthority    `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []*SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string          `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email       string          `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable      int             `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}

func (s *SysUser) GetUsername() string {
	return s.Username
}

func (s *SysUser) GetNickname() string {
	return s.NickName
}

func (s *SysUser) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *SysUser) GetUserId() uint {
	return s.ID
}

func (s *SysUser) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *SysUser) GetUserInfo() any {
	return *s
}
