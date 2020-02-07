package logic

import (
	"github.com/micro-plat/hydra/component"

	"github.com/micro-plat/sso/apiserver/apiserver/modules/access/datapermission"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

//IDataPermissionLogic 数据权限
type IDataPermissionLogic interface {
	GetUserDataPermission(req model.DataPermissionGetReq) (result string, err error)
	SyncDataPermission(req model.DataPermissionSyncReq) error
}

//DataPermissionLogic 数据权限
type DataPermissionLogic struct {
	db datapermission.IDBDataPermission
}

//NewDataPermissionLogic 创建登录对象
func NewDataPermissionLogic(c component.IContainer) *DataPermissionLogic {
	return &DataPermissionLogic{
		db: datapermission.NewDBDataPermission(c),
	}
}

//GetUserDataPermission 获取用户有权限的　[数据权限]　数据
func (m *DataPermissionLogic) GetUserDataPermission(req model.DataPermissionGetReq) (result string, err error) {
	return m.db.GetUserDataPermission(req)
}

//SyncDataPermission 同步子系统的　[数据权限]　数据
func (m *DataPermissionLogic) SyncDataPermission(req model.DataPermissionSyncReq) error {
	return m.db.SyncDataPermission(req)
}
