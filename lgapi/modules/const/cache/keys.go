package cache

const (
	//CacheLoginUser 给登录用户生成code
	CacheLoginUser = "{sso}:login:state-user:{@key}"

	//CacheLoginFailCount 记录登录失败次数
	CacheLoginFailCount = "{sso}:login:failcount-user:{@user_name}"

	//CacheLoginFailUnLockTime 自动解锁时间
	CacheLoginFailUnLockTime = "{sso}:login:failunlocktime-user:{@user_name}"
)
