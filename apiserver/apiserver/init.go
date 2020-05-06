package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/apiserver/services/login"
	"github.com/micro-plat/sso/apiserver/apiserver/services/member"
	"github.com/micro-plat/sso/apiserver/apiserver/services/password"
	"github.com/micro-plat/sso/apiserver/apiserver/services/permission"
	"github.com/micro-plat/sso/apiserver/apiserver/services/system"
	"github.com/micro-plat/sso/apiserver/apiserver/services/user"
	"github.com/micro-plat/sso/common/module/model"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) init() {

	//设置配置参数
	r.install()

	//挂载请求处理函数
	r.handling()

	r.Initializing(func(c component.IContainer) error {
		//检查应用配置
		var conf model.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		if err := conf.Valid(); err != nil {
			return err
		}
		model.SaveConf(c, &conf)

		if _, err := c.GetDB(); err != nil {
			return err
		}

		if _, err := c.GetCache(); err != nil {
			return err
		}
		return nil
	})

	r.Micro("/member/menu/get", member.NewMenuHandler)               //获取用户菜单数据
	r.Micro("/member/tags/get", member.NewTagHandler)                //获取用户有权限的tag数据
	r.Micro("/member/info/get", member.NewMemberHandler)             //获取用户信息
	r.Micro("/role/user/get", member.NewRoleHandler)                 //获取角色下的所有用户
	r.Micro("/member/system/get", member.NewMemberSysHandler)        //获取用户可用的子系统
	r.Micro("/member/all/get", member.NewMemberGetAllHandler)        //获取所有用户信息
	r.Micro("/system/info/get", system.NewInfoHandler)               //获取子系统信息
	r.Micro("/login/auth", login.NewAuthorizeHandler)                //用户跳转登录后的认证(不是用户名密码登录)
	r.Micro("/permission/config", permission.NewDataPerssionHandler) //【数据权限】相关接口
	r.Micro("/member/forget/password", password.NewPasswordHandler)  // 忘记密码再修改密码

	//以下接口是为sass系统使用
	r.Micro("/user", user.NewUserHandler)                  //用户相关接口
	r.Micro("/verifycode/get", login.NewVerifyCodeHandler) //生成图片验证码(这个现在没用,以后可能会用到)
}
