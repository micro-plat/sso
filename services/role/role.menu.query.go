package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/role"
)

type AuthMenuHandler struct {
	container component.IContainer
	roleLib   role.IRole
}

func NewAuthMenuHandler(container component.IContainer) (u *AuthMenuHandler) {
	return &AuthMenuHandler{
		container: container,
		roleLib:   role.NewRole(container),
	}
}

func (u *AuthMenuHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------角色授权菜单--------")
	ctx.Log.Info("1.参数校验")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	if err := ctx.Request.Check("role_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	var sysID int64
	if l.SystemID == 0 {
		sysID = ctx.Request.GetInt64("sys_id")
	} else {
		sysID = int64(l.SystemID)
	}

	ctx.Log.Info("2.执行操作")
	res, err := u.roleLib.QueryAuthMenu(sysID, ctx.Request.GetInt64("role_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果。")
	return res
}
