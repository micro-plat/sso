package function

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrapi/modules/const/sql"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

type IDbSystemFunc interface {
	Get(sysid int) (data []map[string]interface{}, err error)
	ChangeStatus(id int, status int) (err error)
	Delete(id int) (err error)
	Edit(input *model.SystemFuncEditInput) (err error)
	Add(input *model.SystemFuncAddInput) (err error)
}

type DbSystemFunc struct {
	c component.IContainer
}

func NewDbSystemFunc(c component.IContainer) *DbSystemFunc {
	return &DbSystemFunc{
		c: c,
	}
}

//Query 获取功能信息列表
func (u *DbSystemFunc) Get(sysid int) (results []map[string]interface{}, err error) {
	db := u.c.GetRegularDB()
	data, q, a, err := db.Query(sql.QuerySysFuncList, map[string]interface{}{
		"sysid": sysid,
	})
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
							children3 := make([]map[string]interface{}, 0, 8)
							for _, row4 := range data {
								if row4.GetInt("parent") == row3.GetInt("id") && row4.GetInt("level_id") == 4 {
									children3 = append(children3, row4)
								}
							}
							children2 = append(children2, row3)
							row3["children"] = children3
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
	return result, nil
}

func (u *DbSystemFunc) ChangeStatus(id int, status int) (err error) {
	db := u.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.EnableSysFunc, map[string]interface{}{
		"id":     id,
		"enable": status,
	})
	if err != nil {
		return fmt.Errorf("禁用/启用系统功能发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	return nil
}

func (u *DbSystemFunc) Delete(id int) (err error) {
	db := u.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.DeleteSysFunc, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return fmt.Errorf("删除系统功能发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	return nil
}

func (u *DbSystemFunc) Edit(input *model.SystemFuncEditInput) (err error) {
	db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"id":      input.Id,
		"name":    input.Name,
		"icon":    input.Icon,
		"path":    input.Path,
		"is_open": input.IsOpen,
	}
	_, q, a, err := db.Execute(sql.EditSysFunc, params)
	if err != nil {
		return fmt.Errorf("编辑系统功能发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	return nil
}

func (u *DbSystemFunc) Add(input *model.SystemFuncAddInput) (err error) {
	db := u.c.GetRegularDB()

	params := map[string]interface{}{
		"sys_id":   input.Sysid,
		"name":     input.Name,
		"icon":     input.Icon,
		"path":     input.Path,
		"parent":   input.Parentid,
		"level_id": input.ParentLevel + 1,
		"is_open":  input.IsOpen,
	}
	var (
		sortrank interface{}
		q        string
		arg      []interface{}
	)

	//1: 查询目录结构中的最大值
	sortrank, q, arg, err = db.Scalar(sql.GetSysFuncSortRank, params)
	if err != nil {
		return fmt.Errorf("添加系统功能发生错误(err:%v),sql:%s,参数：%v", err, q, arg)
	}

	params["sortrank"] = sortrank
	_, q, arg, err = db.Execute(sql.AddSysFunc, params)
	if err != nil {
		return fmt.Errorf("添加系统功能发生错误(err:%v),sql:%s,参数：%v", err, q, arg)
	}
	return nil
}
