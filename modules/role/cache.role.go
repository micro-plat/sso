package role

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type ICacheRole interface {
	Query(s *QueryRoleInput) (data db.QueryRows, count interface{}, err error)
	Save(s *QueryRoleInput, data db.QueryRows, count interface{}) error
	Delete() error
	SaveAuthMenu(sysID int64, roleID int64, data []map[string]interface{}) error
	QueryAuthMenu(sysID int64, roleID int64) (data []map[string]interface{}, err error)
	DeleteAuthMenu() error
}

//CacheRole 控制用户角色缓存
type CacheRole struct {
	c         component.IContainer
	cacheTime int
}

const (
	cacheRoleListFormat      = "sso:role:list:"
	cacheRoleListCountFormat = "sso:role:list-count:"
	cacheRoleFormat          = "sso:role:menu:"
)

//NewCacheRole 创建角色缓存对象
func NewCacheRole(c component.IContainer) *CacheRole {
	return &CacheRole{
		c:         c,
		cacheTime: 3600 * 24,
	}
}

//Save 缓存角色列表信息
func (l *CacheRole) Save(s *QueryRoleInput, data db.QueryRows, count interface{}) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	buff1 := count.(string)
	cache := l.c.GetRegularCache()
	key := cacheRoleListFormat + s.ToString()
	key1 := cacheRoleListCountFormat + s.ToString()
	if err := cache.Set(key, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(key1, string(buff1), l.cacheTime)
}

//Query 获取角色列表数据
func (l *CacheRole) Query(s *QueryRoleInput) (data db.QueryRows, count interface{}, err error) {
	//从缓存中查询角色列表数据
	cache := l.c.GetRegularCache()
	key := cacheRoleListFormat + s.ToString()
	key1 := cacheRoleListCountFormat + s.ToString()
	v, err := cache.Get(key)
	if err != nil {
		return nil, nil, err
	}
	if v == "" {
		return nil, nil, fmt.Errorf("无角色列表数据")
	}
	if err = json.Unmarshal([]byte(v), &data); err != nil {
		return nil, nil, err
	}
	nmap := make(db.QueryRows, 0, 8)
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, nil, err
	}

	c, err := cache.Get(key1)
	if err != nil {
		return nil, nil, err
	}
	if c == "" {
		return nil, nil, fmt.Errorf("无角色列表数据")
	}
	if err = json.Unmarshal([]byte(v), &data); err != nil {
		return nil, nil, err
	}
	ni := new(interface{})
	if err = json.Unmarshal([]byte(c), &ni); err != nil {
		return nil, nil, err
	}

	return nmap, ni, err
}

//Delete 缓存角色列表信息删除
func (l *CacheRole) Delete() error {
	cache := l.c.GetRegularCache()
	key := cacheRoleListFormat + "*"
	return cache.Delete(key)
}

//SaveAuthMenu 缓存角色菜单信息
func (l *CacheRole) SaveAuthMenu(sysID int64, roleID int64, data []map[string]interface{}) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := cacheRoleFormat + strconv.FormatInt(roleID, 10) + "-" + strconv.FormatInt(sysID, 10)
	return cache.Set(key, string(buff), l.cacheTime)
}

//QueryAuthMenu 获取角色菜单数据
func (l *CacheRole) QueryAuthMenu(sysID int64, roleID int64) (data []map[string]interface{}, err error) {
	//从缓存中查询角色菜单数据
	cache := l.c.GetRegularCache()
	key := cacheRoleFormat + strconv.FormatInt(roleID, 10) + "-" + strconv.FormatInt(sysID, 10)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, fmt.Errorf("无角色菜单列表数据")
	}
	nmap := make([]map[string]interface{}, 0)
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	return nmap, err
}

//DeleteAuthMenu 缓存角色菜单信息删除
func (l *CacheRole) DeleteAuthMenu() error {
	cache := l.c.GetRegularCache()
	key := cacheRoleFormat + "*"
	return cache.Delete(key)
}
