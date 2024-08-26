package appapi

import (
	"github.com/gin-gonic/gin"
	apprequest "github.com/lazybearlee/yuedong-fitness/model/app/request"
	appresponse "github.com/lazybearlee/yuedong-fitness/model/app/response"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/utils"
)

type UserApi struct{}

// GetUserInfo
// @Tags AppUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=appresponse.UserInfo} "获取用户信息"
// @Router /user/get_user_info [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	// 拿到token中的用户信息
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}
	// 获取用户信息
	user, err := userService.UserGetInfoWithID(uid)
	if err != nil {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}
	// 封装为UserInfo
	userInfo := appresponse.UserInfo{
		Username:  user.Username,
		Nickname:  user.NickName,
		Gender:    user.Gender,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
	}
	response.SuccessWithDetailed(userInfo, "获取用户信息成功", c)
}

// UpdateUserInfo
// @Tags AppUser
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body apprequest.UserUpdateInfoReq true "更新用户信息"
// @Success 200 {object} response.Response{data=string} "更新用户信息"
// @Router /user/update_user_info [post]
func (u *UserApi) UpdateUserInfo(c *gin.Context) {
	// 拿到token中的用户信息
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}
	// 更新用户信息
	var req apprequest.UserUpdateInfoReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	// 更新用户信息
	user := sysmodel.SysUser{
		Phone:  req.Phone,
		Email:  req.Email,
		Gender: req.Gender,
	}
	user.ID = uid
	err = userService.UserSetInfo(user)
	if err != nil {
		response.ErrorWithMessage("更新用户信息失败", c)
		return
	}
	response.SuccessWithMessage("更新用户信息成功", c)
}

// UpdateUserPassword
// @Tags AppUser
// @Summary 更新用户密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body apprequest.UserUpdatePasswordReq true "更新用户密码"
// @Success 200 {object} response.Response{data=string} "更新用户密码"
// @Router /user/update_user_password [post]
func (u *UserApi) UpdateUserPassword(c *gin.Context) {
	// 拿到token中的用户信息
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}
	// 更新用户密码
	var req apprequest.UserUpdatePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	// 更新用户密码
	var user sysmodel.SysUser
	user.ID = uid
	user.Password = req.OldPassword
	_, err = userService.UserChangePassword(user, req.NewPassword)
	if err != nil {
		response.ErrorWithMessage("更新用户密码失败", c)
		return
	}
	response.SuccessWithMessage("更新用户密码成功", c)
}
