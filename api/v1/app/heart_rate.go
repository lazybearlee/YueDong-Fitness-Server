package appapi

import (
	"github.com/gin-gonic/gin"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"github.com/lazybearlee/yuedong-fitness/utils"
)

type HeartRateApi struct{}

// CreateHeartRate
// @Tags HeartRate
// @Summary 创建心率
// @Security ApiKeyAuth
// @Produce json
// @Param data body appmodel.HeartRate true "创建心率"
// @Success 200 {object} response.Response{data=string} "创建心率"
// @Router /heart_rate/create_heart_rate [post]
func (h *HeartRateApi) CreateHeartRate(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var rate appmodel.HeartRate
	err := c.ShouldBindJSON(&rate)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	rate.UserId = uid

	// 插入数据
	err = heartRateService.CreateHeartRate(&rate)
	if err != nil {
		response.ErrorWithMessage("创建心率失败", c)
		return
	}

	response.SuccessWithMessage("创建心率成功", c)
}

// GetAllHeartRateOfUser
// @Tags HeartRate
// @Summary 获取用户所有心率数据
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} response.Response{data=[]appmodel.HeartRate} "获取用户所有心率数据"
// @Router /heart_rate/get_all_heart_rate_of_user [get]
func (h *HeartRateApi) GetAllHeartRateOfUser(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	rates, err := heartRateService.GetAllHeartRateOfUser(uid)
	if err != nil {
		response.ErrorWithMessage("获取用户所有心率数据失败", c)
		return
	}

	response.SuccessWithDetailed(rates, "获取用户所有心率数据成功", c)
}

// GetLatestHeartRateOfUser
// @Tags HeartRate
// @Summary 获取用户最新心率数据
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} response.Response{data=appmodel.HeartRate} "获取用户最新心率数据"
// @Router /heart_rate/get_latest_heart_rate_of_user [get]
func (h *HeartRateApi) GetLatestHeartRateOfUser(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	rate, err := heartRateService.GetLatestHeartRateOfUser(uid)
	if err != nil {
		response.ErrorWithMessage("获取用户最新心率数据失败", c)
		return
	}

	response.SuccessWithDetailed(rate, "获取用户最新心率数据成功", c)
}

// DeleteHeartRate
// @Tags HeartRate
// @Summary 删除心率数据
// @Security ApiKeyAuth
// @Produce json
// @Param id query int true "ID"
// @Success 200 {object} response.Response{data=string} "删除心率数据"
// @Router /heart_rate/delete_heart_rate [delete]
func (h *HeartRateApi) DeleteHeartRate(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	id, err := utils.StrToUInt(c.Query("id"))
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}

	// 删除数据
	err = heartRateService.DeleteHeartRate(id, uid)
	if err != nil {
		response.ErrorWithMessage("删除心率数据失败", c)
		return
	}

	response.SuccessWithMessage("删除心率数据成功", c)
}
