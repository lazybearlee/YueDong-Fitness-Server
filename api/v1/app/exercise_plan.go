package appapi

import (
	"github.com/gin-gonic/gin"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	apprequest "github.com/lazybearlee/yuedong-fitness/model/app/request"
	"github.com/lazybearlee/yuedong-fitness/model/common/response"
	"github.com/lazybearlee/yuedong-fitness/utils"
)

type ExercisePlanApi struct{}

// CreateExercisePlan 创建训练计划
// @Tags Plan
// @Summary 创建训练计划
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body appmodel.ExercisePlan true "创建训练计划"
// @Success 200 {object} response.Response{data=string} "创建训练计划"
// @Router /plan/create_exercise_plan [post]
func (ep *ExercisePlanApi) CreateExercisePlan(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var e appmodel.ExercisePlan
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	e.UID = uid
	// 检查开始日期和结束日期
	if err := utils.StartEndFormatCheck(e.StartDate, e.EndDate); err != nil {
		response.ErrorWithMessage(err.Error(), c)
		return
	}

	// 插入数据
	err = exercisePlanService.CreateExercisePlan(&e)
	if err != nil {
		response.ErrorWithMessage("创建训练计划失败", c)
		return
	}

	response.SuccessWithMessage("创建训练计划成功", c)
}

// GetExercisePlans 获取训练计划
// @Tags Plan
// @Summary 	获取训练计划
// @Security 	ApiKeyAuth
// @Accept  	application/json
// @Produce  	application/json
// @Param 		data 	body 				apprequest.SearchExercisePlanParams 		true 			"查询运动计划参数"
// @Success 	200 	{object} 			response.Response{data=response.PageResponse,msg=string} 	"获取训练计划，返回包括列表，总数，页数，页大小"
// @Router 		/plan/get_exercise_plans 	[get]
func (ep *ExercisePlanApi) GetExercisePlans(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var params apprequest.SearchExercisePlanParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	params.UID = uid
	// 检查开始日期和结束日期
	if err := utils.StartEndFormatCheck(params.StartDate, params.EndDate); err != nil {
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	params.Page, params.PageSize = utils.PageFormatCheck(params.Page, params.PageSize)

	plans, total, err := exercisePlanService.GetExercisePlans(&params)
	if err != nil {
		response.ErrorWithMessage("获取训练计划失败", c)
		return
	}
	// 返回数据
	response.SuccessWithDetailed(response.PageResponse{
		List:     plans,
		Total:    total,
		Page:     params.Page,
		PageSize: params.PageSize,
	}, "获取训练计划成功", c)
}

// GetAllExercisePlans 获取所有训练计划
// @Tags Plan
// @Summary 获取所有训练计划
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]appmodel.ExercisePlan} "获取所有训练计划"
// @Router /plan/get_all_exercise_plans [get]
func (ep *ExercisePlanApi) GetAllExercisePlans(c *gin.Context) {
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	plans, err := exercisePlanService.GetAllExercisePlans(uid)
	if err != nil {
		response.ErrorWithMessage("获取所有训练计划失败", c)
		return
	}

	response.SuccessWithDetailed(plans, "获取所有训练计划成功", c)
}

// UpdateExercisePlan 更新训练计划
// @Tags Plan
// @Summary 更新训练计划
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body appmodel.ExercisePlan true "更新训练计划"
// @Success 200 {object} response.Response{data=string} "更新训练计划"
// @Router /plan/update_exercise_plan [put]
func (ep *ExercisePlanApi) UpdateExercisePlan(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var e appmodel.ExercisePlan
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}
	e.UID = uid
	// 检查开始日期和结束日期
	if err := utils.StartEndFormatCheck(e.StartDate, e.EndDate); err != nil {
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	if e.CreatedAt.IsZero() {
		response.ErrorWithMessage("创建时间不能为空", c)
		return
	}

	// 更新数据
	err = exercisePlanService.UpdateExercisePlan(&e)
	if err != nil {
		response.ErrorWithMessage("更新训练计划失败", c)
		return
	}

	response.SuccessWithMessage("更新训练计划成功", c)
}

// DeleteExercisePlans 删除训练计划
// @Tags Plan
// @Summary 删除训练计划
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body apprequest.DeleteExercisePlansParams true "删除训练计划"
// @Success 200 {object} response.Response{data=string} "删除训练计划"
// @Router /plan/delete_exercise_plans [delete]
func (ep *ExercisePlanApi) DeleteExercisePlans(c *gin.Context) {
	// 拿到token
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	var e apprequest.DeleteExercisePlansParams
	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.ErrorWithMessage("参数绑定失败", c)
		return
	}

	// 删除数据
	err = exercisePlanService.DeleteExercisePlan(e, uid)
	if err != nil {
		response.ErrorWithMessage("删除训续计划失败", c)
		return
	}

	response.SuccessWithMessage("删除训练计划成功", c)
}

// GetUnCompletedExercisePlans 获取未完成的训练计划
// @Tags Plan
// @Summary 获取未完成的训练计划
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]appmodel.ExercisePlan} "获取未完成的训续计划"
// @Router /plan/get_uncompleted_exercise_plans [get]
func (ep *ExercisePlanApi) GetUnCompletedExercisePlans(c *gin.Context) {
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	plans, err := exercisePlanService.GetUnCompletedExercisePlans(uid)
	if err != nil {
		response.ErrorWithMessage("获取未完成的训练计划失败", c)
		return
	}

	response.SuccessWithDetailed(plans, "获取未完成的训续计划成功", c)
}

// GetStartedExercisePlans 获取已开始的训练计划
// @Tags Plan
// @Summary 获取已开始的训练计划
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]appmodel.ExercisePlan} "获取已开始的训续计划"
// @Router /plan/get_started_exercise_plans [get]
func (ep *ExercisePlanApi) GetStartedExercisePlans(c *gin.Context) {
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.ErrorWithMessage("获取用户信息失败", c)
		return
	}

	plans, err := exercisePlanService.GetStartedExercisePlans(uid)
	if err != nil {
		response.ErrorWithMessage("获取已开始的训练计划失败", c)
		return
	}

	response.SuccessWithDetailed(plans, "获取已开始的训续计划成功", c)
}
