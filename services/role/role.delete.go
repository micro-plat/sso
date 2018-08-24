package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/operate"
	"github.com/micro-plat/sso/modules/role"
)

type RoleDelHandler struct {
	container component.IContainer
	roleLib   role.IRole
	op        operate.IOperate
}

func NewRoleDelHandler(container component.IContainer) (u *RoleDelHandler) {
	return &RoleDelHandler{
		container: container,
		roleLib:   role.NewRole(container),
		op:        operate.NewOperate(container),
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
	ctx.Log.Info("3.记录行为")
	if err := u.op.RoleOperate(
		member.Query(ctx, u.container),
		"删除角色",
		"role_id",
		ctx.Request.GetInt("role_id"),
	); err != nil {
		return err
	}
	ctx.Log.Info("3.返回结果。")
	return "success"
}
