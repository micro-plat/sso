package logger

import (
	"fmt"

	"github.com/micro-plat/sso/sdk/model"
)

//Infof 记录日志
func Infof(format string, content ...interface{}) {
	if model.SysInfoConfig.Log != nil {
		model.SysInfoConfig.Log.Infof(format, content)
	} else {
		fmt.Printf(format, content)
	}
}
