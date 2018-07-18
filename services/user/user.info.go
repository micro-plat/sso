package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/user"
)

type UserInfoHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserInfoHandler(container component.IContainer) (u *UserInfoHandler) {
	return &UserInfoHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

func (u *UserInfoHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询用户信息--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.userLib.UserInfo(ctx.Request.GetInt("user_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"userinfo": data,
	}
}
