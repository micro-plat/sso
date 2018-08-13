package app

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/wechat/mp"
	"strings"
)

//Conf 应用程序配置
type Conf struct {
	AppID           string `json:"appid" valid:"ascii,required"`
	Secret          string `json:"secret" valid:"ascii,required"`
	WechatTSAddr    string `json:"wechat-url" valid:"required"`
	HostName		string `json:"hostname" valid:"required"`
}

//Valid 验证配置参数是否合法
func (c Conf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//获取绑定url
func (c *Conf) GetBindUrl() string {
	return strings.Join([]string{c.HostName,"/user/bind?email=%s"},"")
}

//获取二维码登录url
func (c *Conf) GetQRLoginCheckURL() string {
	return strings.Join([]string{c.HostName,"/member/check"},"")
}
//获取微信登录url
func (c *Conf) GetWXLoginURL() string {
	return strings.Join([]string{c.HostName,"/member/login"},"")
}

//SaveConf 保存当前应用程序配置
func SaveConf(c component.IContainer, m *Conf) {
	c.Set("__AppConf__", m)
}

//GetConf 获取当前应用程序配置
func GetConf(c component.IContainer) *Conf {
	return c.Get("__AppConf__").(*Conf)
}

//GetWeChatContext 获取微信操作context
func GetWeChatContext(ct component.IContainer) *mp.Context {
	c := GetConf(ct)
	if mp, ok := ct.Get("__wechat_context_").(*mp.Context); ok {
		return mp
	}
	tk := mp.NewDefaultAccessTokenByURL(c.AppID, c.Secret, c.WechatTSAddr)
	wectx := mp.NewContext(tk)
	ct.Set("__wechat_context_", wectx)
	return wectx
}
