package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/access/menu"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/access/system"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/enum"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

// IMenuLogic interface
type IMenuLogic interface {
	Query(uid int64, ident string) ([]map[string]interface{}, error)
}

// MenuLogic 菜单
type MenuLogic struct {
	c     component.IContainer
	db    menu.IMenu
	dbSys system.IDbSystem
}

// NewMenuLogic new
func NewMenuLogic(c component.IContainer) *MenuLogic {
	return &MenuLogic{
		c:     c,
		db:    menu.NewMenu(c),
		dbSys: system.NewDbSystem(c),
	}
}

//Query 获取用户指定系统的菜单信息
func (m *MenuLogic) Query(uid int64, ident string) ([]map[string]interface{}, error) {
	status, err := m.dbSys.QuerySystemStatus(ident)
	if err != nil {
		return nil, err
	}
	if status == enum.SystemDisable {
		return nil, context.NewError(model.ERR_SYS_LOCKED, "系统被禁用")
	}

	return m.db.Query(uid, ident)
}
