package appapi

import (
	"github.com/gin-gonic/gin"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"github.com/lazybearlee/yuedong-fitness/utils"
)

type RecordApi struct{}

// InsertExerciseRecord insert record
// @Tags Record
// @Summary 插入运动记录
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body appmodel.ExerciseRecord true "插入运动记录"
// @Success 200 {object} response.Response{data=string} "插入运动记录"
// @Router /record/insert_exercise_record [post]
func (r *RecordApi) InsertExerciseRecord(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var e appmodel.ExerciseRecord
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	e.UID = uid

	// TODO:做一步去重，防止重复插入

	// 插入数据
	err = exerciseRecordService.InsertExerciseRecord(e)
	if err != nil {
		response.ErrorWithMessage("插入运动记录失败", c)
		return
	}
	response.SuccessWithMessage("插入运动记录成功", c)
}
