package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
	mem "github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/services/member"
)

//AppConf 应用程序配置
type AppConf struct {
}

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func bindConf(app *hydra.MicroApp) {
	app.Conf.API.SetSubConf("auth", `
		{
			"jwt": {
				"exclude": ["/member/login","/member/check"],
				"expireAt": 36000,
				"mode": "HS512",
				"name": "__jwt__",
				"secret": "12345678"
			}
		}
		`)

}

//bind 检查应用程序配置文件，并根据配置初始化服务
func bind(r *hydra.MicroApp) {
	bindConf(r)

	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {
		if r.IsDebug {
			return
		}
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
		var m mem.Member
		if err := ctx.Request.GetJWT(&m); err != nil {
			return context.NewError(context.ERR_FORBIDDEN, err)
		}
		mem.Save(ctx, &m)
		return nil
	})

	//初始化
	r.Initializing(func(c component.IContainer) error {

		//获取配置
		var conf AppConf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		if b, err := govalidator.ValidateStruct(&conf); !b {
			return fmt.Errorf("app 配置文件有误:%v", err)
		}

		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}

		//检查cache配置是否正确
		if _, err := c.GetCache(); err != nil {
			return err
		}
		return nil
	})
	r.Micro("/member/login", member.NewLoginHandler)
	r.Micro("/member/check", member.NewCheckHandler)
	r.Micro("/member/menu", member.NewMenuHandler)
}
