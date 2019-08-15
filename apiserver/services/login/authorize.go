package login

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
)

//AuthorizeHandler is
type AuthorizeHandler struct {
	container component.IContainer
	sys       logic.ISystemLogic
	m         logic.IMemberLogic
	op        logic.IOperateLogic
}

//NewAuthorizeHandler is
func NewAuthorizeHandler(container component.IContainer) (u *AuthorizeHandler) {
	return &AuthorizeHandler{
		container: container,
		sys:       logic.NewSystemLogic(container),
		m:         logic.NewMemberLogic(container),
		op:        logic.NewOperateLogic(container),
	}
}

func (u *AuthorizeHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------根据code验证用户登录状态---------")

	if err := ctx.Request.Check("code", "ident"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("1. 根据code验证登录状态")
	info, err := u.m.GetUserInfoByCode(ctx.Request.GetString("code"),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log.Info("2. 保存登录结果")
	u.op.LoginOperate(info)

	ctx.Log.Info("3. 返回用户信息")
	return info
}
