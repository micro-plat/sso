package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/access/base"
)

type IBaseLogic interface {
	QueryUserRoleList(belongID, belongType int) (data db.QueryRows, err error)
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
//现在每个加油站的角色是一样的,后面会改成不同的加油站角色不一样, belongID暂时不要
//belongType 所属类型(0:加油站、1:公司)
func (b *BaseLogic) QueryUserRoleList(belongID, belongType int) (data db.QueryRows, err error) {
	return b.db.QueryUserRoleList(belongID, belongType)
}

//QuerySysList 获取系统列表
func (b *BaseLogic) QuerySysList() (data db.QueryRows, err error) {
	return b.db.QuerySysList()
}
