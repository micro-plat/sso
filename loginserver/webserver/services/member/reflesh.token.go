package member

import (
	"github.com/micro-plat/hydra"
)

//RefleshTokenHandler 用户对象
type RefleshTokenHandler struct {
}

//NewRefleshTokenHandler 用户
func NewRefleshTokenHandler() (u *RefleshTokenHandler) {
	return &RefleshTokenHandler{}
}

//Handle 刷新token 这个接口只是为了刷新sso登录用户的jwt, jwt刷新在框架就做了
func (u *RefleshTokenHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------刷新token---------")

	return "success"
}
