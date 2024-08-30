package appapi

import (
	"github.com/gin-gonic/gin"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	apprequest "github.com/lazybearlee/yuedong-fitness/model/app/request"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"time"
)

type HealthStatusApi struct{}

// PutHealthStatus insert health status
// @Tags HealthStatus
// @Summary 插入用户健康状态/更新用户健康状态
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body appmodel.HealthStatus true "插入用户健康状态"
// @Success 200 {object} response.Response{data=string} "插入用户健康状态"
// @Router /health_status/put_health_status [put]
func (h *HealthStatusApi) PutHealthStatus(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var e appmodel.HealthStatus
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	e.UID = uid
	// 将日期格式化为年月日
	e.Date = time.Date(e.Date.Year(), e.Date.Month(), e.Date.Day(), 0, 0, 0, 0, e.Date.Location())

	// 插入数据
	err = healthStatusService.PutHealthStatus(e)
	if err != nil {
		response.ErrorWithMessage("更新用户健康状态失败", c)
		return
	}
	response.SuccessWithMessage("插入/更新用户健康状态成功", c)
}

// GetHealthStatus get health status
// @Tags HealthStatus
// @Summary 获取用户健康状态
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=appmodel.HealthStatus} "获取用户健康状态"
// @Router /health_status/get_health_status [get]
func (h *HealthStatusApi) GetHealthStatus(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	// 获取数据
	data, err := healthStatusService.GetLatestHealthStatus(uid)
	if err != nil {
		response.ErrorWithMessage("获取用户健康状态失败", c)
		return
	}
	response.SuccessWithData(data, c)
}

// GetHealthStatusList get health status list
// @Tags HealthStatus
// @Summary 获取用户健康状态列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param date body apprequest.GetHealthStatusListReq true "获取用户健康状态列表"
// @Success 200 {object} response.Response{data=response.PageResponse} "获取用户健康状态列表"
// @Router /health_status/get_health_status_list [post]
func (h *HealthStatusApi) GetHealthStatusList(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var req apprequest.GetHealthStatusListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	err = utils.StartEndFormatCheck(req.StartTime, req.EndTime)
	if err != nil {
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	req.Page, req.PageSize = utils.PageFormatCheck(req.Page, req.PageSize)
	req.StartTime = time.Date(req.StartTime.Year(), req.StartTime.Month(), req.StartTime.Day(), 0, 0, 0, 0, req.StartTime.Location())
	req.EndTime = time.Date(req.EndTime.Year(), req.EndTime.Month(), req.EndTime.Day(), 0, 0, 0, 0, req.EndTime.Location())

	// 获取数据
	list, total, err := healthStatusService.GetHealthStatusList(req, uid)
	if err != nil {
		response.ErrorWithMessage("获取用户健康状态列表失败", c)
		return
	}

	response.SuccessWithDetailed(response.PageResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取用户健康状态列表成功", c)
}
