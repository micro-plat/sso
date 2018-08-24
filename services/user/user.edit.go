package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/operate"
	"github.com/micro-plat/sso/modules/user"
)

type UserEditHandler struct {
	container component.IContainer
	userLib   user.IUser
	op        operate.IOperate
}

func NewUserEditHandler(container component.IContainer) (u *UserEditHandler) {
	return &UserEditHandler{
		container: container,
		userLib:   user.NewUser(container),
		op:        operate.NewOperate(container),
	}
}

func (u *UserEditHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑个人资料--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("username", "tel", "email"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	if err := u.userLib.Edit(ctx.Request.GetString("username"), ctx.Request.GetString("tel"), ctx.Request.GetString("email")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.UserOperate(
		member.Query(ctx, u.container),
		"编辑个人资料",
		"username",
		ctx.Request.GetInt("username"),
		"tel",
		ctx.Request.GetInt("tel"),
		"email",
		ctx.Request.GetString("email"),
	); err != nil {
		return err
	}
	ctx.Log.Info("4.返回结果")
	return "success"
}
