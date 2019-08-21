package main

import (
	"github.com/micro-plat/hydra/context"
	mem "github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/member"

	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) handling() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {

		//跳过jwt排除的请求
		if skip, err := ctx.Request.SkipJWTExclude(); err != nil || skip {
			return err
		}

		//缓存用户信息
		var m model.LoginState
		if err := ctx.Request.GetJWT(&m); err != nil {
			return context.NewError(context.ERR_FORBIDDEN, err)
		}

		if err := mem.Save(ctx, &m); err != nil {
			return err
		}
		return nil
	})
}
