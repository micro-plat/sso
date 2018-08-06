package cache


const (
	cacheFormat       = "{sso}:system:info:{@ident}"
	cacheFormatSys    = "{sso}:system:info:{@name}-{@status}-{@pi}-{@ps}"
	cacheFormatSysDel = "{sso}:system:info:*"
	cacheFormatSysCount = "{sso}:system:info:{@name}-{@status}"
) 


const (
	cacheUserListFormat      = "{sso}:user:list:{@userName}-{@roleID}-{@pageSize}-{@pageIndex}"
	cacheUserListAll         = "{sso}:user:list:*"
	cacheUserListCountFormat = "{sso}:user:list-count:{@userName}-{@roleID}"
	cacheUserListCountAll    = "{sso}:user:list-count:*"
	cacheUserFormat          = "{sso}:user:info:{@userID}"
	cacheUserAll             = "{sso}:user:info:*"
	cacheEmail				 = "{sso}:email:{@guid}"
)