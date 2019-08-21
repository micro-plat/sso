package operate

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/sqls"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

type IDbOperate interface {
	// 系统数据操作
	SysOperate(m *model.LoginState, method string, r ...interface{}) (err error)
	// 角色数据操作
	RoleOperate(m *model.LoginState, method string, r ...interface{}) (err error)
	// 菜单数据操作
	MenuOperate(m *model.LoginState, method string, r ...interface{}) (err error)
	// 用户数据操作
	UserOperate(m *model.LoginState, method string, r ...interface{}) (err error)
}

type DbOperate struct {
	c component.IContainer
}

func NewDbOperate(c component.IContainer) *DbOperate {
	return &DbOperate{
		c: c,
	}
}

//SysOperate 系统数据操作
func (d *DbOperate) SysOperate(m *model.LoginState, method string, r ...interface{}) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    20,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", method, r),
	}
	_, q, a, err := db.Execute(sqls.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加系统操作行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

//RoleOperate 角色数据操作
func (d *DbOperate) RoleOperate(m *model.LoginState, method string, r ...interface{}) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    30,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", method, r),
	}
	_, q, a, err := db.Execute(sqls.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加角色操作行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

//MenuOperates 菜单数据操作
func (d *DbOperate) MenuOperate(m *model.LoginState, method string, r ...interface{}) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    40,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", method, r),
	}
	_, q, a, err := db.Execute(sqls.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加菜单操作行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

//UserOperate 用户数据操作
func (d *DbOperate) UserOperate(m *model.LoginState, method string, r ...interface{}) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    50,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", method, r),
	}
	_, q, a, err := db.Execute(sqls.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加用户操作行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}
