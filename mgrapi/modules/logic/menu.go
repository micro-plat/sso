package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrapi/modules/access/menu"
)

//Get 获取全局IMenuLogic
func Get(c component.IContainer) IMenuLogic {
	return c.Get("__imenu__").(IMenuLogic)
}

//Set 保存全局imenu
func Set(c component.IContainer) {
	c.Set("__imenu__", NewMenuLogic(c))
}

// IMenuLogic interface
type IMenuLogic interface {
	Query(uid int64, sysid int) ([]map[string]interface{}, error) // 对外api有在用
	Verify(uid int64, sysid int, menuURL string, method string) error
}

// MenuLogic 菜单
type MenuLogic struct {
	c  component.IContainer
	db menu.IMenu
}

// NewMenuLogic xx
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

//Verify 获取用户指定系统的菜单信息
func (m *MenuLogic) Verify(uid int64, sysid int, menuURL string, method string) error {
	return m.db.Verify(uid, sysid, menuURL, method)
}
