package user

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/const/sql"
)

type IDbSystem interface {
	Get(page int) (data db.QueryRows, count interface{}, err error)
	Query(input map[string]interface{}) (data db.QueryRows, err error)
	Delete(id int) (err error)
	Add(input map[string]interface{}) (err error)
	ChangeStatus(sysID int, status int) (err error)
	Edit(input map[string]interface{}) (err error)
}

type DbSystem struct {
	c component.IContainer
}

func NewDbSystem(c component.IContainer) *DbSystem {
	return &DbSystem{
		c: c,
	}
}

//Get 获取用系统列表
func (u *DbSystem) Get(page int) (data db.QueryRows, count interface{}, err error) {
	cdb := u.c.GetRegularDB()
	count, q, a, err := cdb.Scalar(sql.QuerySubSystemTotalCount, map[string]interface{}{})
	if err != nil {
		return nil, nil, fmt.Errorf("获取系统管理列表条数发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	data, q, a, err = cdb.Query(sql.QuerySubSystemPageList, map[string]interface{}{
		"page":     page,
		"pageSize": 10,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("获取系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, count, nil
}

func (u *DbSystem) Query(input map[string]interface{}) (data db.QueryRows, err error) {

	db := u.c.GetRegularDB()
	data, q, a, err := db.Query(sql.QuerySubSystemList, map[string]interface{}{
		"name":   input["name"],
		"enable": input["status"],
	})
	if err != nil {
		return nil, fmt.Errorf("获取系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}

func (u *DbSystem) Delete(id int) (err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id": id,
	}
	_, q, a, err := Db.Execute(sql.DeleteSubSystemById, params)
	if err != nil {
		return fmt.Errorf("删除系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (u *DbSystem) Add(input map[string]interface{}) (err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"name":     input["name"],
		"addr":     input["addr"],
		"time_out": input["time_out"],
		"logo":     input["logo"],
		"style":    input["style"],
		"theme":    input["theme"],
	}
	_, q, a, err := Db.Execute(sql.AddSubSystem, params)
	if err != nil {
		return fmt.Errorf("添加系统管理数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (u *DbSystem) ChangeStatus(sysId int, status int) (err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id":     sysId,
		"enable": status,
	}
	_, q, a, err := Db.Execute(sql.UpdateEnable, params)
	if err != nil {
		return fmt.Errorf("更新系统管理状态发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (u *DbSystem) Edit(input map[string]interface{}) (err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"enable":        input["enable"],
		"id":            input["id"],
		"index_url":     input["index_url"],
		"login_timeout": input["login_timeout"],
		"logo":          input["logo"],
		"name":          input["name"],
		"layout":        input["layout"],
		"theme":         input["theme"],
	}
	_, q, a, err := Db.Execute(sql.UpdateEdit, params)
	if err != nil {
		return fmt.Errorf("更新系统管理数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}
