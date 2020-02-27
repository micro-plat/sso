package user

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/access/system"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/enum"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/sqls"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
	commodel "github.com/micro-plat/sso/common/module/model"
)

type IDBUser interface {
	AddUser(req model.UserInputNew) error
}

// DBUser  用户管理
type DBUser struct {
	c   component.IContainer
	sys system.IDbSystem
}

//NewDBUser new
func NewDBUser(c component.IContainer) *DBUser {
	return &DBUser{
		c:   c,
		sys: system.NewDbSystem(c),
	}
}

//AddUser 新增用户
func (l *DBUser) AddUser(req model.UserInputNew) error {
	db := l.c.GetRegularDB()
	params, err := types.Struct2Map(req)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	params["password"] = md5.Encrypt(enum.UserDefaultPassword)
	userID, _, q, a, err := dbTrans.Executes(sqls.AddUserInfo, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加用户发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	params["user_id"] = userID
	conf := commodel.GetConf(l.c) //取默认配置的角色
	if strings.EqualFold(conf.AddUserUseDefaultRole, "") {
		dbTrans.Commit()
		return nil
	}
	var roleConfig map[string]int
	err = json.Unmarshal([]byte(conf.AddUserUseDefaultRole), &roleConfig)
	if err != nil {
		dbTrans.Commit()
		return fmt.Errorf("新增用户取默认角色配置出错: %+v", err)
	}
	roleID, flag := roleConfig[req.TargetIdent]
	if !flag {
		dbTrans.Commit()
		return nil
	}

	systemInfo, err := l.sys.Get(req.TargetIdent)
	if err != nil {
		dbTrans.Rollback()
		return err
	}

	params["role_id"] = roleID
	params["sys_id"] = systemInfo.GetInt("id")
	_, q, a, err = dbTrans.Execute(sqls.AddUserRole, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("关联用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}
