package role

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/const/enum"
	"github.com/micro-plat/sso/modules/const/sql"
)

type IDbRole interface {
	Query(input *QueryRoleInput) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(roleID string, status int) (err error)
	Delete(roleID int) (err error)
	Edit(input *RoleEditInput) (err error)
	Add(input *RoleEditInput) (err error)
	Auth(input *RoleAuthInput) (err error)
	QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error)
}

//RoleEditInput 编辑角色参数
type RoleEditInput struct {
	RoleName string `form:"role_name" json:"role_name"`
	RoleID   int64  `form:"role_id" json:"role_id"`
	Status   int64  `form:"status" json:"status"`
	IsAdd    int64  `form:"is_add" json:"is_add"`
}

//RoleAuthInput 角色授权输入参数
type RoleAuthInput struct {
	RoleID     int64  `form:"role_id" json:"role_id"`
	SysID      int64  `form:"sys_id" json:"sys_id"`
	SelectAuth string `form:"selectauth" json:"selectauth"`
}

//QueryRoleInput 查询角色信息所需参数
type QueryRoleInput struct {
	PageIndex int    `form:"pi" json:"pi"`
	PageSize  int    `form:"ps" json:"ps"`
	RoleName  string `form:"role_name" json:"role_name"`
}

type DbRole struct {
	c component.IContainer
}

func NewDbRole(c component.IContainer) *DbRole {
	return &DbRole{
		c: c,
	}
}

//Query 获取角色信息列表
func (r *DbRole) Query(input *QueryRoleInput) (data db.QueryRows, count interface{}, err error) {
	db := r.c.GetRegularDB()
	params, err := types.Struct2Map(input)
	if err != nil {
		return nil, nil, fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	params["role_sql"] = " and t.name like '%" + input.RoleName + "%' "
	count, q, a, err := db.Scalar(sql.QueryRoleInfoListCount, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取角色信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryRoleInfoList, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取角色信息列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return data, count, nil
}

//ChangeStatus 修改角色状态
func (r *DbRole) ChangeStatus(roleID string, status int) (err error) {
	db := r.c.GetRegularDB()
	input := map[string]interface{}{
		"role_id": roleID,
	}
	switch status {
	case enum.Disabled:
		input["status"] = enum.Disabled
	case enum.Normal:
		input["status"] = enum.Normal
	}
	_, q, a, err := db.Execute(sql.UpdateRoleStatus, input)
	if err != nil {
		return fmt.Errorf("修改角色状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}

//Delete 删除角色
func (r *DbRole) Delete(roleID int) (err error) {
	db := r.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.DeleteRole, map[string]interface{}{
		"role_id": roleID,
	})
	if err != nil {
		return fmt.Errorf("删除角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	_, q, a, err = db.Execute(sql.DeleteRoleMenu, map[string]interface{}{
		"role_id": roleID,
	})
	if err != nil {
		return fmt.Errorf("删除角色菜单发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}

//Edit 编辑角色信息
func (r *DbRole) Edit(input *RoleEditInput) (err error) {
	db := r.c.GetRegularDB()
	params, err := types.Struct2Map(input)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}
	_, q, a, err := db.Execute(sql.EditRoleInfo, params)
	if err != nil {
		return fmt.Errorf("编辑角色信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}

//Add 添加角色
func (r *DbRole) Add(input *RoleEditInput) (err error) {
	db := r.c.GetRegularDB()
	params, err := types.Struct2Map(input)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	_, q, a, err := db.Execute(sql.AddRoleInfo, params)
	if err != nil {
		return fmt.Errorf("添加角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}

//Auth 添加角色权限
func (r *DbRole) Auth(input *RoleAuthInput) (err error) {
	db := r.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	//删除原权限
	_, q, a, err := dbTrans.Execute(sql.DelRoleAuth, map[string]interface{}{
		"role_id": input.RoleID,
		"sys_id":  input.SysID,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除角色原权限发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	if input.SelectAuth == "" {
		return nil
	}
	//添加新权限
	s := strings.Split(input.SelectAuth, ",")
	for i := 0; i < len(s); i++ {
		_, q, a, err := dbTrans.Execute(sql.AddRoleAuth, map[string]interface{}{
			"role_id":  input.RoleID,
			"sys_id":   input.SysID,
			"menu_id":  s[i],
			"sortrank": i + 1,
		})
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加角色权限发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	dbTrans.Commit()
	return nil
}

//QueryAuthMenu 查询角色菜单
func (r *DbRole) QueryAuthMenu(sysID int64, roleID int64) (results []map[string]interface{}, err error) {
	db := r.c.GetRegularDB()
	data, q, a, err := db.Query(sql.QuerySysMenucList, map[string]interface{}{
		"role_id": roleID,
		"sys_id":  sysID,
	})
	if err != nil {
		return nil, fmt.Errorf("获取菜单列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
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
