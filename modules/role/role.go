package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IRole interface {
	Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(input map[string]interface{}) (err error)
	Delete(input map[string]interface{}) (err error)
	RoleEdit(input map[string]interface{}) (err error)
	Auth(input map[string]interface{}) (err error)
	AuthMenu(input map[string]interface{}) (results []map[string]interface{}, err error)
}

type Role struct {
	c  component.IContainer
	db IDbRole
}

func NewRole(c component.IContainer) *Role {
	return &Role{
		c:  c,
		db: NewDbRole(c),
	}
}

//Query 获取角色信息列表
func (r *Role) Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error) {
	data, count, err = r.db.Query(input)
	if err != nil {
		return nil, nil, err
	}
	return data, count, nil
}

//ChangeStatus 修改角色状态
func (r *Role) ChangeStatus(input map[string]interface{}) (err error) {
	err = r.db.ChangeStatus(input)
	if err != nil {
		return err
	}
	return nil
}

//Delete 删除角色
func (r *Role) Delete(input map[string]interface{}) (err error) {
	err = r.db.Delete(input)
	if err != nil {
		return err
	}
	return nil
}

//RoleEdit 编辑用户信息
func (r *Role) RoleEdit(input map[string]interface{}) (err error) {
	if input["is_add"].(float64) == 1 {
		err = r.db.Add(input)
		if err != nil {
			return err
		}
	} else {
		err = r.db.Edit(input)
		if err != nil {
			return err
		}
	}
	return nil
}

//Auth 编辑用户信息
func (r *Role) Auth(input map[string]interface{}) (err error) {
	err = r.db.Auth(input)
	if err != nil {
		return err
	}
	return nil
}

//AuthMenu 编辑用户信息
func (r *Role) AuthMenu(input map[string]interface{}) (results []map[string]interface{}, err error) {
	data, err := r.db.AuthMenu(input)
	if err != nil {
		return nil, err
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
							if row3.GetInt("checked") == 1 {
								row3["checked"] = true
							} else {
								row3["checked"] = false
							}
							row3["expanded"] = true
							children2 = append(children2, row3)
						}
					}
					children1 = append(children1, row2)
					row2["children"] = children2
					row2["expanded"] = true
					if row2.GetInt("checked") == 1 {
						row2["checked"] = true
					} else {
						row2["checked"] = false
					}
				}
			}
			row1["children"] = children1
			row1["expanded"] = true
			if row1.GetInt("checked") == 1 {
				row1["checked"] = true
			} else {
				row1["checked"] = false
			}
			result = append(result, row1)
		}
	}
	return result, nil
}
