package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
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

type QueryRoleChangeInput struct {
	RoleID   int64 `form:"role_id" json:"role_id"`
	ExStatus int64 `form:"ex_status" json:"ex_status"`
}

func (u *RoleChangeHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------修改角色状态--------")
	ctx.Log.Info("1.参数校验")
	var inputData QueryRoleChangeInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.执行操作")
	err = u.roleLib.ChangeStatus(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"Status": 200,
	}
}
