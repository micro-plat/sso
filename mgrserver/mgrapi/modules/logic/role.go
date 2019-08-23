package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/role"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

type IRoleLogic interface {
	Get(sysID int, roleID int, path string) (data db.QueryRows, err error)
	Query(input *model.QueryRoleInput) (data db.QueryRows, count int, err error)
	ChangeStatus(roleID string, status int) (err error)
	Delete(roleID int) (err error)
	Save(input *model.RoleEditInput) (err error)
	Auth(input *model.RoleAuthInput) (err error)
	QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error)
}

type RoleLogic struct {
	c     component.IContainer
	db    role.IDbRole
	cache role.ICacheRole
}

func NewRoleLogic(c component.IContainer) *RoleLogic {
	return &RoleLogic{
		c:     c,
		db:    role.NewDbRole(c),
		cache: role.NewCacheRole(c),
	}
}

//Get 获取页面权限
func (r *RoleLogic) Get(sysID int, roleID int, path string) (data db.QueryRows, err error) {
	//从缓存中获取页面权限信息，不存在时从数据库中获取
	data, err = r.cache.Get(sysID, roleID, path)
	if data == nil || err != nil {
		//丛数据库获取数据
		if data, err = r.db.Get(sysID, roleID, path); err != nil {
			return nil, err
		}
		if err = r.cache.SetPageAuth(sysID, roleID, path, data); err != nil {
			return nil, err
		}
	}
	return data, nil
}

//Query 获取角色信息列表
func (r *RoleLogic) Query(input *model.QueryRoleInput) (data db.QueryRows, count int, err error) {
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
func (r *RoleLogic) ChangeStatus(roleID string, status int) (err error) {
	if err := r.cache.Delete(); err != nil {
		return err
	}
	return r.db.ChangeStatus(roleID, status)
}

//Delete 删除角色
func (r *RoleLogic) Delete(roleID int) (err error) {
	if err := r.cache.Delete(); err != nil {
		return err
	}
	return r.db.Delete(roleID)
}

//Save 编辑角色信息
func (r *RoleLogic) Save(input *model.RoleEditInput) (err error) {
	data, err := r.db.QueryRoleInfoByName(input.RoleName)
	if err != nil {
		return err
	}
	if input.IsAdd == 1 {
		if data != nil {
			return context.NewError(model.ERR_ROLE_NAMEEXISTS, "角色名称已被使用")
		}
		err = r.db.Add(input)
	} else {
		if data != nil && data.GetInt64("role_id") != input.RoleID {
			return context.NewError(model.ERR_ROLE_NAMEEXISTS, "角色名称已被使用")
		}
		err = r.db.Edit(input)
	}
	if err != nil {
		return
	}
	if err := r.cache.Delete(); err != nil {
		return err
	}
	return nil
}

//Auth 用户授权
func (r *RoleLogic) Auth(input *model.RoleAuthInput) (err error) {
	if err := r.cache.DeleteAuthMenu(); err != nil {
		return err
	}
	if err := r.cache.Delete(); err != nil {
		return err
	}
	return r.db.Auth(input)
}

//QueryAuthMenu 查询用户菜单
func (r *RoleLogic) QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error) {
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
