package logic

import (
	"github.com/micro-plat/sso/loginserver/webserver/modules/access/operate"
	"github.com/micro-plat/sso/loginserver/webserver/modules/model"
)

// IOperateLogic xx
type IOperateLogic interface {
	// 登录行为
	LoginOperate(m *model.LoginState) (err error) //这个外api在用
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

// LoginOperate 登录行为
func (o *OperateLogic) LoginOperate(m *model.LoginState) (err error) {
	return o.db.LoginOperate(m)
}
