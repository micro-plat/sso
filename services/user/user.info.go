package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/user"
)

type UserInfoHandler struct {
	container component.IContainer
	userLib   user.IUser
	member    member.IMember
}

func NewUserInfoHandler(container component.IContainer) (u *UserInfoHandler) {
	return &UserInfoHandler{
		container: container,
		userLib:   user.NewUser(container),
		member:    member.NewMember(container),
	}
}

func (u *UserInfoHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询用户信息--------")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	ctx.Log.Info("1.参数校验")
	var uid int64
	err := ctx.Request.Check("user_id")
	if err != nil {
		uid = member.Get(ctx).UserID
	} else {
		uid = ctx.Request.GetInt64("user_id")
	}
	ctx.Log.Info("2.验证权限")
	if err = u.member.QueryAuth(int64(l.SystemID), uid); err != nil {
		return context.NewError(context.ERR_FORBIDDEN, err)
	}

	ctx.Log.Info("3.执行操作")
	data, err := u.userLib.Get(int(uid))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("4.返回结果")
	return map[string]interface{}{
		"userinfo": data,
	}
}
