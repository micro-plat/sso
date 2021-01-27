package model

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

//WxBindSecrect 微信绑定的加密串
const WxBindSecrect = `f19179a06954edd0`

var cacheAppConf *LoginConf

//LoginConf 应用程序配置
type LoginConf struct {

	//RequireValidCode 登录是否需要验证码
	RequireValidCode bool `json:"require_valid_code"`

	//ValidCodeType 验证码类型 sms:短信，wechat:微信，aliapy:支付宝
	ValidCodeType string `json:"validcode_type"`

	//消息模板
	SMSTemplateID string `json:"sms_template_id"`

	//SmsSendURL 消息发送地址
	SmsSendURL string `json:"sms_send_url"`

	//UserLoginFailLimit 用户可以输入几次错误密码,之后用户被锁定
	UserLoginFailLimit int `json:"user_loginfail_limit" valid:"required"`

	//UserLockTime 用户锁定时间(默认为秒数)
	UserLockTime int `json:"user_lock_time" valid:"required"`

	//QRCodeTimeOut 后台发送二维图片(绑定微信账户)  过期时间(秒)
	QRCodeTimeOut int `json:"qrcode_timeout"`

	//WechatAppID x
	WechatAppID string `json:"wechat_app_id"`

	//WechatSecret x
	WechatSecret string `json:"wechat_secret"`

	//WechatTokenHost 刷新微信token的host(我们内部的接口)
	WechatTokenHost string `json:"wechat_token_host"`
}

//Valid 验证配置参数是否合法
func (c LoginConf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("var/loginconf/app 配置文件有误:%v", err)
	}
	return nil
}

//SaveLoginConf 保存当前应用程序配置
func SaveLoginConf(m *LoginConf) error {
	cacheAppConf = m
	return nil
}

//GetLoginConf 获取当前应用程序配置
func GetLoginConf() *LoginConf {
	return cacheAppConf
}
