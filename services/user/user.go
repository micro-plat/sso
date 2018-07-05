package user

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/user"
)

type UserHandler struct {
	container component.IContainer
	userLib   user.IUser
}

func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		container: container,
		userLib:   user.NewUser(container),
	}
}

type QueryUserInput struct {
	PageIndex int    `form:"pi" json:"pi"`
	PageSize  int    `form:"ps" json:"ps"`
	UserName  string `form:"username" json:"username"`
	Role      int64  `form:"roleid" json:"roleid"`
}

func (u *UserHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询用户信息数据--------")
	ctx.Log.Info("1.参数校验")
	var inputData QueryUserInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	inputData.PageSize = 10
	inputData.PageIndex = 1
	fmt.Println("input:", inputData)
	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.获取数据。当前页数、每页数据量:", inputData.PageIndex, inputData.PageSize)
	rows, count, err := u.userLib.QueryUserList(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回数据。")
	return map[string]interface{}{
		"count": count.(string),
		"list":  rows,
	}
}
