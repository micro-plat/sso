package system

import (
	"encoding/json"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/transform"
)

const (
	cacheFormat         = "{sso}:system:info:{@ident}"
	cacheFormatSys      = "{sso}:system:info:{@name}-{@status}-{@pi}-{@ps}"
	cacheFormatSysDel   = "{sso}:system:info:*"
	cacheFormatSysCount = "{sso}:system:info:{@name}-{@status}"
)

type ICacheSystem interface {
	Save(s db.QueryRow) (err error)
	Query(ident string) (ls db.QueryRow, err error)
}

type CacheSystem struct {
	c         component.IContainer
	cacheTime int
}

//NewCacheSystem 系统信息缓存对象
func NewCacheSystem(c component.IContainer) *CacheSystem {
	return &CacheSystem{
		c:         c,
		cacheTime: 3600 * 24,
	}
}

//Save 缓存用户信息
func (l *CacheSystem) Save(s db.QueryRow) (err error) {
	buff, err := json.Marshal(s)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheFormat, "ident", s.GetString("ident"))
	return cache.Set(key, string(buff), l.cacheTime)
}

//Query 获取系统信息
func (l *CacheSystem) Query(ident string) (ls db.QueryRow, err error) {
	//从缓存中查询用户数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheFormat, "ident", ident)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, context.NewError(context.ERR_FORBIDDEN, "ident无效")
	}
	nmap := make(map[string]interface{})
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	return nmap, err
}
