package menu

import (
	"github.com/micro-plat/sso/mgrapi/modules/access/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/access/menu"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

//MenuHandler 菜单查询对象
type MenuHandler struct {
	c component.IContainer
	m menu.IMenu
	p menu.IPopular
}

//NewMenuHandler 创建菜单查询对象
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		c: container,
		m: menu.NewMenu(container),
		p: menu.NewPopular(container),
	}
}

//GetHandle 调用远程api得到用户的菜单数据
func (u *MenuHandler) GetHandle(ctx *context.Context) (r interface{}) {
	user := member.Get(ctx)
	menus, err := model.GetSSOClient(u.c).GetUserMenu(int(user.UserID))
	if err != nil {
		return err
	}
	return menus
}

//VerifyHandle 查询用户在指定系统的页面是否有权限
func (u *MenuHandler) VerifyHandle(ctx *context.Context) (r interface{}) {
	path := ctx.Request.GetString("path")
	method := ctx.Request.GetString("method", "get")
	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	err := u.m.Verify(uid, sysid, path, method)
	if err != nil {
		return err
	}
	return ""
}
