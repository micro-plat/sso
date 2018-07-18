package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/role"
)

type RoleAuthHandler struct {
	container component.IContainer
	roleLib   role.IRole
}

func NewRoleAuthHandler(container component.IContainer) (u *RoleAuthHandler) {
	return &RoleAuthHandler{
		container: container,
		roleLib:   role.NewRole(container),
	}
}

func (u *RoleAuthHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------角色授权--------")
	ctx.Log.Info("1.参数校验")
	var inputData role.RoleAuthInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.roleLib.Auth(inputData); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}
