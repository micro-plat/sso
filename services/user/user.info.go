package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
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

type QueryUserInfoInput struct {
	UserID int64 `form:"user_id" json:"user_id"`
}

func (u *UserInfoHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询用户信息--------")
	ctx.Log.Info("1.参数校验")
	var inputData QueryUserInfoInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.userLib.UserInfo(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"userinfo": data,
	}
}
