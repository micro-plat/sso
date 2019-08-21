package member

import (
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

//Save 保存member信息
func Save(ctx *context.Context, m *model.LoginState) error {
	ctx.Meta.Set("login-state", m)
	return nil
}

//Get 获取member信息
func Get(ctx *context.Context) *model.LoginState {
	v, _ := ctx.Meta.Get("login-state")
	if v == nil {
		return nil
	}
	return v.(*model.LoginState)
}
