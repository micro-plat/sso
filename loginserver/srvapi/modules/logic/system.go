package logic

import (
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/access/system"
)

type ISystemLogic interface {
	Get(ident string) (s types.IXMap, err error)
}

type SystemLogic struct {
	cache system.ICacheSystem
	db    system.IDbSystem
}

func NewSystemLogic() *SystemLogic {
	return &SystemLogic{
		cache: system.NewCacheSystem(),
		db:    system.NewDbSystem(),
	}
}

//Get 从数据库中获取系统信息
func (u *SystemLogic) Get(ident string) (s types.IXMap, err error) {
	if s, err = u.db.Get(ident); err != nil {
		return nil, err
	}
	return s, err
}
