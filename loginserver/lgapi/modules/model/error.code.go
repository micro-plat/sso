package model

const (
	//系统被锁定
	ERR_SYS_LOCKED = 901

	//用户被锁定
	ERR_USER_LOCKED = 902

	//用户被禁用
	ERR_USER_FORBIDDEN = 903

	//登录出错,稍后再试
	ERR_LOGIN_ERROR = 904

	//用户不存在
	ERR_USER_NOTEXISTS = 905

	//没有相关系统权限
	ERR_USER_HASNOROLES = 906

	//用户名或密码错误
	ERR_USER_PWDWRONG = 907

	//用户原密码错误
	ERR_USER_OLDPWDWRONG = 908

	//绑定信息错误(绑定用户微信账号)
	ERR_BIND_INFOWRONG = 909

	//用户已绑定微信
	ERR_USER_EXISTSWX = 910

	//绑定超时(微信绑定)
	ERR_BIND_TIMEOUT = 911
)
