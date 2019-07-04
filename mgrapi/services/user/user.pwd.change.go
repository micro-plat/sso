package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/access/member"
	"github.com/micro-plat/sso/mgrapi/modules/logic"
)

type UserPasswordHandler struct {
	container component.IContainer
	userLib   logic.IUserLogic
	member    logic.IMemberLogic
}

func NewUserPasswordHandler(container component.IContainer) (u *UserPasswordHandler) {
	return &UserPasswordHandler{
		container: container,
		userLib:   logic.NewUserLogic(container),
		member:    logic.NewMemberLogic(container),
	}
}

// GetHandle  修改用户密码
func (u *UserPasswordHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------修改用户密码--------")

	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	uid := l.UserID
	if err := ctx.Request.Check("expassword", "newpassword"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.校验旧密码")
	err := u.userLib.ChangePwd(int(uid), ctx.Request.GetString("expassword"), ctx.Request.GetString("newpassword"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回结果")
	return "success"
}

// PostHandle 重置密码
func (u *UserPasswordHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------重置密码-------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	if err := ctx.Request.Bind("user_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.权限验证")
	if err := u.member.QueryAuth(int64(l.SystemID), ctx.Request.GetInt64("user_id")); err != nil {
		return err
	}
	ctx.Log.Info("3.执行操作")
	if err := u.userLib.ResetPwd(ctx.Request.GetInt64("user_id")); err != nil {
		return err
	}
	ctx.Log.Info("4.返回结果")
	return "success"
}
