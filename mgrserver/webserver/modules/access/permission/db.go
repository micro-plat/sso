package permission

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/const/sqls"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/model"
)

type IDbDataPermission interface {
	Query(sysID, name, tableName string, pi int, ps int) (data db.QueryRows, count int, err error)
	Delete(id int) (err error)
	Add(input *model.DataPermissionReq) (err error)
	Edit(input *model.DataPermissionReq) (err error)
	ChangePermissionConfigStatus(id string, status int) error
}

type DbDataPermission struct {
}

func NewDbDataPermission() *DbDataPermission {
	return &DbDataPermission{}
}

//Query 获取数据权限 数据
func (u *DbDataPermission) Query(sysID, name, tableName string, pi int, ps int) (data db.QueryRows, count int, err error) {
	db := components.Def.DB().GetRegularDB()
	c, q, a, err := db.Scalar(sqls.QueryDataPermissionTotalCount, map[string]interface{}{
		"sys_id":     sysID,
		"name":       " and name like '%" + name + "%'",
		"table_name": " and table_name like '%" + tableName + "%'",
	})

	if err != nil {
		return nil, 0, fmt.Errorf("获取系统管理列表条数发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sqls.QueryDataPermissionList, map[string]interface{}{
		"sys_id":     sysID,
		"name":       " and name like '%" + name + "%'",
		"table_name": " and table_name like '%" + tableName + "%'",
		"start":      (pi - 1) * ps,
		"ps":         ps,
	})

	if err != nil {
		return nil, 0, fmt.Errorf("获取数据权限 数据 发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c), nil
}

//Delete 删除权限数据
func (u *DbDataPermission) Delete(id int) (err error) {
	db := components.Def.DB().GetRegularDB()
	_, q, a, err := db.Execute(sqls.DeletePermissionInfoById, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return fmt.Errorf("删除删除权限数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return nil
}

//Add 新增数据权限数据
func (u *DbDataPermission) Add(input *model.DataPermissionReq) (err error) {
	db := components.Def.DB().GetRegularDB()
	sysInfo, qs, as, errs := db.Query(sqls.QuerySystemInfoById, map[string]interface{}{
		"id": input.SysID,
	})
	if errs != nil {
		return fmt.Errorf("查询系统信息出错: (err:%+v),sql:%s,输入参数:%v,", errs, qs, as)
	}
	sysfist := sysInfo.Get(0)

	params := map[string]interface{}{
		"name":           input.Name,
		"sys_id":         input.SysID,
		"ident":          sysfist.GetString("ident"),
		"table_name":     input.TableName,
		"operate_action": input.OperateAction,
		"rules":          input.Rules,
		"remark":         input.Remark,
	}

	_, q, a, err := db.Execute(sqls.AddDataPermission, params)
	if err != nil {
		return fmt.Errorf("新增数据权限数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

//Edit 修改数据权限数据
func (u *DbDataPermission) Edit(input *model.DataPermissionReq) (err error) {
	db := components.Def.DB().GetRegularDB()
	_, q, a, err := db.Execute(sqls.UpdateDataPermission, map[string]interface{}{
		"id":             input.ID,
		"name":           input.Name,
		"operate_action": input.OperateAction,
		"rules":          input.Rules,
		"remark":         input.Remark,
	})
	if err != nil {
		return fmt.Errorf("修改数据权限数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

//ChangePermissionConfigStatus 改变规则状态
func (u *DbDataPermission) ChangePermissionConfigStatus(id string, status int) error {
	db := components.Def.DB().GetRegularDB()
	_, q, a, err := db.Execute(sqls.ChangePermissionConfigStatus, map[string]interface{}{
		"id":     id,
		"status": status,
	})
	if err != nil {
		return fmt.Errorf("改变规则状态发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}
