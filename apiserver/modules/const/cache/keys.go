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
)
