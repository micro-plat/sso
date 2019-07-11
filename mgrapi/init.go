package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrapi/modules/model"
	"github.com/micro-plat/sso/mgrapi/services/base"
	"github.com/micro-plat/sso/mgrapi/services/function"
	"github.com/micro-plat/sso/mgrapi/services/image"
	"github.com/micro-plat/sso/mgrapi/services/member"
	"github.com/micro-plat/sso/mgrapi/services/menu"
	"github.com/micro-plat/sso/mgrapi/services/role"
	"github.com/micro-plat/sso/mgrapi/services/system"
	"github.com/micro-plat/sso/mgrapi/services/user"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) init() {
	//初始化
	r.Initializing(func(c component.IContainer) error {
		var conf model.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		model.SaveConf(c, &conf)
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

		logic.Set(c) //保存全局菜单变量
		return nil
	})

	r.Micro("/sso/login", member.NewLoginHandler)               //用户名密码登录
	r.Micro("/sso/menu", menu.NewMenuHandler)                   //系统菜单相关接口
	r.Micro("/sso/ident", system.NewSystemIdentHandler)         //系统信息获取
	r.Micro("/sso/member", member.NewQueryHandler)              //查询登录用户信息
	r.Micro("/sso/user/changepwd", user.NewUserPasswordHandler) // 修改密码
	r.Micro("/sso/base", base.NewBaseUserHandler)
	r.Micro("/sso/user", user.NewUserHandler)                       //用户相关接口
	r.Micro("/sso/auth", role.NewRoleAuthHandler)                   //权限管理
	r.Micro("/sso/role", role.NewRoleHandler)                       //角色管理相关接口
	r.Micro("/sso/sys/manage", system.NewSystemHandler)             //系统管理相关接口
	r.Micro("/sso/sys/func", function.NewSystemFuncHandler)         //系统功能相关接口
	r.Micro("/sso/img/upload", image.NewImageHandler("static/img")) //图片上传
}
