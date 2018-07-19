package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/role"
)

type RoleEditHandler struct {
	container component.IContainer
	roleLib   role.IRole
}

func NewRoleEditHandler(container component.IContainer) (u *RoleEditHandler) {
	return &RoleEditHandler{
		container: container,
		roleLib:   role.NewRole(container),
	}
}

type RoleEditInput struct {
	RoleName string `form:"role_name" json:"role_name"`
	RoleID   int64  `form:"role_id" json:"role_id"`
	Status   int64  `form:"status" json:"status"`
	IsAdd    int64  `form:"is_add" json:"is_add"`
}

func (u *RoleEditHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑/添加角色信息--------")
	ctx.Log.Info("1.参数校验")
	var inputData RoleEditInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.执行操作")
	err = u.roleLib.RoleEdit(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"Status": 200,
	}
}
