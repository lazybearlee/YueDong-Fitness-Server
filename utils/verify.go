package utils

import (
	"errors"
	"regexp"
	"time"
)

// EmailFormatCheck 邮箱格式校验
func EmailFormatCheck(email string) bool {
	// 正则表达式匹配邮箱格式
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// PageFormatCheck 分页格式校验
func PageFormatCheck(page int, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100 // 最大100条
	case pageSize <= 0:
		pageSize = 10 // 默认10条
	}
	return page, pageSize
}

// StartEndFormatCheck 开始结束时间格式校验
func StartEndFormatCheck(start, end time.Time) error {
	if start.IsZero() || end.IsZero() {
		return nil
	}
	if start.After(end) {
		return errors.New("开始日期不能晚于结束日期")
	}
	return nil
}
