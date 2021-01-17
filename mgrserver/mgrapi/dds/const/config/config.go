package config

import (
	"github.com/micro-plat/lib4go/types"
)

//GetURLPrex 获取用户想加的url前缀
func GetURLPrex(options ...Option) string {
	args := getOption(options...)
	value, flag := args["prex"]
	if flag {
		return types.GetString(value)
	}
	return ""
}

//getOption 获取配置信息
func getOption(options ...Option) map[string]interface{} {
	args := make(map[string]interface{})
	for _, opt := range options {
		opt(args)
	}
	return args
}
