package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/common/service"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/logic"
)

//SendCodeHandler 发送微信验证码
type SendCodeHandler struct {
	c   component.IContainer
	mem logic.IMemberLogic
}

//NewSendCodeHandler 发送微信验证码
func NewSendCodeHandler(container component.IContainer) (u *SendCodeHandler) {
	return &SendCodeHandler{
		c:   container,
		mem: logic.NewMemberLogic(container),
	}
}

//Handle 发送微信验证码
func (u *SendCodeHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------发送微信验证码---------")

	ctx.Log.Info("1: 验证参数")
	if err := ctx.Request.Check("username"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2: 验证用户状态以及是否绑定了微信账户")
	openID, err := u.mem.ValidUserInfo(ctx.Request.GetString("username"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3: 发送微信验证码")
	if err := service.SendWxVerifyCode(u.c, ctx.Request.GetString("username"), openID, ctx.Request.GetString("ident")); err != nil {
		return err
	}

	return "success"
}
