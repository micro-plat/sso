package base

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/mgrapi/modules/const/sqls"
)

type IBase interface {
	QueryUserRoleList() (data db.QueryRows, err error)
	QuerySysList() (data db.QueryRows, err error)
}

type Base struct {
	c component.IContainer
}

func NewBase(c component.IContainer) *Base {
	return &Base{
		c: c,
	}
}

//QueryUserRoleList 获取用户角色列表
func (u *Base) QueryUserRoleList() (data db.QueryRows, err error) {
	db := u.c.GetRegularDB()
	data, q, a, err := db.Query(sqls.GetUserRoleList, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取用户角色列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}

//QuerySysList 获取系统列表
func (u *Base) QuerySysList() (data db.QueryRows, err error) {
	db := u.c.GetRegularDB()
	data, q, a, err := db.Query(sqls.GetSysList, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取用户角色列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}
