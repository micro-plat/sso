package system

import (
	"encoding/json"
	"fmt"
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
	SaveSysInfo(name string, status string, pi int, ps int, s db.QueryRows, count int) (err error)
	Query(ident string) (ls db.QueryRow, err error)
	QuerySysInfo(name string, status string, pi int, ps int) (ls db.QueryRows, count int, err error)
	FreshSysInfo() (err error)
}

type CacheSystem struct {
	cacheTime int
}

//NewCacheSystem 创建登录对象
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

//Query 用户登录
func (l *CacheSystem) Query(ident string) (ls db.QueryRow, err error) {
	//从缓存中查询用户数据
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cacheFormat, "ident", ident)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, errs.NewError(http.StatusForbidden, "ident无效")
	}
	nmap := make(map[string]interface{})
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	//	cache.Delete(key)
	return nmap, err
}

//SaveSysInfo  写入系统数据缓存
func (l *CacheSystem) SaveSysInfo(name string, status string, pi int, ps int, s db.QueryRows, count int) (err error) {
	buff, err := json.Marshal(s)
	if err != nil {
		return err
	}
	cache := components.Def.Cache().GetRegularCache()
	keyData := types.Translate(cacheFormatSys, "name", name, "status", status, "pi", pi, "ps", ps)
	keyCount := types.Translate(cacheFormatSysCount, "name", name, "status", status)
	if err := cache.Set(keyData, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(keyCount, fmt.Sprint(count), l.cacheTime)
}

//QuerySysInfo  获取缓存系统数据
func (l *CacheSystem) QuerySysInfo(name string, status string, pi int, ps int) (ls db.QueryRows, count int, err error) {
	//从缓存中获取系统数据
	cache := components.Def.Cache().GetRegularCache()
	keyData := types.Translate(cacheFormatSys, "name", name, "status", status, "pi", pi, "ps", ps)
	keyCount := types.Translate(cacheFormatSysCount, "name", name, "status", status)
	v, err := cache.Get(keyData)
	if err != nil {
		return nil, 0, err
	}
	if v == "" {
		return nil, 0, errs.NewError(http.StatusForbidden, "无数据")
	}
	var nmap db.QueryRows
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, 0, err
	}
	c, err := cache.Get(keyCount)
	if err != nil {
		return nil, 0, err
	}
	return nmap, types.GetInt(c, 0), err
}

//FreshSysInfo 刷新缓存
func (l *CacheSystem) FreshSysInfo() (err error) {
	cache := components.Def.Cache().GetRegularCache()
	return cache.Delete(cacheFormatSysDel)
}
