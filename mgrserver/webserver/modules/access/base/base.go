package base

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/const/sqls"
)

type IBase interface {
	QueryUserRoleList() (data db.QueryRows, err error)
	QuerySysList() (data db.QueryRows, err error)
	GetPermissTypes(sysID string) (data db.QueryRows, err error)
}

type Base struct {
}

func NewBase() *Base {
	return &Base{}
}

//QueryUserRoleList 获取用户角色列表
func (u *Base) QueryUserRoleList() (data db.QueryRows, err error) {
	db := components.Def.DB().GetRegularDB()
	data, q, a, err := db.Query(sqls.GetUserRoleList, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取用户角色列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, nil
}

//QuerySysList 获取系统列表
func (u *Base) QuerySysList() (data db.QueryRows, err error) {
	db := components.Def.DB().GetRegularDB()
	data, q, a, err := db.Query(sqls.GetSysList, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取用户角色列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}

//GetPermissTypes typeslist
func (u *Base) GetPermissTypes(sysID string) (data db.QueryRows, err error) {
	db := components.Def.DB().GetRegularDB()
	data, q, a, err := db.Query(sqls.GetPermissTypes, map[string]interface{}{
		"sys_id": sysID,
	})
	if err != nil {
		return nil, fmt.Errorf("查询某个系统下面所有的数据权限类型(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}
