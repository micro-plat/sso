package function

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/modules/const/sql"
)

type IDbSystemFunc interface {
	Query(sysid int) (data []map[string]interface{}, err error)
	Enable(id int,status int) (err error)
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
func (u *DbSystemFunc) Query(sysid int) (results []map[string]interface{}, err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"sysid": sysid ,
	}
	data, q, a, err := Db.Query(sql.QuerySysFuncList, params)
	if err != nil {
		return nil, fmt.Errorf("获取系统管理列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	result := make([]map[string]interface{}, 0, 4)
	for _, row1 := range data {
		if row1.GetInt("parent") == 0 && row1.GetInt("level_id") == 1 {
			children1 := make([]map[string]interface{}, 0, 4)
			for _, row2 := range data {
				if row2.GetInt("parent") == row1.GetInt("id") && row2.GetInt("level_id") == 2 {
					children2 := make([]map[string]interface{}, 0, 8)
					for _, row3 := range data {
						if row3.GetInt("parent") == row2.GetInt("id") && row3.GetInt("level_id") == 3 {
							children2 = append(children2, row3)
						}
					}
					children1 = append(children1, row2)
					row2["children"] = children2
				}
			}
			row1["children"] = children1
			result = append(result, row1)
		}
	}
	return result,  nil
}

func (u *DbSystemFunc) Enable(id int, status int) (err error) {
	Db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id": id,
		"enable": status,
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

