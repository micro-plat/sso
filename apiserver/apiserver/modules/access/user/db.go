package user

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/enum"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/sqls"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

type IDBUser interface {
	AddUser(req model.UserInputNew) error
}

// DBUser  用户管理
type DBUser struct {
	c component.IContainer
}

//NewDBUser new
func NewDBUser(c component.IContainer) *DBUser {
	return &DBUser{
		c: c,
	}
}

//AddUser 新增用户
func (l *DBUser) AddUser(req model.UserInputNew) error {
	db := l.c.GetRegularDB()
	params, err := types.Struct2Map(req)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	params["password"] = md5.Encrypt(enum.UserDefaultPassword)
	_, _, q, a, err := db.Executes(sqls.AddUserInfo, params)
	if err != nil {
		return fmt.Errorf("添加用户发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}
