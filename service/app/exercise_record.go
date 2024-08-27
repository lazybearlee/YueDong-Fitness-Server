package appservice

import (
	"errors"
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
	"github.com/lazybearlee/yuedong-fitness/model/common/request"
	"gorm.io/gorm"
)

var (
	orderMap = map[string]bool{
		"id":                  true,
		"uid":                 true,
		"start_time":          true,
		"end_time":            true,
		"distance":            true,
		"calories_burned":     true,
		"steps_count":         true,
		"avg_heart_rate":      true,
		"high_blood_pressure": true,
		"low_blood_pressure":  true,
		"blood_oxygen_level":  true,
	}
)

// ExerciseRecordService exercise_record service struct
// @Description: 运动记录服务
// 提供运动记录的增删改查
type ExerciseRecordService struct{}

// InsertExerciseRecord insert exercise_record
// @Description: 插入运动记录
// @Param: exerciseRecord appmodel.ExerciseRecord
// @return: error
func (e *ExerciseRecordService) InsertExerciseRecord(exerciseRecord appmodel.ExerciseRecord) error {
	// 如果存在相同的记录，不插入（即用户id和运动开始时间相同）
	var record appmodel.ExerciseRecord
	err := global.FitnessDb.Where("uid = ? AND start_time = ?", exerciseRecord.UID, exerciseRecord.StartTime).First(&record).Error
	if err == nil {
		// 如果存在相同记录，不插入
		return errors.New("存在相同记录")
	}

	// 插入数据
	err = global.FitnessDb.Create(&exerciseRecord).Error
	return err
}

// UpdateExerciseRecord update exercise_record
// @Description: 更新运动记录
// @Param: exerciseRecord appmodel.ExerciseRecord
// @return: error
func (e *ExerciseRecordService) UpdateExerciseRecord(exerciseRecord appmodel.ExerciseRecord) error {
	// 首先查询是否存在该记录
	var record appmodel.ExerciseRecord
	// 要保证更新的记录是用户自己的记录
	err := global.FitnessDb.Where("id = ?", exerciseRecord.ID).Where("uid = ?", exerciseRecord.UID).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在该记录
		return errors.New("不存在该记录")
	}
	// 更新数据
	err = global.FitnessDb.Save(&exerciseRecord).Error
	return err
}

// DeleteExerciseRecord delete exercise_record
// @Description: 删除运动记录
// @Param: id uint
// @return: error
func (e *ExerciseRecordService) DeleteExerciseRecord(id uint, uid uint) error {
	// 首先查询是否存在该记录
	var record appmodel.ExerciseRecord
	err := global.FitnessDb.First(&record, id).Where("uid = ?", uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在该记录
		return errors.New("不存在该记录")
	}
	// 删除数据
	err = global.FitnessDb.Delete(&record).Error
	return err
}

// GetExerciseRecord get exercise_record
// @Description: 获取运动记录
// @Param: id uint
// @return: appmodel.ExerciseRecord, error
func (e *ExerciseRecordService) GetExerciseRecord(id uint) (appmodel.ExerciseRecord, error) {
	var record appmodel.ExerciseRecord
	err := global.FitnessDb.First(&record, id).Error
	return record, err
}

// GetExerciseRecords get exercise_records
// @Description: 获取运动记录列表
// @Param: record appmodel.ExerciseRecord, pageInfo request.PageInfo, order string, desc bool
// @return: []appmodel.ExerciseRecord, int64, error
func (e *ExerciseRecordService) GetExerciseRecords(record appmodel.ExerciseRecord, pageInfo request.PageInfo, order string, desc bool) ([]appmodel.ExerciseRecord, int64, error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.FitnessDb.Model(&appmodel.ExerciseRecord{})
	var records []appmodel.ExerciseRecord

	// 查询指定用户的运动记录
	if record.UID != 0 {
		db = db.Where("uid = ?", record.UID)
	}

	// 如果开始时间不为空，查询指定时间段的运动记录
	if !record.StartTime.IsZero() {
		db = db.Where("start_time >= ?", record.StartTime)
	}

	// 如果结束时间不为空，查询指定时间段的运动记录
	if !record.EndTime.IsZero() {
		db = db.Where("end_time <= ?", record.EndTime)
	}

	// 如果距离不为0，查询大于等于该距离的记录
	if record.StepsCount != 0 {
		db = db.Where("steps_count >= ?", record.StepsCount)
	}

	if record.ExerciseType != "" {
		db = db.Where("exercise_type = ?", record.ExerciseType)
	}

	// 如果平均心率不为0，查询大于等于该心率的记录
	if record.AvgHeartRate != 0 {
		db = db.Where("avg_heart_rate >= ?", record.AvgHeartRate)
	}

	var total int64
	err := db.Count(&total).Error

	if err != nil {
		return records, total, err
	}

	db.Limit(limit).Offset(offset)

	orderStr := "id desc"
	if order != "" {
		if _, ok := orderMap[order]; !ok {
			return records, total, errors.New("非法排序字段")
		}
		orderStr = order
		if desc {
			orderStr += " desc"
		}
	}

	err = db.Order(orderStr).Find(&records).Error

	return records, total, err
}

// GetAllExerciseRecordsByUID get all exercise_records by uid
// @Description: 获取指定用户的所有运动记录
// @Param: uid uint
// @return: []appmodel.ExerciseRecord, error
func (e *ExerciseRecordService) GetAllExerciseRecordsByUID(uid uint) ([]appmodel.ExerciseRecord, error) {
	var records []appmodel.ExerciseRecord
	err := global.FitnessDb.Where("uid = ?", uid).Find(&records).Error
	return records, err
}

// DeleteExerciseRecords delete exercise_records
// @Description: 批量删除运动记录
// @Param: ids []uint
// @return: error
func (e *ExerciseRecordService) DeleteExerciseRecords(ids []uint, uid uint) error {
	return global.FitnessDb.Transaction(func(tx *gorm.DB) error {
		// 查询所有要删除的记录
		var records []appmodel.ExerciseRecord
		err := tx.Where("uid = ?", uid).Where("id IN ?", ids).Find(&records).Error // 保证用户只能删除自己的记录
		if err != nil {
			return err
		}
		// 删除记录
		err = tx.Where("id IN ?", ids).Delete(&appmodel.ExerciseRecord{}).Error
		return err
	})
}
