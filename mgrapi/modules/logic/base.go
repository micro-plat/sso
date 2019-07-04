package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrapi/modules/access/base"
)

type IBaseLogic interface {
	QueryUserRoleList() (data db.QueryRows, err error)
	QuerySysList() (data db.QueryRows, err error)
}

type BaseLogic struct {
	c  component.IContainer
	db base.IBase
}

// NewBaseLogic xx
func NewBaseLogic(c component.IContainer) *BaseLogic {
	return &BaseLogic{
		c:  c,
		db: base.NewBase(c),
	}
}

//QueryUserRoleList 获取用户角色列表
func (b *BaseLogic) QueryUserRoleList() (data db.QueryRows, err error) {
	return b.db.QueryUserRoleList()
}

//QuerySysList 获取系统列表
func (b *BaseLogic) QuerySysList() (data db.QueryRows, err error) {
	return b.db.QuerySysList()
}
