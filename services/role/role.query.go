package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/role"
)

type RoleHandler struct {
	container component.IContainer
	roleLib   role.IRole
}

func NewRoleHandler(container component.IContainer) (u *RoleHandler) {
	return &RoleHandler{
		container: container,
		roleLib:   role.NewRole(container),
	}
}

func (u *RoleHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询角色信息数据--------")
	ctx.Log.Info("1.参数校验")
	var inputData role.QueryRoleInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	rows, count, err := u.roleLib.Query(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"count": count.(string),
		"list":  rows,
	}
}
