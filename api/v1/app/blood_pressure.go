package appapi

import (
	"github.com/gin-gonic/gin"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"github.com/lazybearlee/yuedong-fitness/utils"
)

type BloodPressureApi struct{}

// CreateBloodPressure
// @Tags BloodPressure
// @Summary 创建血压
// @Security ApiKeyAuth
// @Produce json
// @Param data body appmodel.BloodPressure true "创建血压"
// @Success 200 {object} response.Response{data=string} "创建血压"
// @Router /blood_pressure/create_blood_pressure [post]
func (b *BloodPressureApi) CreateBloodPressure(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var pressure appmodel.BloodPressure
	err := c.ShouldBindJSON(&pressure)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}

	pressure.UserId = uid

	// 插入数据
	err = bloodPressureService.CreateBloodPressure(&pressure)
	if err != nil {
		response.ErrorWithMessage("创建血压失败", c)
		return
	}

	response.SuccessWithMessage("创建血压成功", c)

}

// GetAllBloodPressureOfUser
// @Tags BloodPressure
// @Summary 获取用户所有血压数据
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} response.Response{data=[]appmodel.BloodPressure} "获取用户所有血压数据"
// @Router /blood_pressure/get_all_blood_pressure_of_user [get]
func (b *BloodPressureApi) GetAllBloodPressureOfUser(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	pressures, err := bloodPressureService.GetAllBloodPressureOfUser(uid)
	if err != nil {
		response.ErrorWithMessage("获取用户所有血压数据失败", c)
		return
	}

	response.SuccessWithDetailed(pressures, "获取用户所有血压数据成功", c)
}

// GetLatestBloodPressureOfUser
// @Tags BloodPressure
// @Summary 获取用户最新血压数据
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} response.Response{data=appmodel.BloodPressure} "获取用户最新血压数据"
// @Router /blood_pressure/get_latest_blood_pressure_of_user [get]
func (b *BloodPressureApi) GetLatestBloodPressureOfUser(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	pressure, err := bloodPressureService.GetLatestBloodPressureOfUser(uid)
	if err != nil {
		response.ErrorWithMessage("获取用户最新血压数据失败", c)
		return
	}

	response.SuccessWithDetailed(pressure, "获取用户最新血压数据成功", c)
}

// DeleteBloodPressure
// @Tags BloodPressure
// @Summary 删除血压数据
// @Security ApiKeyAuth
// @Produce json
// @Param id query int true "ID"
// @Success 200 {object} response.Response{data=string} "删除血压数据"
// @Router /blood_pressure/delete_blood_pressure [delete]
func (b *BloodPressureApi) DeleteBloodPressure(c *gin.Context) {
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
	err = bloodPressureService.DeleteBloodPressure(id, uid)
	if err != nil {
		response.ErrorWithMessage("删除血压数据失败", c)
		return
	}
	response.SuccessWithMessage("删除血压数据成功", c)
}
