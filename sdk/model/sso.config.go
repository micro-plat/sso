package model

import "github.com/micro-plat/lib4go/logger"

var SysInfoConfig *Config

//Config 配置信息
type Config struct {
	//ApiHost ssoApi地址(不是跳转地址)
	ApiHost string

	/*系统标识*/
	Ident string

	/*系统秘钥*/
	Secret string

	/*日志*/
	Log logger.ILogger
}
