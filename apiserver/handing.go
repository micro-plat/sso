package main

import (
	"github.com/micro-plat/hydra/context"
	// "github.com/micro-plat/sso/apiserver/modules/access/member"
	// xmenu "github.com/micro-plat/sso/apiserver/modules/logic"
	// "github.com/micro-plat/sso/apiserver/modules/model"
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) handing() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {
		return nil

		/*
			//是否配置jwt
			jwt, err := ctx.Request.GetJWTConfig() //获取jwt配置
			if err != nil {
				return err
			}
			for _, u := range jwt.Exclude { //排除指定请求
				if u == ctx.Service {
					return nil
				}
			}

			//缓存用户信息
			var m model.LoginState
			if err = ctx.Request.GetJWT(&m); err != nil {
				return context.NewError(context.ERR_FORBIDDEN, err)
			}
			if err = member.Save(ctx, &m); err != nil {
				return err
			}
			//检查用户权限
			tags := r.GetTags(ctx.Service)
			menu := xmenu.Get(ctx.GetContainer().(component.IContainer))
			for _, tag := range tags {
				if tag == "*" {
					return nil
				}
				if err = menu.Verify(m.UserID, m.SystemID, tag, ctx.Request.GetMethod()); err == nil {
					return nil
				}
			}
			return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Sprintf("没有权限:%v", tags))
		*/
	})
}
