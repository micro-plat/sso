package menu

import (
	"github.com/micro-plat/sso/modules/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/menu"
)

//GetHandler 菜单查询对象
type GetHandler struct {
	c component.IContainer
	m menu.IGet
}

//NewGetHandler 创建菜单查询对象
func NewGetHandler(container component.IContainer) (u *GetHandler) {
	return &GetHandler{
		c: container,
		m: menu.NewGet(container),
	}
}

//Handle 查询指定用户在指定系统的菜单列表
func (u *GetHandler) Handle(ctx *context.Context) (r interface{}) {
	uid := member.Get(ctx).UserID
	//	sysid := member.Get(ctx).SystemID
	data, err := u.m.Query(uid, 0)
	if err != nil {
		return err
	}
	return data
}
