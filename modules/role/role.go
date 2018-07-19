package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IRole interface {
	Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(roleID int, status int) (err error)
	Delete(input map[string]interface{}) (err error)
	RoleEdit(input map[string]interface{}) (err error)
	Auth(input map[string]interface{}) (err error)
	AuthMenu(input map[string]interface{}) (results []map[string]interface{}, err error)
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
func (r *Role) Query(input *QueryRoleInput) (data db.QueryRows, count interface{}, err error) {
	//从缓存中获取角色信息，不存在时从数据库中获取
	data, count, err = r.cache.Query(input)
	if data == nil || count == nil || err != nil {
		if data, count, err = r.db.Query(input); err != nil {
			return nil, nil, err
		}
		if err = r.cache.Save(input, data, count); err != nil {
			return nil, nil, err
		}
	}
	return data, count, nil
}

//ChangeStatus 修改角色状态
func (r *Role) ChangeStatus(roleID string, status int) (err error) {
	if err := r.db.ChangeStatus(roleID, status); err != nil {
		return err
	}
	return r.cache.Delete()
}

//Delete 删除角色
func (r *Role) Delete(roleID int) (err error) {
	if err := r.db.Delete(roleID); err != nil {
		return err
	}
	return r.cache.Delete()
}

//Save 编辑角色信息
func (r *Role) Save(input *RoleEditInput) (err error) {
	if input.IsAdd == 1 {
		return r.db.Add(input)
	}
	if err := r.db.Edit(input); err != nil {
		return err
	}
	return r.cache.Delete()
}

//Auth 用户授权
func (r *Role) Auth(input *RoleAuthInput) (err error) {
	if err := r.db.Auth(input); err != nil {
		return err
	}
	if err := r.cache.Delete(); err != nil {
		return err
	}
	return r.cache.DeleteAuthMenu()
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
