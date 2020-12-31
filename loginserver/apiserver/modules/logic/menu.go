package logic

import (
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	commodle "github.com/micro-plat/sso/common/module/model"

	"github.com/micro-plat/sso/loginserver/apiserver/modules/access/menu"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/access/system"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/const/enum"
)

// IMenuLogic interface
type IMenuLogic interface {
	Query(uid int64, ident string) ([]map[string]interface{}, error)
	GetTags(uid int64, ident string) (types.XMaps, error)
}

// MenuLogic 菜单
type MenuLogic struct {
	db    menu.IMenu
	dbSys system.IDbSystem
}

// NewMenuLogic new
func NewMenuLogic() *MenuLogic {
	return &MenuLogic{
		db:    menu.NewMenu(),
		dbSys: system.NewDbSystem(),
	}
}

//Query 获取用户指定系统的菜单信息
func (m *MenuLogic) Query(uid int64, ident string) ([]map[string]interface{}, error) {
	status, err := m.dbSys.QuerySystemStatus(ident)
	if err != nil {
		return nil, err
	}
	if status == enum.SystemDisable {
		return nil, errs.NewError(commodle.ERR_SYS_LOCKED, "系统被禁用")
	}

	return m.db.Query(uid, ident)
}

//GetTags 获取按钮级tags
func (m *MenuLogic) GetTags(uid int64, ident string) (types.XMaps, error) {
	return m.db.QueryUserMenuTags(uid, ident)
}
