package appmodel

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	"time"
)

type ExerciseRecord struct {
	global.BaseModel
	UID               uint              `json:"-" gorm:"not null;index;comment:用户ID"`                                      // 用户ID，外键
	SysUser           *sysmodel.SysUser `json:"-" gorm:"foreignKey:UID;references:ID"`                                     // 关联SysUser表
	ExerciseType      string            `json:"exerciseType" gorm:"type:enum('cycling', 'running');not null;comment:运动类型"` // 运动类型
	Duration          int               `json:"duration" gorm:"comment:运动时长(单位: 毫秒)"`                                      // 运动时长 (单位: 毫秒)
	Distance          float64           `json:"distance" gorm:"comment:运动距离(单位: 公里)"`                                      // 运动距离 (单位: 公里)
	CaloriesBurned    float64           `json:"caloriesBurned" gorm:"comment:消耗的卡路里"`                                      // 消耗的卡路里
	StepsCount        int               `json:"stepsCount" gorm:"comment:步数"`                                              // 步数
	AvgHeartRate      float64           `json:"avgHeartRate" gorm:"comment:平均心率"`                                          // 平均心率
	HighBloodPressure float64           `json:"highBloodPressure" gorm:"comment:高压"`                                       // 高压
	LowBloodPressure  float64           `json:"lowBloodPressure" gorm:"comment:低压"`                                        // 低压
	BloodOxygenLevel  float64           `json:"bloodOxygenLevel" gorm:"comment:血氧水平"`                                      // 血氧水平
	StartTime         time.Time         `json:"startTime" gorm:"comment:运动开始时间"`                                           // 运动开始时间
	EndTime           time.Time         `json:"endTime" gorm:"comment:运动结束时间"`                                             // 运动结束时间
	LocationPath      string            `json:"locationPath" gorm:"type:text;comment:运动轨迹(存储JSON)"`                        // 运动轨迹 (存储JSON)
}

func (e ExerciseRecord) TableName() string {
	return "exercise_records"
}
