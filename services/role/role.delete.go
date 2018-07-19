package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/role"
)

type RoleDelHandler struct {
	container component.IContainer
	roleLib   role.IRole
}

func NewRoleDelHandler(container component.IContainer) (u *RoleDelHandler) {
	return &RoleDelHandler{
		container: container,
		roleLib:   role.NewRole(container),
	}
}

func (u *RoleDelHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------删除角色--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("role_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.roleLib.Delete(ctx.Request.GetInt("role_id")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}
