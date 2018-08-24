package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/operate"
	"github.com/micro-plat/sso/modules/user"
)

type UserChangeHandler struct {
	container component.IContainer
	userLib   user.IUser
	op        operate.IOperate
}

func NewUserChangeHandler(container component.IContainer) (u *UserChangeHandler) {
	return &UserChangeHandler{
		container: container,
		userLib:   user.NewUser(container),
		op:        operate.NewOperate(container),
	}
}

type QueryUserChangeInput struct {
	UserID   int64 `form:"user_id" json:"user_id"`
	ExStatus int64 `form:"ex_status" json:"ex_status"`
}

func (u *UserChangeHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------修改用户状态--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_id", "status"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.userLib.ChangeStatus(ctx.Request.GetInt("user_id"), ctx.Request.GetInt("status")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.UserOperate(
		member.Query(ctx, u.container),
		"修改用户状态",
		"user_id",
		ctx.Request.GetInt("user_id"),
		"status",
		ctx.Request.GetInt("status"),
	); err != nil {
		return err
	}
	ctx.Log.Info("4.返回结果")
	return "success"
}
