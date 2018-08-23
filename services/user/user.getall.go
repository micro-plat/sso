package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/user"
)

type UserGetAllHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserGetAllHandler(container component.IContainer) (u *UserGetAllHandler) {
	return &UserGetAllHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

func (u *UserGetAllHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询当前系统下用户列表--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	if err := ctx.Request.Check("pi", "ps"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	rows, count, err := u.userLib.GetAll(l.SystemID, ctx.Request.GetInt("pi"), ctx.Request.GetInt("ps"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}
