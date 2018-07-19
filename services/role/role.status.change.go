package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/role"
)

type RoleChangeHandler struct {
	container component.IContainer
	roleLib   role.IRole
}

func NewRoleChangeHandler(container component.IContainer) (u *RoleChangeHandler) {
	return &RoleChangeHandler{
		container: container,
		roleLib:   role.NewRole(container),
	}
}

func (u *RoleChangeHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------修改角色状态--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("role_id", "status"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.roleLib.ChangeStatus(ctx.Request.GetString("role_id"), ctx.Request.GetInt("status")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
