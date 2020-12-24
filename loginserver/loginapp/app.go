package loginapp

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/loginserver/webserver/services/login"
	"github.com/micro-plat/sso/loginserver/webserver/services/member"
	"github.com/micro-plat/sso/loginserver/webserver/services/system"

	apilogin "github.com/micro-plat/sso/loginserver/apiserver/services/login"
	apimember "github.com/micro-plat/sso/loginserver/apiserver/services/member"
	"github.com/micro-plat/sso/loginserver/apiserver/services/password"
	"github.com/micro-plat/sso/loginserver/apiserver/services/permission"
	apisystem "github.com/micro-plat/sso/loginserver/apiserver/services/system"
	"github.com/micro-plat/sso/loginserver/apiserver/services/user"
)

//17sup_v2_sso   hbs_sso
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

	//服务启动检查
	hydra.OnReady(func() error {
		_, err := components.Def.DB().GetDB("db")
		if err != nil {
			return err
		}

		_, err = components.Def.Cache().GetCache("redis")
		if err != nil {
			return err
		}

		var conf model.Conf
		varConf, err := app.Cache.GetVarConf()
		if err != nil {
			return err
		}

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
		return nil
	})

	//web接口
	App.API("/login/check", login.NewLoginCheckHandler)       //验证用户是否已登录
	App.API("/member/login", login.NewLoginHandler)           //用户登录相关
	App.API("/member/bind", member.NewBindWxHandler)          //绑定微信
	App.API("/member/changepwd", member.NewChangePwdHandler)  //修改密码
	App.API("/member/refresh", member.NewRefleshTokenHandler) //刷新用户token
	App.API("/member/sendcode", member.NewSendCodeHandler)    //发送验证码
	App.API("/member/system/get", member.NewUserSysHandler)   //获取用户可进的系统信息
	App.API("/system/config/get", system.NewSystemHandler)    //获取系统的一些配置信息
	//web接口

	//api 接口
	App.Micro("/member/menu/get", apimember.NewMenuHandler)            //获取用户菜单数据
	App.Micro("/member/tags/get", apimember.NewTagHandler)             //获取用户有权限的tag数据
	App.Micro("/member/info/get", apimember.NewMemberHandler)          //获取用户信息
	App.Micro("/role/user/get", apimember.NewRoleHandler)              //获取角色下的所有用户
	App.Micro("/member/system/get", apimember.NewMemberSysHandler)     //获取用户可用的子系统
	App.Micro("/member/all/get", apimember.NewMemberGetAllHandler)     //获取所有用户信息
	App.Micro("/system/info/get", apisystem.NewInfoHandler)            //获取子系统信息
	App.Micro("/login/auth", apilogin.NewAuthorizeHandler)             //用户跳转登录后的认证(不是用户名密码登录)
	App.Micro("/permission/config", permission.NewDataPerssionHandler) //【数据权限】相关接口
	App.Micro("/member/forget/password", password.NewPasswordHandler)  // 忘记密码再修改密码

	//以下接口是为sass系统使用
	App.Micro("/user", user.NewUserHandler)                     //用户相关接口
	App.Micro("/verifycode/get", apilogin.NewVerifyCodeHandler) //生成图片验证码(这个现在没用,以后可能会用到)
	App.Micro("/check_sign", apilogin.NewCheckSignHandler)      //检查签名
	//api 接口
}
