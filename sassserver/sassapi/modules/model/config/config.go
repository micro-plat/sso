package config

var UserLoginFailCount = 5 //登录错误次数

var UserLockTime = 24 * 60 * 60 //好久自动解锁

var VerifyCodeTimeOut = 5 * 60 //登录验证码过期时间

var DbName = "db" //dbName

var Ident = "sso" //系统标识

//SetConfig 配置信息设置
func SetConfig(dbName, ident string) {
	DbName = dbName
	Ident = ident
}
