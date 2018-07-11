package app

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
)

//Conf 应用程序配置
type Conf struct {
	CheckURL     string `json:"qrcode-login-check-url" valid:"required"`
	AppID        string `json:"appid" valid:"ascii,required"`
	Secret       string `json:"secret" valid:"ascii,required"`
	WechatTSAddr string `json:"wechat-url" valid:"required"`
}

//Valid 验证配置参数是否合法
func (a Conf) Valid() error {
	if b, err := govalidator.ValidateStruct(&a); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//SaveConf 保存当前应用程序配置
func SaveConf(c component.IContainer, m *Conf) {

}

//GetConf 获取当前应用程序配置
func GetConf(c component.IContainer) *Conf {
	return &Conf{CheckURL: "http://192.168.5.71/check"}
}
