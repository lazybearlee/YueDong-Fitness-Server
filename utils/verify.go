package utils

import "errors"

type VerifyRules map[string][]string

type VerifyRulesMap map[string]VerifyRules

var CustomizeMap = make(map[string]VerifyRules) // 用于存储自定义规则

// RegisterRule 注册自定义规则方案建议在路由初始化层即注册
func RegisterRule(key string, rule VerifyRules) (err error) {
	if CustomizeMap[key] != nil {
		return errors.New(key + "已注册,无法重复注册")
	} else {
		CustomizeMap[key] = rule
		return nil
	}
}
