package base

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/logic"
)

//BaseUserHandler is
type BaseUserHandler struct {
	container component.IContainer
	baseLib   logic.IBaseLogic
}

//NewBaseUserHandler is
func NewBaseUserHandler(container component.IContainer) (u *BaseUserHandler) {
	return &BaseUserHandler{
		container: container,
		baseLib:   logic.NewBaseLogic(container),
	}
}

//GetRolesHandle 查询用户角色列表
func (u *BaseUserHandler) GetRolesHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询角色列表--------")
	rows, err := u.baseLib.QueryUserRoleList()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return rows
}

//GetSystemsHandle 查询系统列表
func (u *BaseUserHandler) GetSystemsHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询系统列表--------")

	ctx.Log.Info("1.获取数据")
	rows, err := u.baseLib.QuerySysList()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return rows
}

//GetPermissTypesHandle 查询某个系统下面所有的数据权限类型
func (u *BaseUserHandler) GetPermissTypesHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询某个系统下面所有的数据权限类型--------")

	ctx.Log.Info("验证参数")
	if err := ctx.Request.Check("sys_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("1.获取数据")
	rows, err := u.baseLib.GetPermissTypes(ctx.Request.GetString("sys_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return rows
}
