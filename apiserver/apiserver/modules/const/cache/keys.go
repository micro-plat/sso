package cache

const (
	CacheLoginUser      = "{sso}:login:state-user:{@code}"
	CacheLoginFailCount = "{sso}:login:failcount-user:{@user_name}"
	//CacheLoginFailUnLockTime 自动解锁时间
	CacheLoginFailUnLockTime = "{sso}:login:failunlocktime-user:{@user_name}"

	CacheLoginVerifyCode = "{sso}:login:verify-code:{@user_name}"
)
