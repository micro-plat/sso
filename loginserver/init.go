package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/loginserver/services/login"
	"github.com/micro-plat/sso/loginserver/services/member"
	"github.com/micro-plat/sso/loginserver/services/system"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	//设置配置参数
	install()

	//挂载请求处理函数
	handling()

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
		appConf, err := app.Cache.GetAPPConf(http.Web)
		if err != nil {
			return err
		}
		_, err = appConf.GetServerConf().GetSubObject("app", &conf)
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

	App.API("/login/check", login.NewLoginCheckHandler)       //验证用户是否已登录
	App.API("/member/login", login.NewLoginHandler)           //用户登录相关
	App.API("/member/bind", member.NewBindWxHandler)          //绑定微信
	App.API("/member/changepwd", member.NewChangePwdHandler)  //修改密码
	App.API("/member/refresh", member.NewRefleshTokenHandler) //刷新用户token
	App.API("/member/sendcode", member.NewSendCodeHandler)    //发送验证码
	App.API("/member/system/get", member.NewUserSysHandler)   //获取用户可进的系统信息
	App.API("/system/config/get", system.NewSystemHandler)    //获取系统的一些配置信息
}
