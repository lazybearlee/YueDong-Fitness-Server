package appservice

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	apprequest "github.com/lazybearlee/yuedong-fitness/model/app/request"
	"time"
)

var (
	planOrderMap = map[string]bool{
		"start_date": true,
		"end_date":   true,
	}
)

type ExercisePlanService struct{}

// CreateExercisePlan 创建新的运动计划
func (s *ExercisePlanService) CreateExercisePlan(plan *appmodel.ExercisePlan) error {
	plan.TotalStages = len(plan.Stages)
	return global.FitnessDb.Create(plan).Error
}

// GetExercisePlanByID 通过ID获取运动计划
func (s *ExercisePlanService) GetExercisePlanByID(id uint) (*appmodel.ExercisePlan, error) {
	var plan appmodel.ExercisePlan
	if err := global.FitnessDb.Where("id = ?", id).First(&plan).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

// GetAllExercisePlans 获取所有运动计划
func (s *ExercisePlanService) GetAllExercisePlans(uid uint) ([]appmodel.ExercisePlan, error) {
	var plans []appmodel.ExercisePlan
	if err := global.FitnessDb.Where("uid = ?", uid).Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// UpdateExercisePlan 更新运动计划
func (s *ExercisePlanService) UpdateExercisePlan(plan *appmodel.ExercisePlan) error {
	plan.UpdatedAt = time.Now()
	// 更新除了主键及创建时间、删除时间、用户ID的所有字段
	return global.FitnessDb.Model(plan).Where("uid = ?", plan.UID).Omit("id", "created_at", "deleted_at", "uid").Updates(plan).Error
}

// DeleteExercisePlan 删除运动计划
func (s *ExercisePlanService) DeleteExercisePlan(params apprequest.DeleteExercisePlansParams, uid uint) error {
	// 批量删除
	return global.FitnessDb.Where("uid = ?", uid).Delete(&appmodel.ExercisePlan{}, params.IDs).Error
}

// GetCurrentExercisePlan 获取当前日期的运动计划
func (s *ExercisePlanService) GetCurrentExercisePlan(uid uint) (*appmodel.ExercisePlan, error) {
	var plan appmodel.ExercisePlan
	today := time.Now().Format("2006-01-02")
	if err := global.FitnessDb.Where("uid = ? AND start_date <= ? AND end_date >= ?", uid, today, today).First(&plan).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

// GetStartedExercisePlans 获取已开始的运动计划
func (s *ExercisePlanService) GetStartedExercisePlans(uid uint) ([]appmodel.ExercisePlan, error) {
	var plans []appmodel.ExercisePlan
	today := time.Now()
	if err := global.FitnessDb.Where("uid = ? AND start_date <= ?", uid, today).Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// GetUnCompletedExercisePlans 获取未完成的运动计划
func (s *ExercisePlanService) GetUnCompletedExercisePlans(uid uint) ([]appmodel.ExercisePlan, error) {
	var plans []appmodel.ExercisePlan
	today := time.Now()
	if err := global.FitnessDb.Where("uid = ? AND end_date > ? AND completed = ?", uid, today, false).Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// GetExercisePlans 获取运动计划
func (s *ExercisePlanService) GetExercisePlans(params *apprequest.SearchExercisePlanParams) ([]appmodel.ExercisePlan, int64, error) {
	var plans []appmodel.ExercisePlan
	db := global.FitnessDb.Model(&appmodel.ExercisePlan{}).Where("uid = ?", params.UID)
	if !params.StartDate.IsZero() {
		db = db.Where("start_date >= ?", params.StartDate)
	}
	if !params.EndDate.IsZero() {
		db = db.Where("end_date <= ?", params.EndDate)
	}
	if params.Title != "" {
		db = db.Where("title like ?", "%"+params.Title+"%")
	}
	if params.Description != "" {
		db = db.Where("description like ?", "%"+params.Description+"%")
	}
	if params.CheckComplete {
		db = db.Where("completed = ?", params.Completed)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return plans, total, err
	}

	// offset and limit
	db.Limit(params.PageSize).Offset(params.PageSize * (params.Page - 1))

	// order
	orderStr := "id desc"
	if _, ok := planOrderMap[params.Order]; ok {
		orderStr = params.Order
		if params.Desc {
			orderStr += " desc"
		}
	}

	err = db.Order(orderStr).Find(&plans).Error

	return plans, total, err
}
