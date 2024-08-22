package utils

import (
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
