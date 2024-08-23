package system

import (
	"github.com/lazybearlee/yuedong-fitness/core"
	"github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
)

var authorityEntities = []system.SysAuthority{
	{AuthorityId: core.AdminSuper, AuthorityName: core.AdminSuperStr, ParentId: utils.Pointer[uint](0), DefaultRouter: core.DefaultRouter},
	{AuthorityId: core.AdminUser, AuthorityName: core.AdminUserStr, ParentId: utils.Pointer[uint](0), DefaultRouter: core.DefaultRouter},
	{AuthorityId: core.CommonUser, AuthorityName: core.CommonUserStr, ParentId: utils.Pointer[uint](0), DefaultRouter: core.DefaultRouter},
}

type AuthorityInitializer struct{}
