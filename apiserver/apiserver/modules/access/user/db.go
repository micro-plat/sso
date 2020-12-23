package user

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
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
	sys system.IDbSystem
}

//NewDBUser new
func NewDBUser() *DBUser {
	return &DBUser{
		sys: system.NewDbSystem(),
	}
}

//AddUser 新增用户
func (l *DBUser) AddUser(req model.UserInputNew) error {
	db := components.Def.DB().GetRegularDB()
	info, err := l.GetUserInfoByName(req.UserName)
	if err != nil {
		return err
	}
	if info != nil {
		return errs.NewError(commodel.ERR_USER_NAMEEXISTS, "此登录名[user_name]已被使用")
	}

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

//GetUserInfoByName 根据用户名查询用户信息
func (l *DBUser) GetUserInfoByName(userName string) (data db.QueryRow, err error) {
	db := components.Def.DB().GetRegularDB()
	result, _, _, err := db.Query(sqls.GetUserInfoByName, map[string]interface{}{"user_name": userName})
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() {
		return nil, nil
	}
	return result.Get(0), nil
}
