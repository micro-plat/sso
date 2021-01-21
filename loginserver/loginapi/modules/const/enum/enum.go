package enum

const (
	UserNormal int = iota
	UserLock
	UserDisable
)

const (
	SystemDisable = 0
	SystemNormal  = 1
)

const (
	//ValidCodeTypeSMS sms
	ValidCodeTypeSMS = "sms"
	//ValidCodeTypeWechat wechat
	ValidCodeTypeWechat = "wechat"
	//ValidCodeTypeAlipay alipay
	ValidCodeTypeAlipay = "alipay"
)
