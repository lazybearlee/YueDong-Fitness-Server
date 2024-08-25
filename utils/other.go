package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

func StructSpaceTrim(target interface{}) {
	// 通过反射获取
	// 获取类型
	t := reflect.TypeOf(target)
	// 判断是否是指针类型
	if t.Kind() != reflect.Ptr {
		return
	}
	// 获取指针指向的类型
	t = t.Elem()
	// 获取值
	v := reflect.ValueOf(target).Elem()
	// 遍历字段，去除空格
	for i := 0; i < t.NumField(); i++ {
		// 判断字段类型
		switch v.Field(i).Kind() {
		case reflect.String:
			// 去除空格
			v.Field(i).SetString(strings.TrimSpace(v.Field(i).String()))
		default:
			continue
		}
	}
}

// GetRandomNumberString 获取指定长度的随机数字字符串
func GetRandomNumberString(length int) string {
	// 随机数字
	number := "0123456789"
	// 随机字符串
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, number[rand.Intn(len(number))])
	}
	return string(result)
}

// GenerateVerificationCode 随机生成一个6位数的验证码。
func GenerateVerificationCode() string {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return code
}
