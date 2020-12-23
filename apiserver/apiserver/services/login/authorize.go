package login

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
)

//AuthorizeHandler is
type AuthorizeHandler struct {
	m  logic.IMemberLogic
	op logic.IOperateLogic
}

//NewAuthorizeHandler is
func NewAuthorizeHandler() (u *AuthorizeHandler) {
	return &AuthorizeHandler{

		m:  logic.NewMemberLogic(),
		op: logic.NewOperateLogic(),
	}
}

func (u *AuthorizeHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------根据code验证用户登录状态---------")

	if err := ctx.Request().Check("code", "ident"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("1. 根据code验证登录状态")
	info, err := u.m.GetUserInfoByCode(ctx.Request().GetString("code"),
		ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 保存登录结果")
	u.op.LoginOperate(info)

	ctx.Log().Info("3. 返回用户信息")
	return info
}
