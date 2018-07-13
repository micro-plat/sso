package user

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/const/sql"
)

type IDbSystem interface {
	Query() (data db.QueryRows, count interface{}, err error)
	QueryWithField(input map[string]interface{}) (data db.QueryRows,err error)
	DeleteById(id int) (err error)
	Add(input map[string]interface{}) (err error)
	UpdateEnable(input map[string]interface{}) (err error)
	UpdateEdit(input map[string]interface{}) (err error)
}

type DbSystem struct {
	c component.IContainer
}

func NewDbSystem(c component.IContainer) *DbSystem {
	return &DbSystem{
		c: c,
	}
}

//Query 获取用户信息列表
func (u *DbSystem) Query() (data db.QueryRows, count interface{}, err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id":  0,

	}
	count, q, a, err := Db.Scalar(sql.QuerySubSystemListCount, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取系统管理列表条数发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	data, q, a, err = Db.Query(sql.QuerySubSystemList, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, count, nil
}


func (u *DbSystem) QueryWithField(input map[string]interface{}) (data db.QueryRows, err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"name":  input["name"],
		"enable": input["status"],
	}

	data, q, a, err := Db.Query(sql.QuerySubSystemListWithField, params)
	if err != nil {
		return nil, fmt.Errorf("获取系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}

func (u *DbSystem) DeleteById(id int) (err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id":  id,
	}
	_,q,a,err := Db.Execute(sql.DeleteSubSystemById,params)
	if err != nil {
		return  fmt.Errorf("删除系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}


func(u *DbSystem) Add(input map[string]interface{}) (err error){
	Db := u.c.GetRegularDB()

	params := map[string]interface{}{
		"name": input["name"],
		"addr": input["addr"],
		"time_out": input["time_out"],
		"logo": input["logo"],
	}

	_,_,_,err = Db.Execute(sql.AddSubSystem,params)
	if err != nil {
		return err
	}
	return nil
}

func (u *DbSystem) UpdateEnable(input map[string]interface{}) (err error){
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id": input["id"],
		"enable": input["status"],
	}
	_,_,_,err = Db.Execute(sql.UpdateEnable,params)
	if err != nil {
		return err
	}
	return nil
}

func (u *DbSystem) UpdateEdit(input map[string]interface{}) (err error){
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"enable": input["enable"],
		"id": input["id"],
		"index_url": input["index_url"],
		"login_timeout": input["login_timeout"],
		"logo": input["logo"],
		"name": input["name"],
	}
	_,_,_,err = Db.Execute(sql.UpdateEdit,params)
	if err != nil {
		return err
	}
	return nil
}

