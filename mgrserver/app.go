package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	_ "github.com/micro-plat/hydra/components/caches/cache/gocache"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	_ "github.com/micro-plat/hydra/components/queues/mq/redis"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/sso/common/config"
	"github.com/micro-plat/sso/common/dds"
	_ "github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/sqls/mysql"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	cmodel "github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/base"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/function"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/image"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/permission"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/role"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/system"
	"github.com/micro-plat/sso/mgrserver/mgrapi/services/user"

	ssoSdk "github.com/micro-plat/sso/sdk/sso"
)

var App = hydra.NewApp(
	hydra.WithPlatName("sso_new", "新版sso"),
	hydra.WithSystemName("mgrserver", "sso单点登录管理系统"),
	hydra.WithUsage("单点登录管理系统"),
	hydra.WithServerTypes(http.Web),
	hydra.WithClusterName("prod"))

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	install()

	dds.Bind(App)
	ssoSdk.Bind(App)

	//每个请求执行前执行
	App.OnHandleExecuting(func(ctx hydra.IContext) (rt interface{}) {
		ctx.Log().Info("handling.....")
		//验证jwt并缓存登录用户信息
		if err := ssoSdk.CheckAndSetMember(ctx); err != nil {
			return err
		}
		return nil
	})

	//启动事检查配置是否正确
	App.OnStarting(func(appconf app.IAPPConf) error {
		var vueconf cmodel.VueConf
		if _, err := appconf.GetServerConf().GetSubObject("vueconf", &vueconf); err != nil {
			return err
		}

		if err := vueconf.Valid(); err != nil {
			return err
		}

		if err := cmodel.SaveVueConf(&vueconf); err != nil {
			return err
		}

		//检查配置信息
		var conf model.Conf
		_, err := appconf.GetServerConf().GetSubObject("app", &conf)
		if err != nil {
			return err
		}

		if err := model.SaveConf(&conf); err != nil {
			return err
		}

		_, err = components.Def.DB().GetDB()
		if err != nil {
			return err
		}

		_, err = components.Def.Cache().GetCache("redis")
		if err != nil {
			return err
		}

		if err := ssoSdk.BindConfig(conf.SsoApiHost, conf.Ident, conf.Secret); err != nil {
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

	//vue config
	App.Micro("/config/vue", config.VueHandler)
}
