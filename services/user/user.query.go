package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/system"
	"github.com/micro-plat/sso/modules/user"
)

type UserHandler struct {
	container component.IContainer
	userLib   user.IUser
	sys       system.ISystem
}

func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		container: container,
		userLib:   user.NewUser(container),
		sys:       system.NewSystem(container),
	}
}

// 获取用户所拥有的系统列表
func (u *UserHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------查询用户所有系统列表---------")
	userID := member.Get(ctx).UserID
	datas, err := u.sys.GetAll(userID)
	if err != nil {
		return err
	}
	ctx.Log.Info("返回数据")
	return datas
}

func (u *UserHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询用户信息数据--------")
	ctx.Log.Info("1.参数校验")
	var inputData user.QueryUserInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	rows, count, err := u.userLib.Query(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}
