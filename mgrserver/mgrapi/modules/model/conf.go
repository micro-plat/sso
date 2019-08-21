package model

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/sdk/sso"
)

//Conf 应用程序配置
type Conf struct {
	PicHost    string `json:"pic_host" valid:"required"`
	Secret     string `json:"secret" valid:"ascii,required"`
	SsoApiHost string `json:"sso_api_host" valid:"ascii,required"`
	Ident      string `json:"ident"`
}

//Valid 验证配置参数是否合法
func (c Conf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//GetWebHostName 获取前端域名,上传图片使用
func (c *Conf) GetWebHostName() string {
	return c.PicHost
}

//SaveConf 保存当前应用程序配置
func SaveConf(c component.IContainer, m *Conf) {
	c.Set("__AppConf__", m)
}

//GetConf 获取当前应用程序配置
func GetConf(c component.IContainer) *Conf {
	return c.Get("__AppConf__").(*Conf)
}

//SaveSSOClient  保存sso client
func SaveSSOClient(c component.IContainer, m *sso.Client) {
	c.Set("__SsoClient__", m)
}

//GetSSOClient  获取sso client
func GetSSOClient(c component.IContainer) *sso.Client {
	return c.Get("__SsoClient__").(*sso.Client)
}
