package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
)

//RoleHandler 角色相关功能
type RoleHandler struct {
	c   component.IContainer
	sys logic.IMemberLogic
}

//NewRoleHandler new
func NewRoleHandler(container component.IContainer) (u *RoleHandler) {
	return &RoleHandler{
		c:   container,
		sys: logic.NewMemberLogic(container),
	}
}

//GetHandle: 获取和当前用户同一个角色的用户ids
func (u *RoleHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------获取当前用户可用的【数据权限】数据------")

	ctx.Log.Info("-------验证数据的有效性------")
	if err := ctx.Request.Check("user_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("-------获取数据------")
	result, err := u.sys.GetAllUserInfoByUserRole(ctx.Request.GetInt("user_id"), ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log.Info("------返回结果------")
	return result
}
