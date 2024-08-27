package apprequest

import (
	"github.com/lazybearlee/yuedong-fitness/model/common/request"
	"time"
)

type GetRankListRequest struct {
	request.PageInfo
	Date time.Time `json:"date" binding:"required"`
}
