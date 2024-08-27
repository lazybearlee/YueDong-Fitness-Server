package appapi

import (
	"github.com/gin-gonic/gin"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	apprequest "github.com/lazybearlee/yuedong-fitness/model/app/request"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"strconv"
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

// UpdateExerciseRecord update record
// @Tags Record
// @Summary 更新运动记录
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body appmodel.ExerciseRecord true "更新运动记录"
// @Success 200 {object} response.Response{data=string} "更新运动记录"
// @Router /record/update_exercise_record [put]
func (r *RecordApi) UpdateExerciseRecord(c *gin.Context) {
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

	// 更新数据
	err = exerciseRecordService.UpdateExerciseRecord(e)
	if err != nil {
		response.ErrorWithMessage("更新运动记录失败", c)
		return
	}
	response.SuccessWithMessage("更新运动记录成功", c)
}

// DeleteExerciseRecord delete record
// @Tags Record
// @Summary 删除运动记录
// @Security ApiKeyAuth
// @Produce  application/json
// @Param id path int true "ID"
// @Success 200 {object} response.Response{data=string} "删除运动记录"
// @Router /record/delete_exercise_record/{id} [delete]
func (r *RecordApi) DeleteExerciseRecord(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	// 获取参数id，id为uint类型
	id_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	id := uint(id_)

	// 删除数据
	err = exerciseRecordService.DeleteExerciseRecord(id, uid)
	if err != nil {
		response.ErrorWithMessage("删除运动记录失败", c)
		return
	}
	response.SuccessWithMessage("删除运动记录成功", c)
}

// DeleteExerciseRecords delete records
// @Tags Record
// @Summary 批量删除运动记录
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body apprequest.DeleteExerciseRecordsParams true "批量删除运动记录"
// @Success 200 {object} response.Response{data=string} "批量删除运动记录"
// @Router /record/delete_exercise_records [delete]
func (r *RecordApi) DeleteExerciseRecords(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var p apprequest.DeleteExerciseRecordsParams
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}

	// 删除数据
	err = exerciseRecordService.DeleteExerciseRecords(p.IDs, uid)
	if err != nil {
		response.ErrorWithMessage("批量删除运动记录失败", c)
		return
	}
	response.SuccessWithMessage("批量删除运动记录成功", c)
}

// GetExerciseRecord get record
// @Tags Record
// @Summary 获取运动记录
// @Security ApiKeyAuth
// @Produce  application/json
// @Param id path int true "ID"
// @Success 200 {object} response.Response{data=appmodel.ExerciseRecord} "获取运动记录"
// @Router /record/get_exercise_record/{id} [get]
func (r *RecordApi) GetExerciseRecord(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	// 获取参数id，id为uint类型
	id_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	id := uint(id_)

	// 查询数据
	record, err := exerciseRecordService.GetExerciseRecord(id)
	if err != nil {
		response.ErrorWithMessage("获取运动记录失败", c)
		return
	}
	response.SuccessWithDetailed(record, "获取运动记录成功", c)
}

// GetExerciseRecordList get record list
// @Tags Record
// @Summary 获取运动记录列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body apprequest.SearchExerciseRecordParams true "获取运动记录列表"
// @Success 200 {object} response.Response{data=response.PageResponse} "获取运动记录列表"
// @Router /record/get_exercise_record_list [post]
func (r *RecordApi) GetExerciseRecordList(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var p apprequest.SearchExerciseRecordParams
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	p.UID = uid

	// 查询数据
	list, total, err := exerciseRecordService.GetExerciseRecords(p.ExerciseRecord, p.PageInfo, p.Order, p.Desc)
	if err != nil {
		response.ErrorWithMessage("获取运动记录列表失败", c)
		return
	}
	response.SuccessWithDetailed(response.PageResponse{
		List:     list,
		Total:    total,
		Page:     p.PageInfo.Page,
		PageSize: p.PageInfo.PageSize,
	}, "获取运动记录列表成功", c)
}

// GetAllExerciseRecordListOfUser get all record list of user
// @Tags Record
// @Summary 获取用户的所有运动记录
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]appmodel.ExerciseRecord} "获取用户的所有运动记录"
// @Router /record/get_all_exercise_record_of_user [get]
func (r *RecordApi) GetAllExerciseRecordListOfUser(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	// 查询数据
	list, err := exerciseRecordService.GetAllExerciseRecordsByUID(uid)
	if err != nil {
		response.ErrorWithMessage("获取用户的所有运动记录失败", c)
		return
	}
	response.SuccessWithDetailed(list, "获取用户的所有运动记录成功", c)
}
