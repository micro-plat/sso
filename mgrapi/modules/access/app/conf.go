package xx

/*
package app

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
)

//Valid 验证配置参数是否合法
func (c Conf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//GetWebHostName 获取前端域名,上传图片使用
func (c *Conf) GetWebHostName() string {
	return c.WebHostName
}

//SaveConf 保存当前应用程序配置
func SaveConf(c component.IContainer, m *Conf) {
	c.Set("__AppConf__", m)
}

//GetConf 获取当前应用程序配置
func GetConf(c component.IContainer) *Conf {
	return c.Get("__AppConf__").(*Conf)
}
*/
