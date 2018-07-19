package function

import (
	"encoding/json"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/transform"
)

const (
	cacheFormat = "sso:system:func:{@sysid}"
	cacheFormatDel = "sso:system:func:*"
)


type ICacheSystemFunc interface {
	Save(sysID int,data []map[string]interface{}) (err error)
	Query(sysID int) (data []map[string]interface{},err error)
	Fresh()(err error)

}

type CacheSystemFunc struct {
	c         component.IContainer
	cacheTime int
}

//NewCacheSystem 创建登录对象
func NewCacheSystemFunc(c component.IContainer) *CacheSystemFunc {
	return &CacheSystemFunc{
		c:         c,
		cacheTime: 3600 * 24,
	}
}

//Save 缓存功能信息
func (l *CacheSystemFunc) Save(sysID int,data []map[string]interface{}) (err error) {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheFormat, "sysid",sysID)
	return cache.Set(key, string(buff), l.cacheTime)
}

//Query 获取缓存中功能数据
func (l *CacheSystemFunc) Query(sysID int) (data []map[string]interface{}, err error) {
	//从缓存中查询功能数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheFormat, "sysid", sysID)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, context.NewError(context.ERR_FORBIDDEN, "无数据")
	}
	var nmap []map[string]interface{}
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	return nmap, err
}

//Fresh 刷新缓存
func (l *CacheSystemFunc) Fresh()(err error){
	cache := l.c.GetRegularCache()
	return cache.Delete(cacheFormatDel)
}