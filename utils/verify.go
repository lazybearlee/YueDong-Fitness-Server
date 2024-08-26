package utils

import "regexp"

// EmailFormatCheck 邮箱格式校验
func EmailFormatCheck(email string) bool {
	// 正则表达式匹配邮箱格式
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
