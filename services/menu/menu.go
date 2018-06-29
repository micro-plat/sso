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
func (u *MenuHandler) GetHandle(ctx *context.Context) (r interface{}) {
	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	data, err := u.m.Query(uid, sysid)
	if err != nil {
		return err
	}
	return data
}
