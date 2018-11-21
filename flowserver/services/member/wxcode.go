package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

type WxcodeHandler struct {
	container  component.IContainer
	wx         member.IWxcode
	appid      string
	secret     string
	serverAddr string
}

func NewWxcodeHandler(appid string, secret string, serverAddr string) func(c component.IContainer) (u *WxcodeHandler) {
	return func(c component.IContainer) (u *WxcodeHandler) {
		return &WxcodeHandler{
			container:  c,
			wx:         member.NewWxcode(c),
			appid:      appid,
			secret:     secret,
			serverAddr: serverAddr,
		}
	}
}

//Handle 发送微信验证码
func (u *WxcodeHandler) Handle(ctx *context.Context) (r interface{}) {
	//检查输入参数
	if err := ctx.Request.Check("username", "ident"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	code := u.wx.GetWXCode()
	if err := u.wx.Send(
		ctx.Request.GetString("username"),
		ctx.Request.GetString("ident"),
		u.appid,
		u.secret,
		u.serverAddr,
		code); err != nil {
		return err
	}

	return "success"
}
