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
	commodel "github.com/micro-plat/sso/loginserver/loginapi/modules/model"
	"github.com/micro-plat/sso/sso/errorcode"

	"github.com/micro-plat/sso/loginserver/srvapi/modules/access/system"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/const/enum"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/const/sqls"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/model"
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
		return errs.NewError(errorcode.ERR_USER_NAMEEXISTS, "此登录名[user_name]已被使用")
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
	userID, _, err := dbTrans.Executes(sqls.AddUserInfo, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加用户发生错误(err:%v)", err)
	}

	l.adapterRoleID(req)
	if req.RoleID != 0 {

		params["user_id"] = userID

		systemInfo, err := l.sys.Get(req.TargetIdent)
		if err != nil {
			dbTrans.Rollback()
			return err
		}

		params["role_id"] = req.RoleID
		params["sys_id"] = systemInfo.GetInt("id")
		_, err = dbTrans.Execute(sqls.AddUserRole, params)
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("关联用户角色发生错误(err:%v)", err)
		}
	}

	dbTrans.Commit()
	return nil
}

func (l *DBUser) adapterRoleID(req model.UserInputNew) {
	if req.RoleID != 0 {
		return
	}
	conf := commodel.GetLoginConf() //取默认配置的角色
	if strings.EqualFold(conf.AddUserUseDefaultRole, "") {
		return
	}
	var roleConfig map[string]int
	err := json.Unmarshal([]byte(conf.AddUserUseDefaultRole), &roleConfig)
	if err != nil {
		return
	}
	roleID, ok := roleConfig[req.TargetIdent]
	if !ok {
		return
	}
	req.RoleID = roleID
}

//GetUserInfoByName 根据用户名查询用户信息
func (l *DBUser) GetUserInfoByName(userName string) (data db.QueryRow, err error) {
	db := components.Def.DB().GetRegularDB()
	result, err := db.Query(sqls.GetUserInfoByName, map[string]interface{}{"user_name": userName})
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() {
		return nil, nil
	}
	return result.Get(0), nil
}
