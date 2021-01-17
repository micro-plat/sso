package errorcode

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

	//验证码为空
	ERR_USER_EMPTY_VALIDATECODE = 913

	//验证码过期
	ERR_VALIDATECODE_TIMEOUT = 914

	//验证码错误
	ERR_VALIDATECODE_WRONG = 915

	//二维码超时(用户系统生成的二维码时间过期)
	ERR_QRCODE_TIMEOUT = 916

	//一个微信只能绑定一个账户
	ERR_OPENID_ONLY_BIND_Once = 917

	//用户姓名已存在
	ERR_USER_FULLNAMEEXISTS = 918

	//用户没有相关页面权限
	ERR_USER_HASNOPAGEAUTHORITY = 919

	//系统已存在菜单数据
	ERR_SYSTEM_HASMENUS = 920
)
