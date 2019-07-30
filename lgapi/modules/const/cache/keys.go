package cache

const (
	CacheFormat         = "{sso}:system:info:{@ident}"
	CacheFormatSys      = "{sso}:system:info:{@name}-{@status}-{@pi}-{@ps}"
	CacheFormatSysDel   = "{sso}:system:info:*"
	CacheFormatSysCount = "{sso}:system:info:{@name}-{@status}"
)

const (
	CacheUserSysFormat       = "{sso}:user:sys:{@sysID}-{@pi}-{@ps}"
	CacheUserSysCountFormat  = "{sso}:user:sys-count:{@sysID}"
	CacheUserDeleteFormat    = "{sso}:user:*"
	CacheUserListFormat      = "{sso}:user:list:{@userName}-{@roleID}-{@pageSize}-{@pageIndex}"
	CacheUserListAll         = "{sso}:user:list:*"
	CacheUserListCountFormat = "{sso}:user:list-count:{@userName}-{@roleID}"
	CacheUserListCountAll    = "{sso}:user:list-count:*"
	CacheUserFormat          = "{sso}:user:info:{@userID}"
	CacheUserAll             = "{sso}:user:info:*"
	CacheEmail               = "{sso}:email:{@guid}"
	CacheEamilOutTime        = 60 * 5

	LockFormat     = "{sso}:login:state-locker:{@userName}"
	CacheLoginUser = "{sso}:login:state-user:{@key}"

	//微信手机扫码登录用到
	WxLoginStateCode = "{sso}:wxlogin:state-code:{@code}"

	//微信验证码登录用到
	WechatValidcodeCacheKey           = `{sso}:wechat.validcode:@senduser`
	WechatValidcodeErrorCountCacheKey = `{sso}:wechat.validcodecount:@senduser`
)
