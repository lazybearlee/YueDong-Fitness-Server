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
// @Router /health_status/post_health_status [put]
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
// @Param date body apprequest.GetHealthStatusReq true "获取用户健康状态"
// @Success 200 {object} response.Response{data=appmodel.HealthStatus} "获取用户健康状态"
// @Router /health_status/get_health_status [get]
func (h *HealthStatusApi) GetHealthStatus(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var e apprequest.GetHealthStatusReq
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	// 将日期格式化为年月日
	if !e.Date.IsZero() {
		e.Date = time.Date(e.Date.Year(), e.Date.Month(), e.Date.Day(), 0, 0, 0, 0, e.Date.Location())
	}

	// 获取数据
	data, err := healthStatusService.GetLatestHealthStatus(uid, e)
	if err != nil {
		response.ErrorWithMessage("获取用户健康状态失败", c)
		return
	}
	response.SuccessWithData(data, c)
}
