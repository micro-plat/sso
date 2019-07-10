package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/modules/access/menu"
)

// IMenuLogic interface
type IMenuLogic interface {
	Query(uid int64, sysid int) ([]map[string]interface{}, error)
}

// MenuLogic 菜单
type MenuLogic struct {
	c  component.IContainer
	db menu.IMenu
}

// NewMenuLogic new
func NewMenuLogic(c component.IContainer) *MenuLogic {
	return &MenuLogic{
		c:  c,
		db: menu.NewMenu(c),
	}
}

//Query 获取用户指定系统的菜单信息
func (m *MenuLogic) Query(uid int64, sysid int) ([]map[string]interface{}, error) {
	return m.db.Query(uid, sysid)
}
