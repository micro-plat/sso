package cache

const (
	//CacheLoginUser 给登录用户生成code
	CacheLoginUser = "{sso}:login:state-user:{@key}"

	//CacheLoginFailCount 记录登录失败次数
	CacheLoginFailCount = "{sso}:login:failcount-user:{@user_name}"

	//CacheLoginFailUnLockTime 自动解锁时间
	CacheLoginFailUnLockTime = "{sso}:login:failunlocktime-user:{@user_name}"

	//微信手机扫码登录,绑定用到(主要是一个凭证,传给weixin,然后又传回, 里面存的是user_id)
	CacheWxStateCode = "{sso}:wx:state-code:{@code}"
)
