package appapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/global"
	apprequest "github.com/lazybearlee/yuedong-fitness/model/app/request"
	appresponse "github.com/lazybearlee/yuedong-fitness/model/app/response"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	sysresponse "github.com/lazybearlee/yuedong-fitness/model/system/response"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"go.uber.org/zap"
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
		NickName: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Gender:   req.Gender,
		Enable:   1,
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

// UploadUserAvatar
// @Tags AppUser
// @Summary 上传用户头像
// @Security ApiKeyAuth
// @accept    multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传头像"
// @Success 200 {object} response.Response{data=string} "上传用户头像"
// @Router /user/upload_user_avatar [post]
func (u *UserApi) UploadUserAvatar(c *gin.Context) {
	// 拿到token中的用户信息
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}
	// 上传用户头像
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.FitnessLog.Error("接收文件失败", zap.Error(err))
		response.ErrorWithMessage("接收文件失败", c)
		return
	}
	// 上传头像
	file, err := fileService.UploadFile(header, "1")
	if err != nil {
		global.FitnessLog.Error("上传用户头像失败", zap.Error(err))
		response.ErrorWithMessage("上传用户头像失败", c)
		return
	}
	// 更新用户头像
	url := "http://" + global.FitnessConfig.System.Addr + fmt.Sprintf(":%d/", global.FitnessConfig.System.Port) + file.Url
	err = userService.UserSetAvatar(url, uid)
	if err != nil {
		global.FitnessLog.Error("更新用户头像失败", zap.Error(err))
		response.ErrorWithMessage("更新用户头像失败", c)
		return
	}
	response.SuccessWithDetailed(sysresponse.ExaFileResponse{File: file}, "上传用户头像成功", c)
}
