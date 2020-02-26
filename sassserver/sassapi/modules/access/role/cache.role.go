package role

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
)

type ICacheRole interface {
	Get(sysID, roleID int, path string) (data db.QueryRows, err error)
	SetPageAuth(sysID int, roleID int, path string, data db.QueryRows) error
	Query(s *model.QueryRoleInput) (data db.QueryRows, count int, err error)
	Save(s *model.QueryRoleInput, data db.QueryRows, count int) error
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
	cacheRoleListFormat      = "{sso}:role:list:{@roleName}-{@pageSize}-{@pageIndex}"
	cacheRoleListAll         = "{sso}:role:list:*"
	cacheRoleListCountFormat = "{sso}:role:list-count:{@roleName}"
	cacheRoleListCountAll    = "{sso}:role:list-count:*"
	cacheRoleFormat          = "{sso}:role:menu:{@roleID}-{@sysID}"
	cacheRoleAll             = "{sso}:role:menu:*"
	cachePageAuth            = "{sso}:page:auth:{@sysID}-{@roleID}-{@path}"
	cachePageAuthAll         = "{sso}:page:auth:*"
)

//NewCacheRole 创建角色缓存对象
func NewCacheRole(c component.IContainer) *CacheRole {
	return &CacheRole{
		c:         c,
		cacheTime: 3600 * 24,
	}
}
func (l *CacheRole) Get(sysID, roleID int, path string) (data db.QueryRows, err error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachePageAuth, "sysID", sysID, "roleID", roleID, "path", path)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, fmt.Errorf("无数据")
	}
	nmap := make(db.QueryRows, 0, 0)
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	return nmap, err
}

func (l *CacheRole) SetPageAuth(sysID int, roleID int, path string, data db.QueryRows) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachePageAuth, "sysID", sysID, "roleID", roleID, "path", path)
	return cache.Set(key, string(buff), l.cacheTime)
}

//Save 缓存角色列表信息
func (l *CacheRole) Save(s *model.QueryRoleInput, data db.QueryRows, count int) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheRoleListFormat, "roleName", s.RoleName, "pageSize", s.PageSize, "pageIndex", s.PageIndex)
	keyCount := transform.Translate(cacheRoleListCountFormat, "roleName", s.RoleName)
	if err := cache.Set(keyData, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(keyCount, fmt.Sprint(count), l.cacheTime)
}

//Query 获取角色列表数据
func (l *CacheRole) Query(s *model.QueryRoleInput) (data db.QueryRows, count int, err error) {
	//从缓存中查询角色列表数据
	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheRoleListFormat, "roleName", s.RoleName, "pageSize", s.PageSize, "pageIndex", s.PageIndex)
	keyCount := transform.Translate(cacheRoleListCountFormat, "roleName", s.RoleName)
	v, err := cache.Get(keyData)
	if err != nil {
		return nil, 0, err
	}
	if v == "" {
		return nil, 0, fmt.Errorf("无角色列表数据")
	}
	nmap := make(db.QueryRows, 0, 8)
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, 0, err
	}

	c, err := cache.Get(keyCount)
	if err != nil {
		return nil, 0, err
	}
	return nmap, types.GetInt(c, 0), err
}

//Delete 缓存角色列表信息删除
func (l *CacheRole) Delete() error {
	cache := l.c.GetRegularCache()
	if err := cache.Delete(cacheRoleListAll); err != nil {
		return err
	}
	if err := cache.Delete(cachePageAuthAll); err != nil {
		return err
	}
	return cache.Delete(cacheRoleListCountAll)
}

//SaveAuthMenu 缓存角色菜单信息
func (l *CacheRole) SaveAuthMenu(sysID int64, roleID int64, data []map[string]interface{}) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheRoleFormat, "roleID", roleID, "sysID", sysID)
	return cache.Set(key, string(buff), l.cacheTime)
}

//QueryAuthMenu 获取角色菜单数据
func (l *CacheRole) QueryAuthMenu(sysID int64, roleID int64) (data []map[string]interface{}, err error) {
	//从缓存中查询角色菜单数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheRoleFormat, "roleID", roleID, "sysID", sysID)
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
	return cache.Delete(cacheRoleAll)
}
