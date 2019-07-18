package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/lgapi/services/member"
	"github.com/micro-plat/sso/lgapi/services/user"
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
	r.API("/lg/login", member.NewLoginHandler) //用户登录相关 /lg/login/post 登录 /lg/login/check 验证是否登录 /lg/login/refresh 刷新token
	r.API("/lg/user", user.NewUserHandler)     //用户可访问的子系统
}
