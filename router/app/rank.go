package approuter

import "github.com/gin-gonic/gin"

type RankRouter struct{}

func (r RankRouter) InitRankRouter(Router *gin.RouterGroup) {
	rankRouter := Router.Group("rank")
	{
		rankRouter.POST("get_rank_list", rankApi.GetRankList)
		rankRouter.POST("get_distance_rank", rankApi.GetDistanceRank)
	}
}
