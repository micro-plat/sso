package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
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

type QueryRoleAuthInput struct {
	RoleID     int64  `form:"role_id" json:"role_id"`
	SysID      int64  `form:"sys_id" json:"sys_id"`
	SelectAuth string `form:"selectauth" json:"selectauth"`
}

func (u *RoleAuthHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------角色授权--------")
	ctx.Log.Info("1.参数校验")
	var inputData QueryRoleAuthInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.执行操作")
	err = u.roleLib.Auth(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"Status": 200,
	}
}
