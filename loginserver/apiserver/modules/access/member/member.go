package member

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/model"
)

//Get 获取member信息
func Get(ctx hydra.IContext) *model.LoginState {
	var s = model.LoginState{}
	if err := ctx.User().Auth().Bind(&s); err != nil {
		return nil
	}
	return &s
}
