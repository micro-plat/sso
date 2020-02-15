package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/sso/apiserver/apiserver/modules/access/datapermission"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

//IDataPermissionLogic 数据权限
type IDataPermissionLogic interface {
	GetUserDataPermissionConfigs(req model.DataPermissionGetReq) (result db.QueryRows, err error)
	GetAllUserInfoByUserRole(userID int, ident string) (string, error)
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

//GetUserDataPermissionConfigs 获取用户有权限的　[数据权限]　规则信息
func (m *DataPermissionLogic) GetUserDataPermissionConfigs(req model.DataPermissionGetReq) (result db.QueryRows, err error) {
	return m.db.GetUserDataPermissionConfigs(req)
}

//GetAllUserInfoByUserRole 获取和当前用户同一个角色的用户ids
func (m *DataPermissionLogic) GetAllUserInfoByUserRole(userID int, ident string) (string, error) {
	return m.db.GetAllUserInfoByUserRole(userID, ident)
}
