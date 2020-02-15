package datapermission

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/sqls"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

type IDBDataPermission interface {
	GetUserDataPermissionConfigs(req model.DataPermissionGetReq) (result db.QueryRows, err error)
	GetAllUserInfoByUserRole(userID int, ident string) (string, error)
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
	fmt.Println(data.Get(0).GetString("permissions"))

	configs, q, args, err := db.Query(sqls.QueryPermissionConfig, map[string]interface{}{
		"ids": data.Get(0).GetString("permissions"),
	})
	fmt.Println(q)
	fmt.Println(args)

	if err != nil {
		return result, fmt.Errorf("QueryPermissionConfig出错: sql:%s, args:%+v, err:%+v", q, args, err)
	}

	fmt.Println(configs)
	fmt.Println("dddddd")
	return configs, nil
}

//GetAllUserInfoByUserRole 获取和当前用户同一个角色的用户ids
func (l *DBDataPermission) GetAllUserInfoByUserRole(userID int, ident string) (string, error) {
	db := l.c.GetRegularDB()
	userInfo, q, args, err := db.Query(sqls.GetAllUserInfoByUserRole, map[string]interface{}{
		"user_id": userID,
		"ident":   ident,
	})
	if err != nil {
		return "", fmt.Errorf("GetAllUserInfoByUserRole出错: sql:%s, args:%+v, err:%+v", q, args, err)
	}
	userIDArray := make([]string, 0)
	for _, item := range userInfo {
		userIDArray = append(userIDArray, item.GetString("user_id"))
	}
	return strings.Join(userIDArray, ","), nil
}
