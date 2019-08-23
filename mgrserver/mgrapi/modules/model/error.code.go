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

	//此用户名已被使用
	ERR_USER_NAMEEXISTS = 909

	//角色名称已被使用
	ERR_ROLE_NAMEEXISTS = 910

	//系统名称或英文名称已存在
	ERR_SYS_NAMEORIDENTEXISTS = 911

	//请先保存系统根节点
	ERR_SYSFUNC_ROOTNOTEXISTS = 912
)
