package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/user"
)

type UserSaveHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserSaveHandler(container component.IContainer) (u *UserSaveHandler) {
	return &UserSaveHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

func (u *UserSaveHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑/添加用户信息--------")
	ctx.Log.Info("1.参数校验")
	var inputData user.UserEditInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.userLib.Save(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}
