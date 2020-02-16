package datapermission

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/sqls"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

type IDBDataPermission interface {
	GetUserDataPermissionConfigs(req model.DataPermissionGetReq) (result db.QueryRows, err error)
}

//DBDataPermission 数据权限
type DBDataPermission struct {
	c component.IContainer
}

//NewDBDataPermission new
func NewDBDataPermission(c component.IContainer) *DBDataPermission {
	return &DBDataPermission{
		c: c,
	}
}

//GetUserDataPermissionConfigs 获取某个用户的数据权限规则配置信息
func (l *DBDataPermission) GetUserDataPermissionConfigs(req model.DataPermissionGetReq) (result db.QueryRows, err error) {
	db := l.c.GetRegularDB()
	data, q, a, err := db.Query(sqls.QueryUserDataPermission, map[string]interface{}{
		"user_id":        req.UserID,
		"ident":          req.Ident,
		"table_name":     req.TableName,
		"operate_action": req.OperateAction,
	})
	if err != nil {
		return nil, fmt.Errorf("获取某个用户的数据权限规则配置信息 QueryUserDataPermission 出错: q:%s,a:%+v, err:%+v", q, a, err)
	}
	if data.IsEmpty() {
		return result, nil
	}
	configs, q, args, err := db.Query(sqls.QueryPermissionConfig, map[string]interface{}{
		"ids": data.Get(0).GetString("permissions"),
	})
	if err != nil {
		return result, fmt.Errorf("QueryPermissionConfig出错: sql:%s, args:%+v, err:%+v", q, args, err)
	}
	return configs, nil
}
