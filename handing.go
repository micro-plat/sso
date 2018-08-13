package main

import (
	"github.com/micro-plat/hydra/component"
	"fmt"

	"github.com/micro-plat/hydra/context"
	mem "github.com/micro-plat/sso/modules/member"
	xmenu "github.com/micro-plat/sso/modules/menu"
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) handing() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {		

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
		var m mem.LoginState
		if err = ctx.Request.GetJWT(&m); err != nil {
			return context.NewError(context.ERR_FORBIDDEN, err)
		}
		if err = mem.Save(ctx, &m); err != nil {
			return err
		}

		//检查用户权限
		tags := r.GetTags(ctx.Service)
		menu := xmenu.Get(ctx.GetContainer().(component.IContainer))
		for _, tag := range tags {
			if tag == "*" {
				return nil
			}
			if err = menu.Verify(m.UserID, m.SystemID, tag,ctx.Request.GetMethod()); err == nil {
				return nil
			}
		}
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Sprintf("没有权限:%v", tags))
	})
}
