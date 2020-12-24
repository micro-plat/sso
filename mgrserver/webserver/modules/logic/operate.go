package logic

import (
	"github.com/micro-plat/sso/mgrserver/webserver/modules/access/operate"
	"github.com/micro-plat/sso/sdk/sso"
)

// IOperateLogic xx
type IOperateLogic interface {
	// 系统数据操作
	SysOperate(m *sso.LoginState, method string, r ...interface{}) (err error)
	// 角色数据操作
	RoleOperate(m *sso.LoginState, method string, r ...interface{}) (err error)
	// 菜单数据操作
	MenuOperate(m *sso.LoginState, method string, r ...interface{}) (err error)
	// 用户数据操作
	UserOperate(m *sso.LoginState, method string, r ...interface{}) (err error)
}

// OperateLogic 操作日志
type OperateLogic struct {
	db operate.IDbOperate
}

// NewOperateLogic xx
func NewOperateLogic() *OperateLogic {
	return &OperateLogic{
		db: operate.NewDbOperate(),
	}
}

// SysOperate 系统数据操作
func (o *OperateLogic) SysOperate(m *sso.LoginState, method string, r ...interface{}) (err error) {
	return o.db.SysOperate(m, method, r...)
}

// RoleOperate 角色数据操作
func (o *OperateLogic) RoleOperate(m *sso.LoginState, method string, r ...interface{}) (err error) {
	return o.db.RoleOperate(m, method, r...)
}

// MenuOperate 菜单数据操作
func (o *OperateLogic) MenuOperate(m *sso.LoginState, method string, r ...interface{}) (err error) {
	return o.db.MenuOperate(m, method, r...)
}

// UserOperate 用户数据操作
func (o *OperateLogic) UserOperate(m *sso.LoginState, method string, r ...interface{}) (err error) {
	return o.db.UserOperate(m, method, r...)
}
