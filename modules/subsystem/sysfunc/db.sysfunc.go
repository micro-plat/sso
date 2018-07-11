package sysfunc

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/const/sql"
)

type IDbSystemFunc interface {
	Query(sysid int) (data db.QueryRows, err error)
	Enable(input map[string]interface{}) (err error)
	Delete(id int) (err error)
	Edit(input map[string]interface{}) (err error)
	Add(input map[string]interface{}) (err error)
}

type DbSystemFunc struct {
	c component.IContainer
}

func NewDbSystemFunc(c component.IContainer) *DbSystemFunc {
	return &DbSystemFunc{
		c: c,
	}
}

//Query 获取用户信息列表
func (u *DbSystemFunc) Query(sysid int) (data db.QueryRows, err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"sysid": sysid ,
	}
	data, q, a, err := Db.Query(sql.QuerySysFuncList, params)
	if err != nil {
		return nil, fmt.Errorf("获取系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data,  nil
}

func (u *DbSystemFunc) Enable(input map[string]interface{}) (err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id": input["Id"],
		"enable": input["Status"],
	}
	_,q,a,err := Db.Execute(sql.EnableSysFunc,params)
	if err != nil {
		return fmt.Errorf("禁用/启用系统功能发生错误(err:%v),sql:%s,参数：%v", err, q,a)
	}
	return   nil
}

func (u *DbSystemFunc) Delete(id int) (err error){
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id": id,
	}
	_,q,a,err := Db.Execute(sql.DeleteSysFunc,params)
	if err != nil {
		return fmt.Errorf("删除系统功能发生错误(err:%v),sql:%s,参数：%v", err, q,a)
	}
	return   nil
}

func (u *DbSystemFunc) Edit(input map[string]interface{}) (err error){
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id": input["id"],
		"name": input["name"],
		"icon": input["icon"],
		"path": input["path"],
	}
	_,q,a,err := Db.Execute(sql.EditSysFunc,params)
	if err != nil {
		return fmt.Errorf("编辑系统功能发生错误(err:%v),sql:%s,参数：%v", err, q,a)
	}
	return   nil
}

func (u *DbSystemFunc) Add(input map[string]interface{}) (err error){
	Db := u.c.GetRegularDB()

	params := map[string]interface{}{
		"sys_id": input["sys_id"],
		"name": input["name"],
		"icon": input["icon"],
		"path": input["path"],
		"parent": input["parentid"],
		"level_id": input["level_id"],
	}
	_,q,a,err := Db.Execute(sql.AddSysFunc,params)

	if err != nil {
		return fmt.Errorf("添加系统功能发生错误(err:%v),sql:%s,参数：%v", err, q,a)
	}
	return nil
}

