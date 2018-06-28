package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type ISystem interface {
	Query(sysid int) (s db.QueryRow, err error)
}

type System struct {
	c     component.IContainer
	cache ICacheSystem
	db    IDbSystem
}

func NewSystem(c component.IContainer) *System {
	return &System{
		c:     c,
		cache: NewCacheSystem(c),
		db:    NewDbSystem(c),
	}
}

//Query 从数据库中获取系统信息
func (m *System) Query(sysid int) (s db.QueryRow, err error) {
	//从缓存中获取用户信息，不存在时从数据库中获取
	s, err = m.cache.Query(sysid)
	if s == nil || err != nil {
		if s, err = m.db.Query(sysid); err != nil {
			return nil, err
		}
		//保存用户数据到缓存
		if err = m.cache.Save(s); err != nil {
			return nil, err
		}
	}
	return s, err
}
