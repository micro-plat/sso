package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
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

type QueryRoleInput struct {
	PageIndex int    `form:"pi" json:"pi"`
	PageSize  int    `form:"ps" json:"ps"`
	RoleName  string `form:"role_name" json:"role_name"`
}

func (u *RoleHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询角色信息数据--------")
	ctx.Log.Info("1.参数校验")
	var inputData QueryRoleInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	rows, count, err := u.roleLib.Query(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"count": count.(string),
		"list":  rows,
	}
}
