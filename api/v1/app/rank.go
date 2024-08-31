package appapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	apprequest "github.com/lazybearlee/yuedong-fitness/model/app/request"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"time"
)

type RankApi struct{}

// GetRankList get rank list
// @Tags Rank
// @Summary 获取步数排行榜
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body apprequest.GetRankListRequest true "获取排行榜"
// @Success 200 {object} response.Response{data=response.PageResponse} "获取排行榜"
// @Router /rank/get_rank_list [post]
func (r *RankApi) GetRankList(c *gin.Context) {
	var req apprequest.GetRankListRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	req.Page, req.PageSize = utils.PageFormatCheck(req.Page, req.PageSize)
	// 判断查询日期是否为今日或昨日
	date := req.Date.Format("2006-01-02")
	cacheStr := ""
	switch date {
	case time.Now().Format("2006-01-02"):
		cacheStr = global.StepRankToday
	case time.Now().AddDate(0, 0, -1).Format("2006-01-02"):
		cacheStr = global.StepRankYesterday
	default:
	}
	if cacheStr != "" {
		// 从缓存中获取
		list, ok := global.FitnessCache.Get(cacheStr)
		if ok {
			length := len(list.([]appmodel.UserStepRank))
			// 根据分页参数返回数据,如果分页出错,则返回全部数据
			if (req.Page * req.PageSize) >= length {
				// 超出范围
				response.SuccessWithDetailed(response.PageResponse{
					List:     list.([]appmodel.UserStepRank),
					Total:    int64(length),
					Page:     1,
					PageSize: length,
				}, "获取排行榜成功", c)
				return
			}
			response.SuccessWithDetailed(response.PageResponse{
				List:     list.([]appmodel.UserStepRank)[(req.Page-1)*req.PageSize : req.Page*req.PageSize],
				Total:    int64(length),
				Page:     req.Page,
				PageSize: req.PageSize,
			}, "获取排行榜成功", c)
		}
	}
	// 从数据库中获取
	list, err := rankService.GetStepRank(req.Date, 50)
	if err != nil {
		response.ErrorWithMessage("获取排行榜失败", c)
		return
	}
	// 缓存
	switch cacheStr {
	case global.StepRankToday:
		global.FitnessCache.Set(global.StepRankToday, list, 10*time.Minute)
	case global.StepRankYesterday:
		global.FitnessCache.Set(global.StepRankYesterday, list, 10*time.Hour)
	}
	// 根据分页参数返回数据,如果分页出错，则返回全部数据
	length := len(list)
	if (req.Page * req.PageSize) >= length {
		// 超出范围
		response.SuccessWithDetailed(response.PageResponse{
			List:     list,
			Total:    int64(length),
			Page:     1,
			PageSize: length,
		}, "获取排行榜成功", c)
	}
	response.SuccessWithDetailed(response.PageResponse{
		List:     list[(req.Page-1)*req.PageSize : req.Page*req.PageSize],
		Total:    int64(length),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取排行榜成功", c)
}

// GetDistanceRank get distance rank
// @Tags Rank
// @Summary 获取今日距离排行榜
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body apprequest.GetRankListRequest true "获取距禽排行榜"
// @Success 200 {object} response.Response{data=response.PageResponse} "获取距离排行榜"
// @Router /rank/get_distance_rank [get]
func (r *RankApi) GetDistanceRank(c *gin.Context) {
	var req apprequest.GetDistanceRankListRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	req.Page, req.PageSize = utils.PageFormatCheck(req.Page, req.PageSize)
	// 从缓存中获取
	var list []appmodel.UserDistanceRank
	l, ok := global.FitnessCache.Get(global.DistanceRankToday)
	list = l.([]appmodel.UserDistanceRank)
	if !ok {
		// 从数据库中获取
		list, err = rankService.GetDistanceRank(50)
		if err != nil {
			response.ErrorWithMessage("获取距离排行榜失败", c)
			return
		}
		// 缓存
		global.FitnessCache.Set(global.DistanceRankToday, list, 10*time.Minute)
	}
	// 根据分页参数返回数据,如果分页出错，则返回全部数据
	length := len(list)
	if (req.Page * req.PageSize) >= length {
		// 超出范围
		response.SuccessWithDetailed(response.PageResponse{
			List:     list,
			Total:    int64(length),
			Page:     1,
			PageSize: length,
		}, "获取距离排行榜成功", c)
	}
	response.SuccessWithDetailed(response.PageResponse{
		List:     list[(req.Page-1)*req.PageSize : req.Page*req.PageSize],
		Total:    int64(length),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取距离排行榜成功", c)
}
