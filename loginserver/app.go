package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	_ "github.com/micro-plat/hydra/components/caches/cache/gocache"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	_ "github.com/micro-plat/hydra/components/queues/mq/redis"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/sso/common/module/model"
	cmember "github.com/micro-plat/sso/loginserver/loginapi/modules/access/member"
	cmodel "github.com/micro-plat/sso/loginserver/loginapi/modules/model"
	"github.com/micro-plat/sso/loginserver/loginapi/services/login"
	"github.com/micro-plat/sso/loginserver/loginapi/services/member"
	"github.com/micro-plat/sso/loginserver/loginapi/services/system"

	apilogin "github.com/micro-plat/sso/loginserver/apiserver/services/login"
	apimember "github.com/micro-plat/sso/loginserver/apiserver/services/member"
	"github.com/micro-plat/sso/loginserver/apiserver/services/password"
	"github.com/micro-plat/sso/loginserver/apiserver/services/permission"
	apisystem "github.com/micro-plat/sso/loginserver/apiserver/services/system"
	"github.com/micro-plat/sso/loginserver/apiserver/services/user"
	"github.com/micro-plat/sso/loginserver/apiserver/services/vueconf"
)

var App = hydra.NewApp(
	hydra.WithPlatName("sso_new", "新版sso"),
	hydra.WithSystemName("loginserver", "sso单点登录服务"),
	hydra.WithUsage("单点登录服务"),
	hydra.WithServerTypes(http.Web, http.API),
	hydra.WithClusterName("prod"))

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	//设置配置参数
	install()

	//启动时参数配置检查
	App.OnStarting(func(appConf app.IAPPConf) error {
		_, err := components.Def.DB().GetDB()
		if err != nil {
			return err
		}

		_, err = components.Def.Cache().GetCache("redis")
		if err != nil {
			return err
		}

		var conf model.Conf
		varConf := appConf.GetVarConf()
		_, err = varConf.GetObject("loginconf", "app", &conf)
		if err != nil {
			return err
		}

		if err := conf.Valid(); err != nil {
			return err
		}

		if err := model.SaveConf(&conf); err != nil {
			return err
		}

		var vueconf cmodel.VueConf
		if _, err = appConf.GetServerConf().GetSubObject("vueconf", &vueconf); err != nil {
			return err
		}

		if err := vueconf.Valid(); err != nil {
			return err
		}

		if err := cmodel.SaveConf(&vueconf); err != nil {
			return err
		}
		return nil
	})

	App.OnHandleExecuting(func(ctx hydra.IContext) interface{} {
		if ctx.User().Auth().Request() != nil {
			var out cmodel.LoginState
			if err := ctx.User().Auth().Bind(&out); err != nil {
				return err
			}
			cmember.Save(ctx, &out)
		}
		return nil
	})

	//对web后台接口添加全局后处理钩子函数   如果存在jwt的登录信息,就设置到response中返回给前端
	App.OnHandleExecuted(func(ctx hydra.IContext) interface{} {
		auth := ctx.User().Auth().Request()
		if auth != nil {
			ctx.User().Auth().Response(auth)
		}
		return nil
	})

	//web接口
	App.Web("/mgrweb/login/check", login.NewLoginCheckHandler)       //验证用户是否已登录
	App.Web("/mgrweb/member/login", login.NewLoginHandler)           //用户登录相关
	App.Web("/mgrweb/member/bind", member.NewBindWxHandler)          //绑定微信
	App.Web("/mgrweb/member/changepwd", member.NewChangePwdHandler)  //修改密码
	App.Web("/mgrweb/member/refresh", member.NewRefleshTokenHandler) //刷新用户token
	App.Web("/mgrweb/member/sendcode", member.NewSendCodeHandler)    //发送验证码
	App.Web("/mgrweb/member/system/get", member.NewUserSysHandler)   //获取用户可进的系统信息
	App.Web("/mgrweb/system/config/get", system.NewSystemHandler)    //获取系统的一些配置信息
	App.Web("/vue/config/get", vueconf.NewGetVueConfHandler)         //获取前端页面配置

	//web接口

	//api 接口
	App.Micro("/member/menu/get", apimember.NewMenuHandler)            //获取用户菜单数据
	App.Micro("/member/tags/get", apimember.NewTagHandler)             //获取用户有权限的tag数据
	App.Micro("/member/info/get", apimember.NewMemberHandler)          //获取用户信息
	App.Micro("/member/forget/password", password.NewPasswordHandler)  // 忘记密码再修改密码
	App.Micro("/member/system/get", apimember.NewMemberSysHandler)     //获取用户可用的子系统
	App.Micro("/member/all/get", apimember.NewMemberGetAllHandler)     //获取所有用户信息
	App.Micro("/role/user/get", apimember.NewRoleHandler)              //获取角色下的所有用户
	App.Micro("/system/info/get", apisystem.NewInfoHandler)            //获取子系统信息
	App.Micro("/login/auth", apilogin.NewAuthorizeHandler)             //用户跳转登录后的认证(不是用户名密码登录)
	App.Micro("/permission/config", permission.NewDataPerssionHandler) //【数据权限】相关接口

	//以下接口是为sass系统使用
	App.Micro("/user", user.NewUserHandler)                     //用户相关接口
	App.Micro("/verifycode/get", apilogin.NewVerifyCodeHandler) //生成图片验证码(这个现在没用,以后可能会用到)
	App.Micro("/check_sign", apilogin.NewCheckSignHandler)      //检查签名
	//api 接口
}
