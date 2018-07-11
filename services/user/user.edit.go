package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/user"
)

type UserEditHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserEditHandler(container component.IContainer) (u *UserEditHandler) {
	return &UserEditHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

type QueryUserEditInput struct {
	UserName string `form:"user_name" json:"user_name"`
	UserID   int64  `form:"user_id" json:"user_id"`
	RoleID   int64  `form:"role_id" json:"role_id"`
	Mobile   int64  `form:"mobile" json:"mobile"`
	Status   int64  `form:"status" json:"status"`
	IsAdd    int64  `form:"is_add" json:"is_add"`
}

func (u *UserEditHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑/添加用户信息--------")
	ctx.Log.Info("1.参数校验")
	var inputData QueryUserEditInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.执行操作")
	err = u.userLib.UserEdit(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"Status": 200,
	}
}
