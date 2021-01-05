package datapermission

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/const/sqls"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/model"
)

type IDBDataPermission interface {
	GetUserDataPermissionConfigs(req model.DataPermissionGetReq) (result db.QueryRows, err error)
}

//DBDataPermission 数据权限
type DBDataPermission struct {
}

//NewDBDataPermission new
func NewDBDataPermission() *DBDataPermission {
	return &DBDataPermission{}
}

//GetUserDataPermissionConfigs 获取某个用户的数据权限规则配置信息
func (l *DBDataPermission) GetUserDataPermissionConfigs(req model.DataPermissionGetReq) (result db.QueryRows, err error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryUserDataPermission, map[string]interface{}{
		"user_id":        req.UserID,
		"ident":          req.Ident,
		"table_name":     req.TableName,
		"operate_action": req.OperateAction,
	})
	if err != nil {
		return nil, fmt.Errorf("获取某个用户的数据权限规则配置信息 QueryUserDataPermission 出错: err:%+v", err)
	}
	return data, nil

	// if data.IsEmpty() {
	// 	return result, nil
	// }
	// configs, err := db.Query(sqls.QueryPermissionConfig, map[string]interface{}{
	// 	"ids": data.Get(0).GetString("permissions"),
	// })
	// if err != nil {
	// 	return result, fmt.Errorf("QueryPermissionConfig出错: err:%+v", err)
	// }
	// return configs, nil
}
