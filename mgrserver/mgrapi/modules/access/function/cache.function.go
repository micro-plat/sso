package function

import (
	"encoding/json"
	"net/http"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

const (
	cacheFormat    = "{sso}:system:func:{@sysid}"
	cacheFormatDel = "{sso}:system:func:*"
	cacheMenuAll   = "{sso}:role:menu:*"
)

type ICacheSystemFunc interface {
	Save(sysID int, data []map[string]interface{}) (err error)
	Query(sysID int) (data []map[string]interface{}, err error)
	Fresh() (err error)
}

type CacheSystemFunc struct {
	cacheTime int
}

//NewCacheSystem 创建登录对象
func NewCacheSystemFunc() *CacheSystemFunc {
	return &CacheSystemFunc{
		cacheTime: 3600 * 24,
	}
}

//Save 缓存功能信息
func (l *CacheSystemFunc) Save(sysID int, data []map[string]interface{}) (err error) {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := components.Def.Cache().GetRegularCache("redis")
	key := types.Translate(cacheFormat, "sysid", sysID)
	return cache.Set(key, string(buff), l.cacheTime)
}

//Query 获取缓存中功能数据
func (l *CacheSystemFunc) Query(sysID int) (data []map[string]interface{}, err error) {
	//从缓存中查询功能数据
	cache := components.Def.Cache().GetRegularCache("redis")
	key := types.Translate(cacheFormat, "sysid", sysID)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, errs.NewError(http.StatusForbidden, "无数据")
	}
	var nmap []map[string]interface{}
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	return nmap, err
}

//Fresh 刷新缓存
func (l *CacheSystemFunc) Fresh() (err error) {
	cache := components.Def.Cache().GetRegularCache("redis")
	_ = cache.Delete(cacheMenuAll)
	return cache.Delete(cacheFormatDel)
}
