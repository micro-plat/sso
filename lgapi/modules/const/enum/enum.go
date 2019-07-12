package enum

const UserDefaultPassword = `123456`

const (
	Normal   = 0
	Locked   = 1
	Disabled = 2
	Unlock   = 11
)

const (
	UserNormal int = iota
	UserLock
	UserDisable
)

const WxApiCode = `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx9e02ddcc88e13fd4&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect`
