package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//RefleshTokenHandler 用户对象
type RefleshTokenHandler struct {
	c component.IContainer
}

//NewRefleshTokenHandler 用户
func NewRefleshTokenHandler(container component.IContainer) (u *RefleshTokenHandler) {
	return &RefleshTokenHandler{
		c: container,
	}
}

//Handle 刷新token 这个接口只是为了刷新sso登录用户的jwt, jwt刷新在框架就做了
func (u *RefleshTokenHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------刷新token---------")

	return "success"
}
