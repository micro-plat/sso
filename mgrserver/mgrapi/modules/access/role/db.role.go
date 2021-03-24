package role

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/enum"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/sqls"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

type IDbRole interface {
	//Get(sysID int, roleID int, path string) (data types.XMaps, err error)
	Query(input *model.QueryRoleInput) (data types.XMaps, count int, err error)
	ChangeStatus(roleID string, status int) (err error)
	Delete(roleID int) (err error)
	Edit(input *model.RoleEditInput) (err error)
	Add(input *model.RoleEditInput) (err error)
	Auth(input *model.RoleAuthInput) (err error)
	QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error)
	QueryRoleInfoByName(roleName string) (data types.IXMap, err error)
	QueryAuthDataPermission(req model.RolePermissionQueryReq) (data types.XMaps, err error)
	SaveRolePermission(req model.RolePermissionReq) error
	ChangeRolePermissionStatus(id string, status int) error
	DelRolePermission(id string) error
}

type DbRole struct {
}

func NewDbRole() *DbRole {
	return &DbRole{}
}

//QueryRoleInfoByName 通过名称查询角色信息
func (r *DbRole) QueryRoleInfoByName(roleName string) (data types.IXMap, err error) {
	db := components.Def.DB().GetRegularDB()
	result, err := db.Query(sqls.QueryRoleInfoByName, map[string]interface{}{"role_name": roleName})
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() {
		return nil, nil
	}
	return result.Get(0), nil
}

//Query 获取角色信息列表
func (r *DbRole) Query(input *model.QueryRoleInput) (data types.XMaps, count int, err error) {
	db := components.Def.DB().GetRegularDB()
	if err != nil {
		return nil, 0, fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	params := map[string]interface{}{
		"status":   input.Status,
		"role_sql": " and t.name like '%" + input.RoleName + "%' ",
		"start":    (input.PageIndex - 1) * input.PageSize,
		"ps":       input.PageSize,
	}

	c, err := db.Scalar(sqls.QueryRoleInfoListCount, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取角色信息列表条数发生错误(err:%v)", err)
	}

	data, err = db.Query(sqls.QueryRoleInfoList, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取角色信息列表发生错误(err:%v)", err)
	}
	return data, types.GetInt(c), nil
}

//ChangeStatus 修改角色状态
func (r *DbRole) ChangeStatus(roleID string, status int) (err error) {
	db := components.Def.DB().GetRegularDB()
	input := map[string]interface{}{
		"role_id": roleID,
	}
	switch status {
	case enum.Disabled:
		input["status"] = enum.Disabled
	case enum.Normal:
		input["status"] = enum.Normal
	}
	_, err = db.Execute(sqls.UpdateRoleStatus, input)
	if err != nil {
		return fmt.Errorf("修改角色状态发生错误(err:%v)", err)
	}
	return nil
}

//Delete 删除角色
func (r *DbRole) Delete(roleID int) (err error) {
	db := components.Def.DB().GetRegularDB()
	_, err = db.Execute(sqls.DeleteRole, map[string]interface{}{
		"role_id": roleID,
	})
	if err != nil {
		return fmt.Errorf("删除角色发生错误(err:%v)", err)
	}

	_, err = db.Execute(sqls.DeleteRoleMenu, map[string]interface{}{
		"role_id": roleID,
	})
	if err != nil {
		return fmt.Errorf("删除角色菜单发生错误(err:%v)", err)
	}
	return nil
}

//Edit 编辑角色信息
func (r *DbRole) Edit(input *model.RoleEditInput) (err error) {
	db := components.Def.DB().GetRegularDB()
	params, err := types.Struct2Map(input)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}
	_, err = db.Execute(sqls.EditRoleInfo, params)
	if err != nil {
		return fmt.Errorf("编辑角色信息发生错误(err:%v)", err)
	}
	return nil
}

//Add 添加角色
func (r *DbRole) Add(input *model.RoleEditInput) (err error) {
	db := components.Def.DB().GetRegularDB()
	params, err := types.Struct2Map(input)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	_, err = db.Execute(sqls.AddRoleInfo, params)
	if err != nil {
		return fmt.Errorf("添加角色发生错误(err:%v)", err)
	}
	return nil
}

//Auth 添加角色权限
func (r *DbRole) Auth(input *model.RoleAuthInput) (err error) {
	db := components.Def.DB().GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	checkResult, err := db.Query(sqls.CheckSysMenu, map[string]interface{}{
		"sys_id":      input.SysID,
		"select_auth": input.SelectAuth,
	})
	if err != nil {
		return fmt.Errorf("检查菜单列表发生错误(err:%v)", err)
	}
	if checkResult.IsEmpty() {
		dbTrans.Rollback()
		return errs.NewErrorf(enum.Invalid_Code, "角色权限修改未成功，请重试")
	}

	//删除原权限
	_, err = dbTrans.Execute(sqls.DelRoleAuth, map[string]interface{}{
		"role_id": input.RoleID,
		"sys_id":  input.SysID,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除角色原权限发生错误(err:%v)", err)
	}

	if input.SelectAuth == "" {
		dbTrans.Commit()
		return nil
	}
	//添加新权限
	s := strings.Split(input.SelectAuth, ",")
	var num int64 = 0
	for i := 0; i < len(s); i++ {
		r, err := dbTrans.Execute(sqls.AddRoleAuth, map[string]interface{}{
			"role_id":  input.RoleID,
			"sys_id":   input.SysID,
			"menu_id":  s[i],
			"sortrank": i + 1,
		})
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加角色权限发生错误(err:%v)", err)
		}
		num = num + r
	}

	dbTrans.Commit()

	return nil
}

