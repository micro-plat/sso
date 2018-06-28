package system

import (
	"encoding/json"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/transform"
)

const cacheFormat = "sso:system:info:{@sys_id}"

type ICacheSystem interface {
	Save(s db.QueryRow) (err error)
	Query(sysid int) (ls db.QueryRow, err error)
}

type CacheSystem struct {
	c         component.IContainer
	cacheTime int
}

//NewCacheSystem 创建登录对象
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
	key := transform.Translate(cacheFormat, "sys_id", s.GetInt("id"))
	return cache.Set(key, string(buff), l.cacheTime)
}

//Query 用户登录
func (l *CacheSystem) Query(sysid int) (ls db.QueryRow, err error) {
	//从缓存中查询用户数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheFormat, "sys_id", sysid)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, context.NewError(context.ERR_FORBIDDEN, "sys_id无效")
	}
	nmap := make(map[string]interface{})
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	cache.Delete(key)
	return nmap, err
}
