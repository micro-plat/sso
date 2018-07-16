package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/role"
)

type AuthMenuHandler struct {
	container component.IContainer
	roleLib   role.IRole
}

func NewAuthMenuHandler(container component.IContainer) (u *AuthMenuHandler) {
	return &AuthMenuHandler{
		container: container,
		roleLib:   role.NewRole(container),
	}
}

type AuthMenuInput struct {
	SysID int64 `form:"sys_id" json:"sys_id"`
}

func (u *AuthMenuHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------角色授权菜单--------")
	ctx.Log.Info("1.参数校验")
	var inputData AuthMenuInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.执行操作")
	res, err := u.roleLib.AuthMenu(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return res
}
