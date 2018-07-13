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
