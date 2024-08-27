package apprequest

import (
	"github.com/lazybearlee/yuedong-fitness/model/common/request"
	"time"
)

// GetHealthStatusListReq 获取健康状态列表请求参数
type GetHealthStatusListReq struct {
	request.PageInfo
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Order     string    `json:"order"` // 排序字段
	Desc      bool      `json:"desc"`  // 是否倒序
}

// GetHealthStatusReq 获取健康状态请求参数
type GetHealthStatusReq struct {
	ID   uint      `json:"id"`
	Date time.Time `json:"date"`
}
