package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/lgapi/modules/access/system"
)

// ISystemLogic xx
type ISystemLogic interface {
	QueryUserSystem(userId int64) (db.QueryRows, error)
}

// SystemLogic 操作日志
type SystemLogic struct {
	c     component.IContainer
	dbSys system.IDbSystem
}

// NewSystemLogic xx
func NewSystemLogic(c component.IContainer) *SystemLogic {
	return &SystemLogic{
		c:     c,
		dbSys: system.NewDbSystem(c),
	}
}

// QueryUserSystem 返回用户可用的子系统信息
func (s *SystemLogic) QueryUserSystem(userId int64) (db.QueryRows, error) {
	return s.dbSys.QueryUserSystem(userId)
}
