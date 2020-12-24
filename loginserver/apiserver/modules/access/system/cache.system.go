package system

import (
	"encoding/json"
	"net/http"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
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
	cacheTime int
}

//NewCacheSystem 系统信息缓存对象
func NewCacheSystem() *CacheSystem {
	return &CacheSystem{
		cacheTime: 3600 * 24,
	}
}

//Save 缓存用户信息
func (l *CacheSystem) Save(s db.QueryRow) (err error) {
	buff, err := json.Marshal(s)
	if err != nil {
		return err
	}
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cacheFormat, "ident", s.GetString("ident"))
	return cache.Set(key, string(buff), l.cacheTime)
}

//Query 获取系统信息
func (l *CacheSystem) Query(ident string) (ls db.QueryRow, err error) {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cacheFormat, "ident", ident)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, errs.NewError(http.StatusBadRequest, "ident无效")
	}
	nmap := make(map[string]interface{})
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	return nmap, err
}
