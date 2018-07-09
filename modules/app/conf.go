package app

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/wechat/mp"
)

//Conf 应用程序配置
type Conf struct {
	WXLoginURL   string `json:"wxlogin-url" valid:"required"`
	AppID        string `json:"appid" valid:"ascii,required"`
	Secret       string `json:"secret" valid:"ascii,required"`
	WechatTSAddr string `json:"wechat-url" valid:"required"`
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

//GetWeChatContext 获取微信操作context
func (c *Conf) GetWeChatContext(ct component.IContainer) *mp.Context {
	if mp, ok := ct.Get("__wechat_context_").(*mp.Context); ok {
		return mp
	}
	tk := mp.NewDefaultAccessTokenByURL(c.AppID, c.Secret, c.WechatTSAddr)
	wectx := mp.NewContext(tk)
	ct.Set("__wechat_context_", wectx)
	return wectx
}
