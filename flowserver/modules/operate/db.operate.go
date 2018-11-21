package operate

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/modules/const/sql"
	"github.com/micro-plat/sso/modules/member"
)

type IDbOperate interface {
	// 登录行为
	LoginOperate(m *member.LoginState) (err error)
	// 系统数据操作
	SysOperate(m *member.LoginState, method string, r ...interface{}) (err error)
	// 角色数据操作
	RoleOperate(m *member.LoginState, method string, r ...interface{}) (err error)
	// 菜单数据操作
	MenuOperate(m *member.LoginState, method string, r ...interface{}) (err error)
	// 用户数据操作
	UserOperate(m *member.LoginState, method string, r ...interface{}) (err error)
}

type DbOperate struct {
	c component.IContainer
}

func NewDbOperate(c component.IContainer) *DbOperate {
	return &DbOperate{
		c: c,
	}
}

// 登录行为
func (d *DbOperate) LoginOperate(m *member.LoginState) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    10,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", "用户登录", m.UserName),
	}
	_, q, a, err := db.Execute(sql.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加登录行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil

}

// 系统数据操作
func (d *DbOperate) SysOperate(m *member.LoginState, method string, r ...interface{}) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    20,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", method, r),
	}
	_, q, a, err := db.Execute(sql.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加系统操作行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

// 角色数据操作
func (d *DbOperate) RoleOperate(m *member.LoginState, method string, r ...interface{}) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    30,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", method, r),
	}
	_, q, a, err := db.Execute(sql.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加角色操作行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

// 菜单数据操作
func (d *DbOperate) MenuOperate(m *member.LoginState, method string, r ...interface{}) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    40,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", method, r),
	}
	_, q, a, err := db.Execute(sql.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加菜单操作行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

// 用户数据操作
func (d *DbOperate) UserOperate(m *member.LoginState, method string, r ...interface{}) (err error) {
	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		"type":    50,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf("{\"desc\":%s,\"data\":%v}", method, r),
	}
	_, q, a, err := db.Execute(sql.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加用户操作行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}
