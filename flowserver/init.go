package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/flowserver/modules/app"
	xmenu "github.com/micro-plat/sso/flowserver/modules/menu"
	"github.com/micro-plat/sso/flowserver/services/base"
	"github.com/micro-plat/sso/flowserver/services/function"
	"github.com/micro-plat/sso/flowserver/services/image"
	"github.com/micro-plat/sso/flowserver/services/member"
	"github.com/micro-plat/sso/flowserver/services/menu"
	"github.com/micro-plat/sso/flowserver/services/role"
	"github.com/micro-plat/sso/flowserver/services/subsys"
	"github.com/micro-plat/sso/flowserver/services/system"
	"github.com/micro-plat/sso/flowserver/services/user"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) init() {
	//初始化
	r.Initializing(func(c component.IContainer) error {
		var conf app.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		app.SaveConf(c, &conf)
		if err := conf.Valid(); err != nil {
			return err
		}

		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}

		//检查缓存配置是否正确
		if _, err := c.GetCache(); err != nil {
			return err
		}

		xmenu.Set(c) //保存全局菜单变量
		return nil
	})

	r.Micro("/sso/login", member.NewLoginHandler, "*") //用户名密码登录
	r.Micro("/sso/menu", menu.NewMenuHandler, "*")     //系统菜单相关接口

	r.Micro("/subsys/login", subsys.NewLoginHandler, "*") //子系统远程登录
	r.Micro("/subsys/menu", subsys.NewMenuHandler, "*")   //子系统远程登录
	r.Micro("/subsys/user", subsys.NewUserHandler, "*")   //子系统,获取用户列表
	r.Micro("/subsys/pwd", subsys.NewPwdHandler, "*")     //子系统,修改密码
	r.Micro("/subsys/info", subsys.NewInfoHandler, "*")   //子系统,获取系统信息

	r.Micro("/sso/ident", system.NewSystemIdentHandler, "*") //系统信息获取

	r.Micro("/sso/member", member.NewQueryHandler, "*") //查询登录用户信息

	r.Micro("/sso/changepwd", user.NewUserPasswordHandler, "*") // 修改密码

	r.Micro("/sso/base", base.NewBaseUserHandler, "*")

	r.Micro("/sso/user", user.NewUserHandler, "*") //用户相关接口

	r.Micro("/sso/auth", role.NewRoleAuthHandler, "/user/role") //权限管理

	r.Micro("/sso/role", role.NewRoleHandler, "/user/role") //角色管理相关接口

	r.Micro("/sso/sys/manage", system.NewSystemHandler, "*") //系统管理相关接口

	r.Micro("/sso/sys/func", function.NewSystemFuncHandler, "/sys/index") //系统功能相关接口

	r.Micro("/sso/img/upload", image.NewImageHandler("./static/static/img", "http://sso.sinopecscsy.com"), "*") //图片上传

	// r.Micro("/sso/notify/info", notify.NewNotifyHandler, "*") //获取报警消息列表

	// r.Micro("/sso/notify/settings", notify.NewNotifySetHandler, "*") //报警消息设置

	//r.CRON("/sso/notify/send", notify.NewNotifySendHandler, "*") // 发送消息

}
