package role

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

//RoleAuthHandler is
type RoleAuthHandler struct {
	roleLib logic.IRoleLogic
}

//NewRoleAuthHandler is
func NewRoleAuthHandler() (u *RoleAuthHandler) {
	return &RoleAuthHandler{
		roleLib: logic.NewRoleLogic(),
	}
}

//SaveHandle 角色授权
func (u *RoleAuthHandler) SaveHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------角色授权--------")

	ctx.Log().Info("1.参数校验")
	var inputData model.RoleAuthInput
	if err := ctx.Request().Bind(&inputData); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.执行操作")
	err := u.roleLib.Auth(&inputData)
	if _, ok := err.(*errs.Error); ok {
		return err
	}
	if err != nil {
		return errs.NewError(http.StatusNotImplemented, err)
	}

	ctx.Log().Info("3.返回结果。")
	return "success"
}

//QueryHandle 角色授权菜单数据
func (u *RoleAuthHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------角色授权菜单--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("role_id", "sys_id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.执行操作")
	res, err := u.roleLib.QueryAuthMenu(ctx.Request().GetInt64("sys_id"), ctx.Request().GetInt64("role_id"))
	if err != nil {
		return errs.NewError(http.StatusNotImplemented, err)
	}

	ctx.Log().Info("3.返回结果。")
	return res
}
