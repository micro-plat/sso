package login

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
)

//VerifyCodeHandler is 生成登录图片验证码
type VerifyCodeHandler struct {
	container component.IContainer
	m         logic.IMemberLogic
}

//NewVerifyCodeHandler is
func NewVerifyCodeHandler(container component.IContainer) (u *VerifyCodeHandler) {
	return &VerifyCodeHandler{
		container: container,
		m:         logic.NewMemberLogic(container),
	}
}

//Handle 生成登录图片验证码 并存与redis中
func (u *VerifyCodeHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------生成登录图片验证码---------")

	if err := ctx.Request.Check("user_name"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("1. 生成登录图片验证码")
	info, err := u.m.GenerateVerifyCode(ctx.Request.GetString("user_name"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 返回图片base64编码")
	return info
}
