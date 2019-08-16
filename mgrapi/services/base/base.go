package base

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/access/member"
	"github.com/micro-plat/sso/mgrapi/modules/logic"
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

//GetHandle 查询用户角色列表
func (u *BaseUserHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询用户角色列表--------")
	rows, err := u.baseLib.QueryUserRoleList()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return rows

	// return map[string]interface{}{
	// 	"rolelist": rows,
	// }
}

//PostHandle 查询系统列表
func (u *BaseUserHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询系统列表--------")
	ctx.Log.Info("1.获取数据")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	rows, err := u.baseLib.QuerySysList()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return rows

	// return map[string]interface{}{
	// 	"list": rows,
	// }
}
