package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/modules/access/app"
	xmenu "github.com/micro-plat/sso/apiserver/modules/logic"
	"github.com/micro-plat/sso/apiserver/services/system"
	"github.com/micro-plat/sso/apiserver/services/users"
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

	r.Micro("/subsys/login", users.NewLoginHandler, "*")   //子系统远程登录
	r.Micro("/subsys/menu", users.NewMenuHandler, "*")     //子系统获取菜单数据
	r.Micro("/subsys/user", users.NewUserInfoHandler, "*") //子系统,获取用户信息
	r.Micro("/subsys/pwd", users.NewPwdHandler, "*")       //子系统,修改密码
	r.Micro("/subsys/info", system.NewInfoHandler, "*")    //子系统,获取系统信息
}
