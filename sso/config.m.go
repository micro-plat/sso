package sso

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

//ConfigData 配置信息
type ConfigData struct {
	//ApiHost ssoApi地址(不是跳转地址)
	host string `valid:"required"`

	/*系统标识*/
	ident string `valid:"required"`

	/*系统秘钥*/
	secret string `valid:"required"`
}

//Valid 验证传入参数
func (c *ConfigData) Valid() error {
	if b, err := govalidator.ValidateStruct(c); !b {
		return fmt.Errorf("sso 调用配置有误:%v", err)
	}
	return nil
}
