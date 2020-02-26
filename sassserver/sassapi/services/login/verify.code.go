package login

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/logic"
)

//VerifyCodeHandler 用户登录
type VerifyCodeHandler struct {
	container component.IContainer
	subLib    logic.ILoginLogic
}

//NewVerifyCodeHandler new
func NewVerifyCodeHandler(container component.IContainer) (u *VerifyCodeHandler) {
	return &VerifyCodeHandler{
		container: container,
		subLib:    logic.NewLoginLogic(container),
	}
}

//Handle 获取验证码
func (u *VerifyCodeHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------获取登录验证码---------")

	ctx.Log.Info("1:参数验证")
	if err := ctx.Request.Check("mobile"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "手机号不能为空")
	}

	ctx.Log.Info("2: 生成验证码")
	verfifyCodePic, err := u.subLib.GenerateVerifyCode(ctx.Request.GetString("mobile"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3: 返回图片base64编码")
	return verfifyCodePic
}
