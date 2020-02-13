package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/permission"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

//IDataPermissionLogic 数据权限
type IDataPermissionLogic interface {
	GetTypeInfo(sysID string) (s db.QueryRows, err error)
	Query(sysID, tableName string, pi int, ps int) (data db.QueryRows, count int, err error)
	Delete(id int) (err error)
	Add(input *model.DataPermissionReq) (err error)
	Edit(input *model.DataPermissionReq) (err error)
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

//GetTypeInfo 获取类型信息
func (u *DataPermissionLogic) GetTypeInfo(sysID string) (s db.QueryRows, err error) {
	if s, err = u.db.GetTypeInfo(sysID); err != nil {
		return nil, err
	}
	return s, nil
}

//Query 获取数据权限管理列表
func (u *DataPermissionLogic) Query(sysID, tableName string, pi int, ps int) (data db.QueryRows, count int, err error) {
	data, count, err = u.db.Query(sysID, tableName, pi, ps)
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
	if err = u.db.Add(input); err != nil {
		return
	}
	return nil
}

//Edit 编辑
func (u *DataPermissionLogic) Edit(input *model.DataPermissionReq) (err error) {
	if err = u.db.Edit(input); err != nil {
		return
	}
	return nil
}
