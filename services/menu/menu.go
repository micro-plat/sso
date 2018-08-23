package menu

import (
	"github.com/micro-plat/sso/modules/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/menu"
)

//MenuHandler 菜单查询对象
type MenuHandler struct {
	c component.IContainer
	m menu.IMenu
}

//NewMenuHandler 创建菜单查询对象
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		c: container,
		m: menu.NewMenu(container),
	}
}

//Handle 查询指定用户在指定系统的菜单列表
func (u *MenuHandler) Handle(ctx *context.Context) (r interface{}) {
	l := member.Query(ctx, u.c)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	data, err := u.m.Query(l.UserID, l.SystemID)
	if err != nil {
		return err
	}
	return data
}
