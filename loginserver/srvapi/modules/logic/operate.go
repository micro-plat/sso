package logic

import (
	"github.com/micro-plat/sso/loginserver/srvapi/modules/access/operate"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/model"
)

// IOperateLogic xx
type IOperateLogic interface {
	LoginOperate(m *model.LoginState) (err error)
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
