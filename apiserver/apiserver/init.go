package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/apiserver/services/login"
	"github.com/micro-plat/sso/apiserver/apiserver/services/member"
	"github.com/micro-plat/sso/apiserver/apiserver/services/permission"
	"github.com/micro-plat/sso/apiserver/apiserver/services/system"
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

	r.Micro("/member/menu/get", member.NewMenuHandler)        //获取用户菜单数据
	r.Micro("/member/tags/get", member.NewTagHandler)         //获取用户有权限的tag数据
	r.Micro("/member/info/get", member.NewMemberHandler)      //获取用户信息
	r.Micro("/member/system/get", member.NewMemberSysHandler) //获取用户可用的子系统
	r.Micro("/member/all/get", member.NewMemberGetAllHandler) //获取所有用户信息
	r.Micro("/system/info/get", system.NewInfoHandler)        //获取子系统信息
	r.Micro("/login/auth", login.NewAuthorizeHandler)         //用户登录认证

	r.Micro("/permission/data", permission.NewDataPerssionHandler) //【数据权限】相关接口
}
