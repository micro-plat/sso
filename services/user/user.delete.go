package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/user"
)

type UserDelHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserDelHandler(container component.IContainer) (u *UserDelHandler) {
	return &UserDelHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

type QueryUserDelInput struct {
	UserID int64 `form:"user_id" json:"user_id"`
}

func (u *UserDelHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------删除用户--------")
	ctx.Log.Info("1.参数校验")
	var inputData QueryUserDelInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.执行操作")
	err = u.userLib.Delete(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"Status": 200,
	}
}
