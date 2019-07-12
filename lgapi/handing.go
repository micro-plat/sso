package main

import (
	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/model"
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) handling() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {

		//是否配置jwt
		jwt, err := ctx.Request.GetJWTConfig()
		if err != nil {
			return err
		}

		for _, u := range jwt.Exclude {
			if u == ctx.Service {
				return nil
			}
		}

		var m model.LoginState
		if err = ctx.Request.GetJWT(&m); err != nil {
			return context.NewError(context.ERR_FORBIDDEN, err)
		}

		if err = member.Save(ctx, &m); err != nil {
			return err
		}
		return nil
	})
}
