package member

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/loginserver/webserver/modules/model"
)

const maxErrorCnt = 5

//Save 保存member信息
func Save(ctx hydra.IContext, m *model.LoginState) error {
	ctx.Meta().SetValue("login-state", m)
	return nil
}

//Get 获取member信息
func Get(ctx hydra.IContext) *model.LoginState {
	v, _ := ctx.Meta().Get("login-state")
	if v == nil {
		return nil
	}
	return v.(*model.LoginState)
}
