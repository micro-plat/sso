package main

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/sso/mgrserver/modules/model"
	"github.com/micro-plat/sso/mgrserver/services/base"
	"github.com/micro-plat/sso/mgrserver/services/function"
	"github.com/micro-plat/sso/mgrserver/services/image"
	"github.com/micro-plat/sso/mgrserver/services/permission"
	"github.com/micro-plat/sso/mgrserver/services/role"
	"github.com/micro-plat/sso/mgrserver/services/system"
	"github.com/micro-plat/sso/mgrserver/services/user"
	ssoSdk "github.com/micro-plat/sso/sdk/sso"
	"gitlab.100bm.cn/micro-plat/dds/dds"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	install()

	handling()

	hydra.OnReady(func() error {
		//检查配置信息
		var conf model.Conf
		appConf, err := app.Cache.GetAPPConf(http.Web)
		if err != nil {
			return err
		}
		_, err = appConf.GetServerConf().GetSubObject("app", &conf)
		if err != nil {
			return err
		}
		if err := model.SaveConf(&conf); err != nil {
			return err
		}

		_, err := components.Def.DB().GetDB("db")
		if err != nil {
			return err
		}

		_, err = components.Def.Cache().GetCache("redis")
		if err != nil {
			return err
		}

		dds.Bind(App, "db")
		if err := ssoSdk.Bind(App, conf.SsoApiHost, conf.Ident, conf.Secret); err != nil {
			return err
		}

		return nil
	})

	App.Micro("/base", base.NewBaseUserHandler)                          //基础数据
	App.Micro("/user", user.NewUserHandler)                              //用户相关接口
	App.Micro("/auth", role.NewRoleAuthHandler)                          //菜单权限管理
	App.Micro("/role", role.NewRoleHandler)                              //角色管理相关接口
	App.Micro("/system/info", system.NewSystemHandler)                   //系统管理相关接口
	App.Micro("/system/menu", system.NewSystemMenuHandler)               //系统菜单管理相关接口
	App.Micro("/system/func", function.NewSystemFuncHandler)             //系统功能相关接口
	App.Micro("/system/permission", permission.NewDataPermissionHandler) //数据权限功能相关接口
	App.Micro("/auth/permission", permission.NewAuthPermissionHandler)   //数据权限管理
	App.Micro("/image/upload", image.NewImageHandler("../image"))        //图片上传
}
