package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/user"
)

type UserPasswordHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserPasswordHandler(container component.IContainer) (u *UserPasswordHandler) {
	return &UserPasswordHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

type UserPasswordInput struct {
	ExPassword  string `form:"expassword" json:"expassword"`
	NewPassword string `form:"newpassword" json:"newpassword"`
}

func (u *UserPasswordHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑用户密码--------")
	ctx.Log.Info("1.参数校验")
	uid := member.Get(ctx).UserID
	if err := ctx.Request.Check("expassword", "newpassword"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.校验旧密码")
	code, err := u.userLib.CheckPswd(ctx.Request.GetString("expassword"), ctx.Request.GetString("newpassword"), uid)
	if err != nil {
		return context.NewError(code, err)
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}
