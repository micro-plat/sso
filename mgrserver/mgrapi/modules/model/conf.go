package model

import (
	"encoding/json"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/sso/sso"
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

func (c Conf) ToJson() (string, error) {
	b, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//GetWebHostName 获取前端域名,上传图片使用
func (c *Conf) GetWebHostName() string {
	return c.PicHost
}

//SaveConf 保存当前应用程序配置
func SaveConf(m *Conf) error {
	cache := components.Def.Cache().GetRegularCache("gocache")
	str, err := m.ToJson()
	if err != nil {
		return err
	}

	return cache.Set("__AppConf__", str, -1)
}

//GetConf 获取当前应用程序配置
func GetConf() *Conf {
	cache := components.Def.Cache().GetRegularCache("gocache")
	confStr, err := cache.Get("__AppConf__")
	if err != nil {
		panic(err)
	}
	var m Conf
	err = json.Unmarshal([]byte(confStr), &m)
	if err != nil {
		panic(err)
	}
	return &m
}

//SaveSSOClient  保存sso client
func SaveSSOClient(m *sso.Client) error {
	cache := components.Def.Cache().GetRegularCache("gocache")
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return cache.Set("__SsoClient__", string(b), -1)
}

//GetSSOClient  获取sso client
func GetSSOClient() *sso.Client {
	cache := components.Def.Cache().GetRegularCache("gocache")
	confStr, err := cache.Get("__SsoClient__")
	if err != nil {
		panic(err)
	}
	var m sso.Client
	err = json.Unmarshal([]byte(confStr), &m)
	if err != nil {
		panic(err)
	}
	return &m
}
