package sysapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	sysrequest "github.com/lazybearlee/yuedong-fitness/model/system/request"
	sysresponse "github.com/lazybearlee/yuedong-fitness/model/system/response"
)

type SysUserApi struct{}

// AdminRegister
// @Tags     Admin
// @Summary  用户注册账号
// @Produce  application/json
// @Param    data  body      sysrequest.RegisterReq                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=sysresponse.UserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /admin/user/register [post]
func (b *SysUserApi) AdminRegister(c *gin.Context) {
	var register sysrequest.RegisterReq
	err := c.ShouldBindJSON(&register)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	var auths []*sysmodel.SysAuthority
	for _, v := range register.AuthorityIds {
		auths = append(auths, &sysmodel.SysAuthority{AuthorityId: v})
	}
	u := sysmodel.SysUser{
		Username:    register.Username,
		NickName:    register.NickName,
		Password:    register.Password,
		HeaderImg:   register.HeaderImg,
		AuthorityId: register.AuthorityId,
		Authorities: auths,
		Phone:       register.Phone,
		Email:       register.Email,
		Enable:      register.Enable,
	}
	ur, err := userService.UserRegister(u)
	if err != nil {
		e := fmt.Sprintf("用户注册失败: %v", err)
		global.FitnessLog.Debug(e)
		response.ErrorWithMessage(e, c)
		return
	}
	response.SuccessWithDetailed(sysresponse.UserResponse{User: ur}, "注册成功", c)
}
