package member

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/logic"
)

//RoleTagsHandler 角色相关功能
type RoleTagsHandler struct {
	sys logic.IMemberLogic
}

//NewRoleTagsHandler new
func NewRoleTagsHandler() (u *RoleTagsHandler) {
	return &RoleTagsHandler{
		sys: logic.NewMemberLogic(),
	}
}

//TagsHandle 根据角色获取所有权限数据
func (u *RoleTagsHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取当前角色的【功能权限】数据------")

	if err := ctx.Request().Check("role_id", "ident"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("-------获取数据------", ctx.Request().GetInt("role_id"), ctx.Request().GetString("ident"))
	result, err := u.sys.GetRoleMenus(ctx.Request().GetInt("role_id"), ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}
	ctx.Log().Info("------返回结果------")
	return result
}
