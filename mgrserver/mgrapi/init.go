package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/base"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/function"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/image"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/permission"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/role"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/system"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/user"
	ssoSdk "github.com/micro-plat/sso/sdk/sso"
	"gitlab.100bm.cn/micro-plat/dds/dds"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) init() {
	r.install()
	r.handling()

	r.Initializing(func(c component.IContainer) error {
		//检查配置信息
		var conf model.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		if err := conf.Valid(); err != nil {
			return err
		}
		model.SaveConf(c, &conf)

		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}

		//检查缓存配置是否正确
		if _, err := c.GetCache(); err != nil {
			return err
		}

		dds.Bind(r.MicroApp, "db")
		if err := ssoSdk.Bind(r.MicroApp, conf.SsoApiHost, conf.Ident, conf.Secret); err != nil {
			return err
		}

		return nil
	})

	r.Micro("/base", base.NewBaseUserHandler)                          //基础数据
	r.Micro("/user", user.NewUserHandler)                              //用户相关接口
	r.Micro("/auth", role.NewRoleAuthHandler)                          //菜单权限管理
	r.Micro("/role", role.NewRoleHandler)                              //角色管理相关接口
	r.Micro("/system/info", system.NewSystemHandler)                   //系统管理相关接口
	r.Micro("/system/menu", system.NewSystemMenuHandler)               //系统菜单管理相关接口
	r.Micro("/system/func", function.NewSystemFuncHandler)             //系统功能相关接口
	r.Micro("/system/permission", permission.NewDataPermissionHandler) //数据权限功能相关接口
	r.Micro("/auth/permission", permission.NewAuthPermissionHandler)   //数据权限管理
	r.Micro("/image/upload", image.NewImageHandler("../image"))        //图片上传
}
