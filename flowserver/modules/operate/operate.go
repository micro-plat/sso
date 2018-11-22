// 记录操作行为
package operate

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/flowserver/modules/member"
)

type IOperate interface {
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

type Operate struct {
	c  component.IContainer
	db IDbOperate
}

func NewOperate(c component.IContainer) *Operate {
	return &Operate{
		c:  c,
		db: NewDbOperate(c),
	}
}

// 登录行为
func (o *Operate) LoginOperate(m *member.LoginState) (err error) {
	return o.db.LoginOperate(m)
}

// 系统数据操作
func (o *Operate) SysOperate(m *member.LoginState, method string, r ...interface{}) (err error) {
	return o.db.SysOperate(m, method, r...)
}

// 角色数据操作
func (o *Operate) RoleOperate(m *member.LoginState, method string, r ...interface{}) (err error) {
	return o.db.RoleOperate(m, method, r...)
}

// 菜单数据操作
func (o *Operate) MenuOperate(m *member.LoginState, method string, r ...interface{}) (err error) {
	return o.db.MenuOperate(m, method, r...)
}

// 用户数据操作
func (o *Operate) UserOperate(m *member.LoginState, method string, r ...interface{}) (err error) {

	return o.db.UserOperate(m, method, r...)
}
