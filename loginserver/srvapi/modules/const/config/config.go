package config

const (
	UserLoginFailLimit = 5            //登录错误次数
	UserLockTime       = 24 * 60 * 60 //好久自动解锁
	VerifyCodeTimeOut  = 5 * 60       //登录验证码过期时间
)
