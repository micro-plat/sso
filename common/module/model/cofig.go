package model

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
)

//Conf 应用程序配置
type Conf struct {

	//BindTimeOut 后台发送二维图片(绑定微信账户)  过期时间(秒)
	BindTimeOut int `json:"bind_timeout"`

	//RequireWxCode 登录是否需要验证码
	RequireWxCode bool `json:"require_wx_code"`

	//UserLoginFailCount 用户可以输入几次错误密码,之后用户被锁定
	UserLoginFailCount int `json:"user_login_failcount" valid:"required"`

	//UserLockTime 用户锁定时间(默认为秒数)
	UserLockTime int `json:"user_lock_time" valid:"required"`

	//WxPhoneLoginURL 微信手机登录地址
	WxPhoneLoginURL string `json:"wx_phone_login_url"`

	//WxAppID WxAppID
	WxAppID string `json:"wx_app_id"`

	//WxSecret WxSecret
	WxSecret string `json:"wx_secret"`

	//RefreshWxTokenHost 刷新微信token的host(我们内部的接口)
	RefreshWxTokenHost string `json:"refresh_wx_token_host"`

	//WxGetTokenURL 获取微信token
	WxTokenURL string `json:"wx_get_token_url"`

	//LoginValidCodeTemplateID 登录验证码模板
	LoginValidCodeTemplateID string `json:"login_validcode_template_id"`

	//WxSendTemplateMsgURL 发送微信模板消息的接口
	WxSendTemplateMsgURL string `json:"wx_send_template_msg_url"`

	AddUserUseDefaultRole string `json:"add_user_use_default_role"`
}

//Valid 验证配置参数是否合法
func (c Conf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//SaveConf 保存当前应用程序配置
func SaveConf(c component.IContainer, m *Conf) {
	c.Set("__AppConf__", m)
}

//GetConf 获取当前应用程序配置
func GetConf(c component.IContainer) *Conf {
	return c.Get("__AppConf__").(*Conf)
}
