package model

import (
	"encoding/json"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/components"
)

//VueConf 前端页面配置
type VueConf struct {

	//Wxcallbackhost 微信回调host
	Wxcallbackhost string `json:"wxcallbackhost"`

	//Wxcallbackurl 微信回调url
	Wxcallbackurl string `json:"wxcallbackurl"`

	//CodeLabel 验证码说明lable
	CodeLabel string `json:"codeLabel" valid:"required"`

	//CodeHolder 短信验证码输入提示
	CodeHolder string `json:"codeHolder" valid:"required"`

	//SendBtnLable 短信验证码触发按钮说明
	SendBtnLable string `json:"sendBtnLable" valid:"required"`

	//ShowText 短信验证码发送触发成功说明
	ShowText string `json:"showText" valid:"required"`

	//StaticImageUrl 静态图片地址
	StaticImageUrl string `json:"staticImageUrl" valid:"required"`

	//CompanyRight 网站所属公司说明
	CompanyRight string `json:"companyRight" valid:"required"`

	//CompanyRightcode 网站所属公司编码
	CompanyRightCode string `json:"companyRightCode" valid:"required"`
}

//Valid 验证配置参数是否合法
func (c VueConf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("vueconf 配置文件有误:%v", err)
	}
	return nil
}

func (c VueConf) ToJson() (string, error) {
	b, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//SaveConf 保存当前应用程序配置
func SaveConf(m *VueConf) error {
	cache := components.Def.Cache().GetRegularCache("gocache")
	str, err := m.ToJson()
	if err != nil {
		return err
	}
	return cache.Set("__VueConf__", str, -1)
}

//GetConf 获取当前应用程序配置
func GetConf() string {
	cache := components.Def.Cache().GetRegularCache("gocache")
	confStr, err := cache.Get("__VueConf__")
	if err != nil {
		panic(err)
	}

	return confStr
}
