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

	r.API("/login/check", login.NewLoginCheckHandler)       //验证用户是否已登录
	r.API("/member/login", login.NewLoginHandler)           //用户登录相关
	r.API("/member/changepwd", member.NewChangePwdHandler)  //修改密码
	r.API("/member/refresh", member.NewRefleshTokenHandler) //刷新用户token
	r.API("/member/system/get", member.NewUserSysHandler)   //获取用户可进的系统信息
	r.API("/system/get", system.NewSystemHandler)           //获取系统信息
}
