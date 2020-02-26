package sso

import (
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
)

//GetMember 获取登录用户信息
func GetMember(ctx *context.Context) *model.LoginState {
	v, _ := ctx.Meta.Get("login-state")
	if v == nil {
		return nil
	}
	return v.(*model.LoginState)
}

//CheckAndSetMember 验证jwt同时保存用户登录信息
func CheckAndSetMember(ctx *context.Context) error {
	if skip, err := ctx.Request.SkipJWTExclude(); err != nil || skip {
		return err
	}

	//验证用户是否登录
	var m model.LoginState
	if err := ctx.Request.GetJWT(&m); err != nil {
		return context.NewError(context.ERR_FORBIDDEN, err)
	}

	//保存登录用户信息
	ctx.Meta.Set("login-state", &m)
	return nil
}
