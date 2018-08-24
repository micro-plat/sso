package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/operate"
	"github.com/micro-plat/sso/modules/role"
)

type RoleSaveHandler struct {
	container component.IContainer
	roleLib   role.IRole
	op        operate.IOperate
}

func NewRoleSaveHandler(container component.IContainer) (u *RoleSaveHandler) {
	return &RoleSaveHandler{
		container: container,
		roleLib:   role.NewRole(container),
		op:        operate.NewOperate(container),
	}
}

func (u *RoleSaveHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑/添加角色信息--------")
	ctx.Log.Info("1.参数校验")
	var inputData role.RoleEditInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.roleLib.Save(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.记录行为")
	if err := u.op.RoleOperate(
		member.Query(ctx, u.container),
		"编辑/添加角色",
		"role_name", &inputData.RoleName, "role_id", &inputData.RoleID, "status", &inputData.Status, "IsAdd", &inputData.IsAdd,
	); err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}
