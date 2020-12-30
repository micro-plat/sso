package model

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

//VueConf 前端页面配置
type VueConf struct {

	//Ident .
	Ident string `json:"ident" valid:"required"`

	//LoginWebHost .
	LoginWebHost string `json:"loginWebHost"  valid:"required"`
}

//Valid 验证配置参数是否合法
func (c VueConf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("vueconf 配置文件有误:%v", err)
	}
	return nil
}
