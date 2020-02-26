package base

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/const/sqls"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/config"
)

type IBase interface {
	QueryUserRoleList(belongID, belongType int) (data db.QueryRows, err error)
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
//belongType 所属类型(0:加油站，1:公司),由于现在只有几个默认角色，还没有区分公司，
//belongID(所属id,如果是加油站就是加油站id, 如果是公司就是公司id) 如果改成加油站有自己的角色那就必须用上这个参数
func (u *Base) QueryUserRoleList(belongID, belongType int) (data db.QueryRows, err error) {
	db := u.c.GetRegularDB(config.DbName)
	data, q, a, err := db.Query(sqls.GetUserRoleList, map[string]interface{}{
		"belong_id":   belongID,
		"belong_type": belongType,
	})
	if err != nil {
		return nil, fmt.Errorf("获取用户角色列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}

//QuerySysList 获取系统列表
func (u *Base) QuerySysList() (data db.QueryRows, err error) {
	db := u.c.GetRegularDB(config.DbName)
	data, q, a, err := db.Query(sqls.GetSysList, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取用户角色列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}
