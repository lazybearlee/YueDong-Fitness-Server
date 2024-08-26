package utils

import (
	"fmt"
	"testing"
)

func TestEmailFormatCheck(t *testing.T) {
	// 测试邮箱格式校验函数
	validEmail := "example@example.com"
	invalidEmail := "example.com"

	fmt.Println("Valid email:", EmailFormatCheck(validEmail))     // 应输出 true
	fmt.Println("Invalid email:", EmailFormatCheck(invalidEmail)) // 应输出 false
}
