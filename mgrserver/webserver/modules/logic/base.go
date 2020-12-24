package logic

import (
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/access/base"
)

type IBaseLogic interface {
	QueryUserRoleList() (data db.QueryRows, err error)
	QuerySysList() (data db.QueryRows, err error)
	GetPermissTypes(sysID string) (data db.QueryRows, err error)
}

type BaseLogic struct {
	db base.IBase
}

// NewBaseLogic xx
func NewBaseLogic() *BaseLogic {
	return &BaseLogic{
		db: base.NewBase(),
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

//GetPermissTypes 查询某个系统下面所有的数据权限类型
func (b *BaseLogic) GetPermissTypes(sysID string) (data db.QueryRows, err error) {
	return b.db.GetPermissTypes(sysID)
}
