package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IRole interface {
	Query(input *QueryRoleInput) (data db.QueryRows, count int, err error)
	ChangeStatus(roleID string, status int) (err error)
	Delete(roleID int) (err error)
	Save(input *RoleEditInput) (err error)
	Auth(input *RoleAuthInput) (err error)
	QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error)
}

type Role struct {
	c     component.IContainer
	db    IDbRole
	cache ICacheRole
}

func NewRole(c component.IContainer) *Role {
	return &Role{
		c:     c,
		db:    NewDbRole(c),
		cache: NewCacheRole(c),
	}
}

//Query 获取角色信息列表
func (r *Role) Query(input *QueryRoleInput) (data db.QueryRows, count int, err error) {
	//从缓存中获取角色信息，不存在时从数据库中获取
	data, count, err = r.cache.Query(input)
	if data == nil || count == 0 || err != nil {
		if data, count, err = r.db.Query(input); err != nil {
			return nil, 0, err
		}
		if err = r.cache.Save(input, data, count); err != nil {
			return nil, 0, err
		}
	}
	return data, count, nil
}

//ChangeStatus 修改角色状态
func (r *Role) ChangeStatus(roleID string, status int) (err error) {
	if err := r.cache.Delete(); err != nil {
		return err
	}
	return r.db.ChangeStatus(roleID, status)
}

//Delete 删除角色
func (r *Role) Delete(roleID int) (err error) {
	if err := r.cache.Delete(); err != nil {
		return err
	}
	return r.db.Delete(roleID)
}

//Save 编辑角色信息
func (r *Role) Save(input *RoleEditInput) (err error) {
	if err := r.cache.Delete(); err != nil {
		return err
	}
	if input.IsAdd == 1 {
		return r.db.Add(input)
	}
	return r.db.Edit(input)
}

//Auth 用户授权
func (r *Role) Auth(input *RoleAuthInput) (err error) {
	if err := r.cache.DeleteAuthMenu(); err != nil {
		return err
	}
	if err := r.cache.Delete(); err != nil {
		return err
	}
	return r.db.Auth(input)
}

//QueryAuthMenu 查询用户菜单
func (r *Role) QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error) {
	data, err := r.cache.QueryAuthMenu(sysID, roleID)
	if data == nil || err != nil {
		if data, err = r.db.QueryAuthMenu(sysID, roleID); err != nil {
			return nil, err
		}
		if err = r.cache.SaveAuthMenu(sysID, roleID, data); err != nil {
			return nil, err
		}
	}
	return data, nil
}
