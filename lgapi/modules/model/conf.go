package model

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
)

//Conf 应用程序配置
type Conf struct {
	WxLoginUrl string `json:"wxlogin_url" valid:"required"`    // 微信登录地址 (微信扫码登录)
	WxTokenUrl string `json:"wxgettoken_url" valid:"required"` // 微信获取token地址 (微信扫码登录)
	Appid      string `json:"appid" valid:"ascii,required"`    // appid (微信扫码登录)
	Secret     string `json:"secret" valid:"ascii,required"`   // secrect (微信扫码登录)

	SendCodeKey    string `json:"sendcode_key" valid:"ascii,required"`    //发微信验证码用到的加密key
	SendCodeReqUrl string `json:"sendcodereq_url" valid:"ascii,required"` //发微信验证码到哪个接口

	RequireWxLogin bool `json:"require_wx_login"` //是否支持微信扫码登录
	RequireCode    bool `json:"require_code"`     //是否支持微信验证码登录
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
