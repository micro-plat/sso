package permission

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/mgrserver/modules/logic"
	"github.com/micro-plat/sso/mgrserver/modules/model"
)

//AuthPermissionHandler 角色数据权限关联接口
type AuthPermissionHandler struct {
	roleLib logic.IRoleLogic
	op      logic.IOperateLogic
}

//NewAuthPermissionHandler new
func NewAuthPermissionHandler() (u *AuthPermissionHandler) {
	return &AuthPermissionHandler{
		roleLib: logic.NewRoleLogic(),
		op:      logic.NewOperateLogic(),
	}
}

//QueryHandle 数据权限关联查询
func (u *AuthPermissionHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------查询角色与数据权限的关联数据--------")

	ctx.Log().Info("1.参数校验")
	var req model.RolePermissionQueryReq
	if err := ctx.Request().Bind(&req); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.执行操作")
	res, err := u.roleLib.QueryAuthDataPermission(req)
	if err != nil {
		return errs.NewError(http.StatusNotImplemented, err)
	}

	ctx.Log().Info("3.返回结果。")
	return map[string]interface{}{
		"list": res,
	}
}

//SaveHandle 保存角色与数据权限的关联关系
func (u *AuthPermissionHandler) SaveHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------保存角色与数据权限的关联关系--------")

	ctx.Log().Info("1.参数校验")
	var req model.RolePermissionReq
	if err := ctx.Request().Bind(&req); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.保存数据")
	err := u.roleLib.SaveRolePermission(req)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回结果。")
	return "success"
}

//EnableHandle 启用角色与数据权限的关联关系
func (u *AuthPermissionHandler) EnableHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------启用角色与数据权限的关联关系--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.保存数据")
	err := u.roleLib.ChangeRolePermissionStatus(ctx.Request().GetString("id"), model.Enable)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回结果。")
	return "success"
}

//DisableHandle 禁用角色与数据权限的关联关系
func (u *AuthPermissionHandler) DisableHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------禁用角色与数据权限的关联关系--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.保存数据")
	err := u.roleLib.ChangeRolePermissionStatus(ctx.Request().GetString("id"), model.Disable)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回结果。")
	return "success"
}

//DelHandle 删除角色与数据权限的关联关系
func (u *AuthPermissionHandler) DelHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------删除角色与数据权限的关联关系--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.保存数据")
	err := u.roleLib.DelRolePermission(ctx.Request().GetString("id"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回结果。")
	return "success"
}
