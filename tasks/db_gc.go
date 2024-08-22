package tasks

import (
	"errors"
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/model/common/gc"
	"gorm.io/gorm"
	"time"
)

var (
	gcConfigs []gc.DBGCConfig = []gc.DBGCConfig{
		// 删除jwt黑名单中过期的数据
		{
			TableName:     "jwt_blacklist",
			ComparedField: "created_at",
			ComparedValue: "168h",
		},
	}
)

// DBGc
// @Description: 定时任务，删除过期数据
func DBGc(db *gorm.DB) error {
	if db == nil {
		return errors.New(gc.ErrDBEmpty)
	}
	for _, config := range gcConfigs {
		// parse duration
		duration, err := time.ParseDuration(config.ComparedValue)
		if err != nil {
			return err
		}
		if duration < 0 {
			return errors.New(gc.ErrDurationParse)
		}
		err = db.Debug().Exec(fmt.Sprintf(gc.GCExecFmt, config.TableName, config.ComparedField), time.Now().Add(-duration)).Error
		if err != nil {
			return err
		}
	}
	return nil
}
