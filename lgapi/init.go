package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/lgapi/services/login"
	"github.com/micro-plat/sso/lgapi/services/member"
	"github.com/micro-plat/sso/lgapi/services/system"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) init() {
	//设置配置参数
	r.install()

	//挂载请求处理函数
	r.handling()

	r.Initializing(func(c component.IContainer) error {
		if _, err := c.GetDB(); err != nil {
			return err
		}

		if _, err := c.GetCache(); err != nil {
			return err
		}
		return nil
	})

	r.API("/login", login.NewLoginHandler)            //用户登录相关
	r.API("/login/check", login.NewLoginCheckHandler) //验证用户是否已登录
	r.API("/member", member.NewMemberHandler)         //用户相关操作(修改密码等)
	r.API("/system", system.NewSystemHandler)         //获取系统信息
}