//QueryAuthMenu 查询角色菜单
func (r *DbRole) QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QuerySysMenucList, map[string]interface{}{
		"role_id": roleID,
		"sys_id":  sysID,
	})
	if err != nil {
		return nil, fmt.Errorf("获取菜单列表发生错误(err:%v)", err)
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
									if row4.GetInt("checked") == 1 {
										row4["checked"] = true
									} else {
										row4["checked"] = false
									}
									row4["expanded"] = true
									children3 = append(children3, row4)
								}
							}
							if row3.GetInt("checked") == 1 {
								row3["checked"] = true
							} else {
								row3["checked"] = false
							}
							row3["expanded"] = true
							children2 = append(children2, row3)
							row3["children"] = children3
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

//QueryAuthDataPermission 查询角色与数据权限的关联关系
func (r *DbRole) QueryAuthDataPermission(req model.RolePermissionQueryReq) (data types.XMaps, err error) {
	db := components.Def.DB().GetRegularDB()
	// c, q, a, err := db.Scalar(sqls.QueryRoleDataPermissionCount, map[string]interface{}{
	// 	"sys_id":  req.SysID,
	// 	"role_id": req.RoleID,
	// })

	// if err != nil {
	// 	return nil, 0, fmt.Errorf("查询角色与数据权限的关联关系发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	// }

	data, err = db.Query(sqls.QueryRoleDataPermission, map[string]interface{}{
		"sys_id":  req.SysID,
		"role_id": req.RoleID,
		// "start":   (req.PageIndex - 1) * req.PageSize,
		// "ps":      req.PageSize,
	})
	if err != nil {
		return nil, fmt.Errorf("查询角色与数据权限的关联关系发生错误(err:%v)", err)
	}
	return data, nil
}

//SaveRolePermission 保存角色与数据权限的关系
func (r *DbRole) SaveRolePermission(req model.RolePermissionReq) error {
	// if req.ID != 0 {
	// 	return r.updateRolePermission(req)
	// }

	db := components.Def.DB().GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("保存角色与数据权限的关系, 开启DB事务出错(err:%v)", err)
	}
	//删除原权限
	_, err = dbTrans.Execute(sqls.DelDataPermissionRoleAuth, map[string]interface{}{
		"role_id": req.RoleID,
		"sys_id":  req.SysID,
		// "table_name":     req.TableName,
		// "operate_action": req.OperateAction,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除[数据权限]－> 角色原权限发生错误(err:%v)", err)
	}

	selectAuth := req.Permissions
	if strings.EqualFold(selectAuth, "") {
		dbTrans.Commit()
		return nil
	}

	//添加新的数据权限 关系
	authArray := strings.Split(selectAuth, ",")
	for i := 0; i < len(authArray); i++ {
		_, err = dbTrans.Execute(sqls.AddRoleDataPermissionAuth, map[string]interface{}{
			"role_id": req.RoleID,
			"sys_id":  req.SysID,
			// "table_name":     req.TableName,
			// "operate_action": req.OperateAction,
			//"name":        req.Name,
			"permission_config_id": authArray[i],
		})
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加角色 -> 数据权限 关系发生错误(err:%v)", err)
		}
	}

	dbTrans.Commit()
	return nil
}

//UpdateRolePermission 更新数据权限关联关系
func (r *DbRole) updateRolePermission(req model.RolePermissionReq) error {
	db := components.Def.DB().GetRegularDB()
	_, err := db.Execute(sqls.UpdateRolePermission, map[string]interface{}{
		//"id":          req.ID,
		//"name":        req.Name,
		"permissions": req.Permissions,
	})
	if err != nil {
		return fmt.Errorf("更新数据权限关联关系时发生错误(err:%v)", err)
	}
	return nil
}

//ChangeRolePermissionStatus 改变状态
func (r *DbRole) ChangeRolePermissionStatus(id string, status int) error {
	db := components.Def.DB().GetRegularDB()

	_, err := db.Execute(sqls.ChangeRolePermissionStatus, map[string]interface{}{
		"id":     id,
		"status": status,
	})
	if err != nil {
		return fmt.Errorf("改变[数据权限]－> 角色与配置的关系时发生错误(err:%v)", err)
	}
	return nil
}

//DelRolePermission delete
func (r *DbRole) DelRolePermission(id string) error {
	db := components.Def.DB().GetRegularDB()
	_, err := db.Execute(sqls.DeleteRolePermission, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return fmt.Errorf("删除角色与数据权限配置信息的关系时发生错误(err:%v)", err)
	}
	return nil
}
