package model

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
)

//Conf 应用程序配置
type Conf struct {
	//UserLoginFailCount 用户可以输入几次错误密码,之后用户被锁定
	UserLoginFailCount int `json:"user_login_failcount" valid:"required"`

	//UserLockTime 用户锁定时间(默认为秒数)
	UserLockTime int `json:"user_lock_time" valid:"required"`
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
