package logic

import (
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/role"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	"github.com/micro-plat/sso/sso/errorcode"
)

type IRoleLogic interface {
	//Get(sysID int, roleID int, path string) (data types.XMaps, err error)
	Query(input *model.QueryRoleInput) (data types.XMaps, count int, err error)
	ChangeStatus(roleID string, status int) (err error)
	Delete(roleID int) (err error)
	Save(input *model.RoleEditInput) (err error)
	Auth(input *model.RoleAuthInput) (err error)
	QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error)
	QueryAuthDataPermission(req model.RolePermissionQueryReq) (data types.XMaps, err error)
	SaveRolePermission(req model.RolePermissionReq) error
	ChangeRolePermissionStatus(id string, status int) error
	DelRolePermission(id string) error
}

type RoleLogic struct {
	db    role.IDbRole
	cache role.ICacheRole
}

func NewRoleLogic() *RoleLogic {
	return &RoleLogic{
		db:    role.NewDbRole(),
		cache: role.NewCacheRole(),
	}
}

// //Get 获取页面权限
// func (r *RoleLogic) Get(sysID int, roleID int, path string) (data types.XMaps, err error) {
// 	if data, err = r.db.Get(sysID, roleID, path); err != nil {
// 		return nil, err
// 	}
// 	return data, nil

// }

//Query 获取角色信息列表
func (r *RoleLogic) Query(input *model.QueryRoleInput) (data types.XMaps, count int, err error) {
	if data, count, err = r.db.Query(input); err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

//ChangeStatus 修改角色状态
func (r *RoleLogic) ChangeStatus(roleID string, status int) (err error) {
	return r.db.ChangeStatus(roleID, status)
}

//Delete 删除角色
func (r *RoleLogic) Delete(roleID int) (err error) {
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
			return errs.NewError(errorcode.ERR_ROLE_NAMEEXISTS, "角色名称已被使用")
		}
		err = r.db.Add(input)
	} else {
		if data != nil && data.GetInt64("role_id") != input.RoleID {
			return errs.NewError(errorcode.ERR_ROLE_NAMEEXISTS, "角色名称已被使用")
		}
		err = r.db.Edit(input)
	}
	if err != nil {
		return
	}
	return nil
}

//Auth 用户授权
func (r *RoleLogic) Auth(input *model.RoleAuthInput) (err error) {
	return r.db.Auth(input)
}

//QueryAuthMenu 查询用户菜单
func (r *RoleLogic) QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error) {
	data, err := r.db.QueryAuthMenu(sysID, roleID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//QueryAuthDataPermission 查询角色与数据权限的关联关系
func (r *RoleLogic) QueryAuthDataPermission(req model.RolePermissionQueryReq) (data types.XMaps, err error) {
	return r.db.QueryAuthDataPermission(req)
}

//SaveRolePermission  保存角色与权限数据的关系
func (r *RoleLogic) SaveRolePermission(req model.RolePermissionReq) error {
	return r.db.SaveRolePermission(req)
}

//ChangeRolePermissionStatus 改变 【角色与权限数据关系】的状态
func (r *RoleLogic) ChangeRolePermissionStatus(id string, status int) error {
	return r.db.ChangeRolePermissionStatus(id, status)
}

//DelRolePermission 删除
func (r *RoleLogic) DelRolePermission(id string) error {
	return r.db.DelRolePermission(id)
}
