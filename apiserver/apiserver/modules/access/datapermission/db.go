package datapermission

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/sqls"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

type IDBDataPermission interface {
	GetUserDataPermission(req model.DataPermissionGetReq) (result string, err error)
	SyncDataPermission(req model.DataPermissionSyncReq) error
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

//GetUserDataPermission 获取某个用户的数据权限数据
func (l *DBDataPermission) GetUserDataPermission(req model.DataPermissionGetReq) (result string, err error) {
	db := l.c.GetRegularDB()
	data, q, a, err := db.Query(sqls.QueryUserDataPermission, map[string]interface{}{
		"user_id": req.UserID,
		"ident":   req.Ident,
		"type":    req.Type,
	})
	if err != nil {
		return "", fmt.Errorf("获取某个用户的数据权限数据 GetUserDataPermission 出错: q:%s,a:%+v, err:%+v", q, a, err)
	}
	if data.IsEmpty() {
		return "", nil
	}
	var temp []string
	for _, item := range data {
		temp = append(temp, item.GetString("value"))
	}
	return strings.Join(temp, ","), nil
}

// SyncDataPermission 同步数据权限数据
func (l *DBDataPermission) SyncDataPermission(req model.DataPermissionSyncReq) error {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QuerySystemInfo, map[string]interface{}{
		"ident": req.Ident,
	})
	if err != nil {
		return context.NewError(context.ERR_SERVICE_UNAVAILABLE, "获取系统信息出错")
	}

	if data.IsEmpty() {
		return context.NewError(context.ERR_SERVICE_UNAVAILABLE, "ident传入有误或者系统被禁用")
	}

	//增加权限数据
	first := data.Get(0)
	_, q, a, err := db.Execute(sqls.AddDataPermission, map[string]interface{}{
		"ident":     req.Ident,
		"sys_id":    first.GetString("id"),
		"name":      req.Name,
		"type":      req.Type,
		"type_name": req.TypeName,
		"value":     req.Value,
		"remark":    req.Remark,
	})
	if err != nil {
		return fmt.Errorf("SyncDataPermission 同步数据发生错误, q:%s, a:%+v, err:%+v", q, a, err)
	}

	//增加一个类型全局数据
	db.Execute(sqls.AddDefaultDataPermissionInfo, map[string]interface{}{
		"ident":     req.Ident,
		"sys_id":    first.GetString("id"),
		"type":      req.Type,
		"type_name": req.TypeName,
	})

	return nil
}
