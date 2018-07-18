package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IRole interface {
	Query(input QueryRoleInput) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(roleID string, status int) (err error)
	Delete(roleID int) (err error)
	RoleEdit(input RoleEditInput) (err error)
	Auth(input RoleAuthInput) (err error)
	AuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error)
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
func (r *Role) Query(input QueryRoleInput) (data db.QueryRows, count interface{}, err error) {
	return r.db.Query(input)
}

//ChangeStatus 修改角色状态
func (r *Role) ChangeStatus(roleID string, status int) (err error) {
	return r.db.ChangeStatus(roleID, status)
}

//Delete 删除角色
func (r *Role) Delete(roleID int) (err error) {
	return r.db.Delete(roleID)
}

//RoleEdit 编辑用户信息
func (r *Role) RoleEdit(input RoleEditInput) (err error) {
	if input.IsAdd == 1 {
		return r.db.Add(input)
	} else {
		return r.db.Edit(input)
	}
}

//Auth 用户授权
func (r *Role) Auth(input RoleAuthInput) (err error) {
	return r.db.Auth(input)
}

//AuthMenu 查询用户菜单
func (r *Role) AuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error) {
	data, err := r.db.AuthMenu(sysID, roleID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
