package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
	"github.com/micro-plat/sso/modules/app"
	mem "github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/services/member"
	"github.com/micro-plat/sso/services/menu"
	"github.com/micro-plat/sso/services/system"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func bindConf(app *hydra.MicroApp) {
	app.Conf.API.SetMainConf(`{"address":":9091"}`)
	app.Conf.WS.SetSubConf("app", `
		{
			"qrcode-login-check-url":"http://192.168.5.71:8090"
		}		
		`)
	app.Conf.API.SetSubConf("app", `
			{
				"qrcode-login-check-url":"http://192.168.5.71:8090"
			}			
			`)
	app.Conf.API.SetSubConf("header", `
				{
					"Access-Control-Allow-Origin": "*", 
					"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS", 
					"Access-Control-Allow-Headers": "__jwt__", 
					"Access-Control-Allow-Credentials": "true"
				}
			`)

	app.Conf.API.SetSubConf("auth", `
		{
			"jwt": {
				"exclude": ["/sso/login","/sso/wxcode/get"],
				"expireAt": 36000,
				"mode": "HS512",
				"name": "__jwt__",
				"redirect":"/sso/login",
				"secret": "12345678"
			}
		}
		`)
	app.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"ora",
			"connString":"sso/123456@orcl136",
			"maxOpen":10,
			"maxIdle":1,
			"lifeTime":10		
	}`)

	app.Conf.Plat.SetVarConf("cache", "cache", `
		{
			"proto":"redis",
			"addrs":[
					"192.168.0.110:6379",
					"192.168.0.122:6379",
					"192.168.0.134:6379",
					"192.168.0.122:6380",
					"192.168.0.110:6380",
					"192.168.0.134:6380"
			],
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}
		
		`)

}

//bind 检查应用程序配置文件，并根据配置初始化服务
func bind(r *hydra.MicroApp) {
	bindConf(r)

	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {
		jwt, err := ctx.Request.GetJWTConfig() //获取jwt配置
		if err != nil {
			return err
		}
		for _, u := range jwt.Exclude { //排除指定请求
			if u == ctx.Service {
				return nil
			}
		}
		//检查jwt配置，并使用member中提供的函数缓存login信息到context中
		var m mem.LoginState
		if err := ctx.Request.GetJWT(&m); err != nil {
			return context.NewError(context.ERR_FORBIDDEN, err)
		}
		return mem.Save(ctx, &m)
	})

	//初始化
	r.Initializing(func(c component.IContainer) error {
		var conf app.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}

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
		return nil
	})
	r.Micro("/sso/login", member.NewLoginHandler)    //用户登录，登录后自动转跳到系统配置地址
	r.Micro("/sso/sys/get", system.NewSystemHandler) //获取系统信息
	r.Micro("/sso/menu/get", menu.NewMenuHandler)    //获取用户菜单
	r.Micro("/sso/popular", menu.NewPopularHandler)  //获取用户常用菜单

	r.WS("/qrcode/query", member.NewQRCodeHandler) //登录二维码获取

	r.Micro("/sso/wxcode/get", member.NewWxcodeHandler) //获取已发送的微信验证码
	r.Micro("/sso/login/check", member.NewCheckHandler) //用户登录状态检查，检查用户jwt是否有效
	r.Micro("/sso/member/get", member.NewGetHandler)    //获取用户信息（不包括角色信息）

	r.Micro("/sso/menu/verify", menu.NewVerifyHandler) //检查用户菜单权限
}
