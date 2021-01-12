package login

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/loginserver/sdkapi/modules/logic"
)

//VerifyCodeHandler is 生成登录图片验证码
type VerifyCodeHandler struct {
	m logic.IMemberLogic
}

//NewVerifyCodeHandler is
func NewVerifyCodeHandler() (u *VerifyCodeHandler) {
	return &VerifyCodeHandler{
		m: logic.NewMemberLogic(),
	}
}

//Handle 生成登录图片验证码 并存与redis中
func (u *VerifyCodeHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------生成登录图片验证码---------")

	if err := ctx.Request().Check("user_name"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("1. 生成登录图片验证码")
	info, err := u.m.GenerateVerifyCode(ctx.Request().GetString("user_name"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3. 返回图片base64编码")
	return info
}
