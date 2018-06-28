package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

type WxcodeHandler struct {
	container component.IContainer
	wx        member.IWxcode
}

func NewWxcodeHandler(container component.IContainer) (u *WxcodeHandler) {
	return &WxcodeHandler{
		container: container,
		wx:        member.NewWxcode(container),
	}
}

//Handle 发送微信验证码
func (u *WxcodeHandler) Handle(ctx *context.Context) (r interface{}) {
	//检查输入参数
	if err := ctx.Request.Check("username", "sysid"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	_, err := u.wx.GetWXCode(ctx.Request.GetString("username"), ctx.Request.GetString("sysid"))
	if err != nil {
		return err
	}
	return "success"
}
