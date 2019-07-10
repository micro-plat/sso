package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/services/system"
	"github.com/micro-plat/sso/apiserver/services/users"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) init() {
	r.Initializing(func(c component.IContainer) error {

		if _, err := c.GetDB(); err != nil {
			return err
		}

		if _, err := c.GetCache(); err != nil {
			return err
		}
		return nil
	})

	r.Micro("/subsys/login", users.NewLoginHandler)   //子系统远程登录
	r.Micro("/subsys/menu", users.NewMenuHandler)     //子系统获取菜单数据
	r.Micro("/subsys/user", users.NewUserInfoHandler) //子系统,获取用户信息
	r.Micro("/subsys/pwd", users.NewPwdHandler)       //子系统,修改密码
	r.Micro("/subsys/info", system.NewInfoHandler)    //子系统,获取系统信息
}
