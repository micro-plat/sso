package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/user"
)

type UserChangeHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserChangeHandler(container component.IContainer) (u *UserChangeHandler) {
	return &UserChangeHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

type QueryUserChangeInput struct {
	UserID   int64 `form:"user_id" json:"user_id"`
	ExStatus int64 `form:"ex_status" json:"ex_status"`
}

func (u *UserChangeHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------修改用户状态--------")
	ctx.Log.Info("1.参数校验")
	var inputData QueryUserChangeInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.执行操作")
	err = u.userLib.CHangeStatus(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"Status": 200,
	}
}
