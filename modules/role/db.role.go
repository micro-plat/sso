package role

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/const/sql"
	"github.com/micro-plat/sso/modules/const/util"
)

type IDbRole interface {
	Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(input map[string]interface{}) (err error)
	Delete(input map[string]interface{}) (err error)
	Edit(input map[string]interface{}) (err error)
	Add(input map[string]interface{}) (err error)
	Auth(input map[string]interface{}) (err error)
}

type DbRole struct {
	c component.IContainer
}

func NewDbRole(c component.IContainer) *DbRole {
	return &DbRole{
		c: c,
	}
}

//Query 获取角色信息列表
func (r *DbRole) Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error) {
	input["role_sql"] = " and t.name like '%" + input["role_name"].(string) + "%' "
	db := r.c.GetRegularDB()
	count, q, a, err := db.Scalar(sql.QueryRoleInfoListCount, input)
	if err != nil {
		return nil, nil, fmt.Errorf("获取角色信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	data, q, a, err = db.Query(sql.QueryRoleInfoList, input)
	if err != nil {
		return nil, nil, fmt.Errorf("获取角色信息列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return data, count, nil
}

//ChangeStatus 修改角色状态
func (r *DbRole) ChangeStatus(input map[string]interface{}) (err error) {
	if input["ex_status"].(float64) == util.RoleDisabled {
		input["status"] = util.RoleNormal
	} else if input["ex_status"].(float64) == util.UserNormal {
		input["status"] = util.RoleDisabled
	}

	db := r.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.UpdateRoleStatus, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("修改角色状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//Delete 删除角色
func (r *DbRole) Delete(input map[string]interface{}) (err error) {
	db := r.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.DeleteRole, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//Edit 编辑角色信息
func (r *DbRole) Edit(input map[string]interface{}) (err error) {
	db := r.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.EditRoleInfo, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("编辑角色信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//Add 添加角色
func (r *DbRole) Add(input map[string]interface{}) (err error) {
	db := r.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.AddRoleInfo, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//Auth 添加角色权限
func (r *DbRole) Auth(input map[string]interface{}) (err error) {
	db := r.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	fmt.Println(input)
	s := input["selectauth"].([]string)
	for i := 0; i < len(s); i++ {
		_, q, a, err := dbTrans.Execute(sql.AddRoleAuth, map[string]interface{}{
			"role_id": input["role_id"],
			"sys_id":  input["sys_id"],
			"menu_id": s[i],
		})
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加角色权限发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	dbTrans.Commit()
	return nil
}
