package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrapi/modules/access/operate"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

// IOperateLogic xx
type IOperateLogic interface {
	// 登录行为
	LoginOperate(m *model.LoginState) (err error) //这个外api在用
	// 系统数据操作
	SysOperate(m *model.LoginState, method string, r ...interface{}) (err error)
	// 角色数据操作
	RoleOperate(m *model.LoginState, method string, r ...interface{}) (err error)
	// 菜单数据操作
	MenuOperate(m *model.LoginState, method string, r ...interface{}) (err error)
	// 用户数据操作
	UserOperate(m *model.LoginState, method string, r ...interface{}) (err error)
}

// OperateLogic 操作日志
type OperateLogic struct {
	c  component.IContainer
	db operate.IDbOperate
}

// NewOperateLogic xx
func NewOperateLogic(c component.IContainer) *OperateLogic {
	return &OperateLogic{
		c:  c,
		db: operate.NewDbOperate(c),
	}
}

// LoginOperate 登录行为
func (o *OperateLogic) LoginOperate(m *model.LoginState) (err error) {
	return o.db.LoginOperate(m)
}

// SysOperate 系统数据操作
func (o *OperateLogic) SysOperate(m *model.LoginState, method string, r ...interface{}) (err error) {
	return o.db.SysOperate(m, method, r...)
}

// RoleOperate 角色数据操作
func (o *OperateLogic) RoleOperate(m *model.LoginState, method string, r ...interface{}) (err error) {
	return o.db.RoleOperate(m, method, r...)
}

// MenuOperate 菜单数据操作
func (o *OperateLogic) MenuOperate(m *model.LoginState, method string, r ...interface{}) (err error) {
	return o.db.MenuOperate(m, method, r...)
}

// UserOperate 用户数据操作
func (o *OperateLogic) UserOperate(m *model.LoginState, method string, r ...interface{}) (err error) {

	return o.db.UserOperate(m, method, r...)
}
