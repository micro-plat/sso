package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/lgapi/modules/access/operate"
	"github.com/micro-plat/sso/lgapi/modules/model"
)

// IOperateLogic xx
type IOperateLogic interface {
	// 登录行为
	LoginOperate(m *model.LoginState) (err error) //这个外api在用
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
