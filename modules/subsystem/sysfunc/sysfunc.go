package sysfunc

import (
	"github.com/micro-plat/hydra/component"

)

type ISystemFunc interface {
	Query(sysid int) (result []map[string]interface{}, err error)
	Enable(input map[string]interface{}) (err error)
	Delete(id int) (err error)
	Edit(input map[string]interface{}) (err error)
	Add(input map[string]interface{}) (err error)
}

type SystemFunc struct {
	c  component.IContainer
	db IDbSystemFunc
}

func NewSystemFunc(c component.IContainer) *SystemFunc {
	return &SystemFunc{
		c:  c,
		db: NewDbSystemFunc(c),
	}
}

//Query 获取用系统管理列表
func (u *SystemFunc) Query(sysid int) (results []map[string]interface{}, err error) {
	data, err := u.db.Query(sysid)
	if err != nil {
		return nil,  err
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

func(u *SystemFunc) Enable(input map[string]interface{}) (err error){
	err = u.db.Enable(input)
	if err != nil {
		return err
	}
	return nil
}

func (u *SystemFunc) Delete(id int) (err error){
	err = u.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *SystemFunc) Edit(input map[string]interface{}) (err error){
	err = u.db.Edit(input)
	if err != nil {
		return err
	}
	return nil
}

func (u *SystemFunc) Add(input map[string]interface{}) (err error){
	err = u.db.Add(input)
	if err != nil {
		return err
	}
	return nil
}