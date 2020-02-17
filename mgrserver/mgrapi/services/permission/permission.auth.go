package permission

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

//AuthPermissionHandler 角色数据权限关联接口
type AuthPermissionHandler struct {
	container component.IContainer
	roleLib   logic.IRoleLogic
	op        logic.IOperateLogic
}

//NewAuthPermissionHandler new
func NewAuthPermissionHandler(container component.IContainer) (u *AuthPermissionHandler) {
	return &AuthPermissionHandler{
		container: container,
		roleLib:   logic.NewRoleLogic(container),
		op:        logic.NewOperateLogic(container),
	}
}

//QueryHandle 数据权限关联查询
func (u *AuthPermissionHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询角色与数据权限的关联数据--------")

	ctx.Log.Info("1.参数校验")
	var req model.RolePermissionQueryReq
	if err := ctx.Request.Bind(&req); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	res, err := u.roleLib.QueryAuthDataPermission(req)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return map[string]interface{}{
		"list": res,
	}
}

//SaveHandle 保存角色与数据权限的关联关系
func (u *AuthPermissionHandler) SaveHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------保存角色与数据权限的关联关系--------")

	ctx.Log.Info("1.参数校验")
	var req model.RolePermissionReq
	if err := ctx.Request.Bind(&req); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.保存数据")
	err := u.roleLib.SaveRolePermission(req)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}

//EnableHandle 启用角色与数据权限的关联关系
func (u *AuthPermissionHandler) EnableHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------启用角色与数据权限的关联关系--------")

	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.保存数据")
	err := u.roleLib.ChangeRolePermissionStatus(ctx.Request.GetString("id"), model.Enable)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}

//DisableHandle 禁用角色与数据权限的关联关系
func (u *AuthPermissionHandler) DisableHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------禁用角色与数据权限的关联关系--------")

	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.保存数据")
	err := u.roleLib.ChangeRolePermissionStatus(ctx.Request.GetString("id"), model.Disable)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}

//DelHandle 删除角色与数据权限的关联关系
func (u *AuthPermissionHandler) DelHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------删除角色与数据权限的关联关系--------")

	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.保存数据")
	err := u.roleLib.DelRolePermission(ctx.Request.GetString("id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}
