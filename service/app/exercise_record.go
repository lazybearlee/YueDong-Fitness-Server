package appservice

import (
	"errors"
	"github.com/lazybearlee/yuedong-fitness/global"
	appmodel "github.com/lazybearlee/yuedong-fitness/model/app"
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
