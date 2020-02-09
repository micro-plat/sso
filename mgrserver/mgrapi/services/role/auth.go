package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

//RoleAuthHandler is
type RoleAuthHandler struct {
	container component.IContainer
	roleLib   logic.IRoleLogic
}

//NewRoleAuthHandler is
func NewRoleAuthHandler(container component.IContainer) (u *RoleAuthHandler) {
	return &RoleAuthHandler{
		container: container,
		roleLib:   logic.NewRoleLogic(container),
	}
}

//SaveHandle 角色授权
func (u *RoleAuthHandler) SaveHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------角色授权--------")

	ctx.Log.Info("1.参数校验")
	var inputData model.RoleAuthInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.roleLib.Auth(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}

//QueryHandle 角色授权菜单数据
func (u *RoleAuthHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------角色授权菜单--------")

	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("role_id", "sys_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	res, err := u.roleLib.QueryAuthMenu(ctx.Request.GetInt64("sys_id"), ctx.Request.GetInt64("role_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return res
}

//PermissionQueryHandle 数据权限关联查询
func (u *RoleAuthHandler) PermissionQueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询角色与数据权限的关联数据--------")

	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("role_id", "sys_id", "data_type"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	res, err := u.roleLib.QueryAuthDataPermission(ctx.Request.GetInt64("sys_id"), ctx.Request.GetInt64("role_id"), ctx.Request.GetString("data_type"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return res
}

//SavePermissionHandle 保存角色与数据权限的关联关系
func (u *RoleAuthHandler) SavePermissionHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------保存角色与数据权限的关联关系--------")

	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("role_id", "sys_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.保存数据")
	err := u.roleLib.SaveRolePermission(ctx.Request.GetInt64("sys_id"), ctx.Request.GetInt64("role_id"), ctx.Request.GetString("select_auth"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}
