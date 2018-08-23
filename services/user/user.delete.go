package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/user"
)

type UserDelHandler struct {
	container component.IContainer
	userLib   user.IUser
	member    member.IMember
}

func NewUserDelHandler(container component.IContainer) (u *UserDelHandler) {
	return &UserDelHandler{
		container: container,
		userLib:   user.NewUser(container),
		member:    member.NewMember(container),
	}
}

func (u *UserDelHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------删除用户--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	if err := ctx.Request.Check("user_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.权限验证")
	if err := u.member.QueryAuth(int64(l.SystemID), ctx.Request.GetInt64("user_id")); err != nil {
		return err
	}
	ctx.Log.Info("3.执行操作")
	if err := u.userLib.Delete(ctx.Request.GetInt("user_id")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("4.返回结果。")
	return "success"
}
