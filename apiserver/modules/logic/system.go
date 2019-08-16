package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/apiserver/modules/access/system"
)

type ISystemLogic interface {
	Get(ident string) (s db.QueryRow, err error)
}

type SystemLogic struct {
	c     component.IContainer
	cache system.ICacheSystem
	db    system.IDbSystem
}

func NewSystemLogic(c component.IContainer) *SystemLogic {
	return &SystemLogic{
		c:     c,
		cache: system.NewCacheSystem(c),
		db:    system.NewDbSystem(c),
	}
}

//Get 从数据库中获取系统信息
func (u *SystemLogic) Get(ident string) (s db.QueryRow, err error) {
	s, err = u.cache.Query(ident)
	if s == nil || err != nil {
		if s, err = u.db.Get(ident); err != nil {
			return nil, err
		}
		if err = u.cache.Save(s); err != nil {
			return nil, err
		}
	}
	return s, err
}
