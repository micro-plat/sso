package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/services/login"
	"github.com/micro-plat/sso/apiserver/services/member"
	"github.com/micro-plat/sso/apiserver/services/system"
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

	r.Micro("/member/menu/get", member.NewMenuHandler)   //获取用户菜单数据
	r.Micro("/member/info/get", member.NewMemberHandler) //获取用户信息
	r.Micro("/system/info/get", system.NewInfoHandler)   //获取用户系统
	r.Micro("/login/auth", login.NewAuthorizeHandler)    //用户登录认证

}
