package model

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

//Config 配置信息
type Config struct {
	//ApiHost ssoApi地址(不是跳转地址)
	Host string `valid:"required"`

	/*系统标识*/
	Ident string `valid:"required"`

	/*系统秘钥*/
	Secret string `valid:"required"`
}

//Valid 验证传入参数
func (c Config) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("sso 调用配置有误:%v", err)
	}
	return nil
}
