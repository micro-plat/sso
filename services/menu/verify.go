package menu

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/menu"
)

//VerifyHandler 菜单权限验证对象
type VerifyHandler struct {
	c component.IContainer
	m menu.IMenu
}

//NewVerifyHandler 创建菜单验证对象
func NewVerifyHandler(container component.IContainer) (u *VerifyHandler) {
	return &VerifyHandler{
		c: container,
		m: menu.NewMenu(container),
	}
}

//Handle 查询用户在指定系统的页面是否有权限
func (u *VerifyHandler) Handle(ctx *context.Context) (r interface{}) {
	path := ctx.Request.GetString("path")
	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	err := u.m.Verify(uid, sysid, path,ctx.Request.GetMethod())
	if err != nil {
		return err
	}
	return ""
}
