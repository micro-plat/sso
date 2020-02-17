package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/permission"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

//IDataPermissionLogic 数据权限
type IDataPermissionLogic interface {
	Query(sysID, name, tableName string, pi int, ps int) (data db.QueryRows, count int, err error)
	Delete(id int) (err error)
	Add(input *model.DataPermissionReq) (err error)
	Edit(input *model.DataPermissionReq) (err error)
	ChangePermissionConfigStatus(id string, status int) error
}

type DataPermissionLogic struct {
	c  component.IContainer
	db permission.IDbDataPermission
}

func NewDataPermissionLogic(c component.IContainer) *DataPermissionLogic {
	return &DataPermissionLogic{
		c:  c,
		db: permission.NewDbDataPermission(c),
	}
}

//Query 获取数据权限管理列表
func (u *DataPermissionLogic) Query(sysID, name, tableName string, pi int, ps int) (data db.QueryRows, count int, err error) {
	data, count, err = u.db.Query(sysID, name, tableName, pi, ps)
	if err != nil {
		return nil, 0, err
	}
	return data, count, err
}

//Delete 删除
func (u *DataPermissionLogic) Delete(id int) (err error) {
	if err = u.db.Delete(id); err != nil {
		return
	}
	return nil
}

//Add 添加
func (u *DataPermissionLogic) Add(input *model.DataPermissionReq) (err error) {
	return u.db.Add(input)
}

//Edit 编辑
func (u *DataPermissionLogic) Edit(input *model.DataPermissionReq) (err error) {
	return u.db.Edit(input)
}

//ChangePermissionConfigStatus 改变状态
func (u *DataPermissionLogic) ChangePermissionConfigStatus(id string, status int) error {
	return u.db.ChangePermissionConfigStatus(id, status)
}
