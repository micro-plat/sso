package cachekey

const (
	//CacheLoginUser 给登录用户生成code
	CacheLoginUser = "{sso}:login:state-user:{@code}"

	//CacheLoginFailCount 记录登录失败次数
	CacheLoginFailCount = "{sso}:login:failcount-user:{@user_name}"

	//CacheLoginFailUnLockTime 自动解锁时间
	CacheLoginFailUnLockTime = "{sso}:login:failunlocktime-user:{@user_name}"

	//CacheWxStateCode 微信手机扫码登录,绑定用到(主要是一个凭证,传给weixin,然后又传回, 里面存的是user_id)
	CacheWxStateCode = "{sso}:wx:state-code:{@code}"

	//CacheLoginValidateCode 缓存用户登录验证码
	CacheLoginValidateCode = "{sso}:login:validate-code:{@user_name}"

	//CacheLoginValidateCodeFaildCount 记录验证码错误次数
	CacheLoginValidateCodeFaildCount = "{sso}:login:validate-code-faild-count:{@user_name}"
)

const (
	//CacheLoginUser      = "{sso}:login:state-user:{@code}"
	//CacheLoginFailCount = "{sso}:login:failcount-user:{@user_name}"
	//CacheLoginFailUnLockTime 自动解锁时间
	//CacheLoginFailUnLockTime = "{sso}:login:failunlocktime-user:{@user_name}"

	CacheLoginVerifyCode = "{sso}:login:verify-code:{@user_name}"
)
