package role

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/const/sql"
	"github.com/micro-plat/sso/modules/const/util"
)

type IDbRole interface {
	Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(roleID int, status int) (err error)
	Delete(input map[string]interface{}) (err error)
	Edit(input map[string]interface{}) (err error)
	Add(input map[string]interface{}) (err error)
	Auth(input map[string]interface{}) (err error)
	AuthMenu(input map[string]interface{}) (data db.QueryRows, err error)
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
	db := r.c.GetRegularDB()
	count, q, a, err := db.Scalar(sql.QueryRoleInfoListCount, input)
	if err != nil {
		return nil, nil, fmt.Errorf("获取角色信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	input["role_sql"] = " and t.name like '%" + input["role_name"].(string) + "%' "
	data, q, a, err = db.Query(sql.QueryRoleInfoList, input)
	if err != nil {
		return nil, nil, fmt.Errorf("获取角色信息列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return data, count, nil
}

//ChangeStatus 修改角色状态
func (r *DbRole) ChangeStatus(roleID int, status int) (err error) {
	db := r.c.GetRegularDB()
	input := map[string]interface{}{
		"role_id": roleID,
	}
	switch status {
	case util.RoleDisabled:
		input["status"] = util.RoleDisabled
	case util.RoleNormal, util.UserUnLock:
		input["status"] = util.RoleNormal
	}
	_, q, a, err := db.Execute(sql.UpdateRoleStatus, input)
	if err != nil {
		return fmt.Errorf("修改角色状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}

//Delete 删除角色
func (r *DbRole) Delete(input map[string]interface{}) (err error) {
	db := r.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.DeleteRole, input)
	if err != nil {
		return fmt.Errorf("删除角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}

//Edit 编辑角色信息
func (r *DbRole) Edit(input map[string]interface{}) (err error) {
	db := r.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.EditRoleInfo, input)
	if err != nil {
		return fmt.Errorf("编辑角色信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
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

	//删除原权限
	_, q, a, err := dbTrans.Execute(sql.DelRoleAuth, map[string]interface{}{
		"role_id": input["role_id"],
		"sys_id":  input["sys_id"],
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除角色原权限发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	if input["selectauth"].(string) == "" {
		return nil
	}
	s := strings.Split(input["selectauth"].(string), ",")
	for i := 0; i < len(s); i++ {
		_, q, a, err := dbTrans.Execute(sql.AddRoleAuth, map[string]interface{}{
			"role_id":  input["role_id"],
			"sys_id":   input["sys_id"],
			"menu_id":  s[i],
			"sortrank": i + 1,
		})
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加角色权限发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	dbTrans.Commit()
	return nil
}

//AuthMenu 添加角色
func (r *DbRole) AuthMenu(input map[string]interface{}) (data db.QueryRows, err error) {
	db := r.c.GetRegularDB()
	data, q, a, err := db.Query(sql.QuerySysMenucList, input)
	if err != nil {
		return nil, fmt.Errorf("获取菜单列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return data, nil
}
