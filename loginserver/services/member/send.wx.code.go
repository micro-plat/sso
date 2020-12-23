// +build !sms

package member

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/common/service"
	"github.com/micro-plat/sso/loginserver/modules/logic"
)

//SendCodeHandler 发送微信验证码
type SendCodeHandler struct {
	mem logic.IMemberLogic
}

//NewSendCodeHandler 发送微信验证码
func NewSendCodeHandler() (u *SendCodeHandler) {
	return &SendCodeHandler{
		mem: logic.NewMemberLogic(),
	}
}

//Handle 发送微信验证码
func (u *SendCodeHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------发送微信验证码---------")

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("username"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2: 验证用户状态以及是否绑定了微信账户")
	openID, err := u.mem.ValidUserInfo(ctx.Request().GetString("username"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 发送微信验证码")
	if err := service.SendWxVerifyCode(ctx.Request().GetString("username"), openID, ctx.Request().GetString("ident")); err != nil {
		return err
	}

	return "success"
}
