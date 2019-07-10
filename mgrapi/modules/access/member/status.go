package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

const maxErrorCnt = 5

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

// Query Query
func Query(ctx *context.Context, container component.IContainer) *model.LoginState {
	m := &model.LoginState{}
	err := ctx.Request.GetJWT(m)
	if m.UserName != "" && err == nil {
		return m
	}
	if err := ctx.Request.Check("code"); err != nil {
		return nil
	}
	codeMemberLib := NewCodeMember(container)
	m, err = codeMemberLib.Query(ctx.Request.GetString("code"))
	if err != nil {
		return nil
	}
	return m
}
