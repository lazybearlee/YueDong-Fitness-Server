package approuter

import "github.com/gin-gonic/gin"

type RankRouter struct{}

func (r RankRouter) InitRankRouter(Router *gin.RouterGroup) {
	rankRouter := Router.Group("rank")
	{
		rankRouter.GET("get_rank_list", rankApi.GetRankList)
	}
}
