package main

import (
	"github.com/lib4dev/vcs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	_ "github.com/micro-plat/hydra/components/caches/cache/gocache"
	_ "github.com/micro-plat/hydra/components/caches/cache/redis"
	_ "github.com/micro-plat/hydra/components/queues/mq/redis"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers/http"
	cmember "github.com/micro-plat/sso/loginserver/loginapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/model"
	cmodel "github.com/micro-plat/sso/loginserver/loginapi/modules/model"
	"github.com/micro-plat/sso/loginserver/loginapi/services/login"
	"github.com/micro-plat/sso/loginserver/loginapi/services/member"
	"github.com/micro-plat/sso/loginserver/loginapi/services/system"

	apilogin "github.com/micro-plat/sso/loginserver/srvapi/services/login"
	apimember "github.com/micro-plat/sso/loginserver/srvapi/services/member"
	"github.com/micro-plat/sso/loginserver/srvapi/services/password"
	"github.com/micro-plat/sso/loginserver/srvapi/services/permission"
	apisystem "github.com/micro-plat/sso/loginserver/srvapi/services/system"
	"github.com/micro-plat/sso/loginserver/srvapi/services/user"
	_ "github.com/micro-plat/sso/sso"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {

	//web接口
	registryAPI()

	//启动时参数配置检查
	App.OnStarting(func(appConf app.IAPPConf) error {
		//检查DB
		if _, err := components.Def.DB().GetDB(); err != nil {
			return err
		}
		//检查Cache
		if _, err := components.Def.Cache().GetCache("redis"); err != nil {
			return err
		}
		//检查登录配置
		if err := checkLoginConf(appConf); err != nil {
			return err
		}

		return nil
	}, http.API)

	App.OnHandleExecuting(func(ctx hydra.IContext) interface{} {
		if ctx.User().Auth().Request() != nil {
			out := &cmodel.LoginState{}
			if err := ctx.User().Auth().Bind(out); err != nil {
				return err
			}
			cmember.Save(ctx, out)
		}
		return nil
	})

}

func checkLoginConf(appConf app.IAPPConf) (err error) {
	var loginCfg model.LoginConf
	varConf := appConf.GetVarConf()
	_, err = varConf.GetObject("loginconf", "app", &loginCfg)
	if err != nil {
		return err
	}

	if err := loginCfg.Valid(); err != nil {
		return err
	}

	if err := model.SaveLoginConf(&loginCfg); err != nil {
		return err
	}

	vcs.SetConfig(vcs.WithCacheConfig("redis", "http"), vcs.WithSmsSendUrl(loginCfg.SmsSendURL))

	return
}

func registryAPI() {
	//web接口
	App.Web("/loginweb/login/check", login.NewLoginCheckHandler)       //验证用户是否已登录
	App.Web("/loginweb/logout", login.NewLogoutHandler)                //验证用户退出登录
	App.Web("/loginweb/member/login", login.NewLoginHandler)           //用户登录相关
	App.Web("/loginweb/member/bind", member.NewBindWxHandler)          //绑定微信
	App.Web("/loginweb/member/changepwd", member.NewChangePwdHandler)  //修改密码
	App.Web("/loginweb/member/refresh", member.NewRefleshTokenHandler) //刷新用户token
	App.Web("/loginweb/member/sendcode", member.NewSendCodeHandler)    //发送验证码
	App.Web("/loginweb/member/system/get", member.NewUserSysHandler)   //获取用户可进的系统信息
	App.Web("/loginweb/system/config/get", system.NewSystemHandler)    //获取系统的一些配置信息

	//api 接口
	App.Micro("/api/menu/get", apimember.NewMenuHandler)                   //获取用户菜单数据
	App.Micro("/api/tags/get", apimember.NewTagHandler)                    //获取用户有权限的tag数据
	App.Micro("/api/role/tags", apimember.NewRoleTagsHandler)              //获取用户有权限的tag数据
	App.Micro("/api/role/get", apimember.NewRoleHandler)                   //获取角色下的所有用户
	App.Micro("/api/user/info/get", apimember.NewMemberHandler)            //获取用户信息
	App.Micro("/api/forget/password", password.NewPasswordHandler)         // 忘记密码再修改密码
	App.Micro("/api/user/system/list", apimember.NewMemberSysHandler)      //获取用户可用的子系统
	App.Micro("/api/user/all/get", apimember.NewMemberGetAllHandler)       //获取所有用户信息
	App.Micro("/api/system/info/get", apisystem.NewInfoHandler)            //获取子系统信息
	App.Micro("/api/login/auth", apilogin.NewAuthorizeHandler)             //用户跳转登录后的认证(不是用户名密码登录)
	App.Micro("/api/permission/config", permission.NewDataPerssionHandler) //【数据权限】相关接口

	//以下接口是为sass系统使用
	App.Micro("/api/user", user.NewUserHandler)                     //用户相关接口
	App.Micro("/api/verifycode/get", apilogin.NewVerifyCodeHandler) //生成图片验证码(这个现在没用,以后可能会用到)
	App.Micro("/check_sign", apilogin.NewCheckSignHandler)          //检查签名

}
