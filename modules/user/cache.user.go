package user

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type ICacheUser interface {
	Query(s *QueryUserInput) (data db.QueryRows, count interface{}, err error)
	Save(s *QueryUserInput, data db.QueryRows, count interface{}) error
	Delete() error
	SaveUser(userID int, data db.QueryRow) error
	QueryUser(userID int) (data db.QueryRow, err error)
	DeleteUser() error
}

//CacheUser 控制用户登录
type CacheUser struct {
	c         component.IContainer
	cacheTime int
}

const (
	cacheUserListFormat      = "sso:user:list:"
	cacheUserListCountFormat = "sso:user:list-count:"
	cacheUserFormat          = "sso:user:info:"
)

//NewCacheUser 创建对象
func NewCacheUser(c component.IContainer) *CacheUser {
	return &CacheUser{
		c:         c,
		cacheTime: 3600 * 24,
	}
}

//Save 缓存用户列表信息
func (l *CacheUser) Save(s *QueryUserInput, data db.QueryRows, count interface{}) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	buff1 := count.(string)
	cache := l.c.GetRegularCache()
	key := cacheUserListFormat + s.ToString()
	key1 := cacheUserListCountFormat + s.ToString()
	if err := cache.Set(key, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(key1, string(buff1), l.cacheTime)
}

//Query 获取用户列表数据
func (l *CacheUser) Query(s *QueryUserInput) (data db.QueryRows, count interface{}, err error) {
	//从缓存中查询用户列表数据
	cache := l.c.GetRegularCache()
	key := cacheUserListFormat + s.ToString()
	key1 := cacheUserListCountFormat + s.ToString()
	v, err := cache.Get(key)
	if err != nil {
		return nil, nil, err
	}
	if v == "" {
		return nil, nil, fmt.Errorf("无用户列表数据")
	}
	nmap := make(db.QueryRows, 0)
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, nil, err
	}

	c, err := cache.Get(key1)
	if err != nil {
		return nil, nil, err
	}
	if c == "" {
		return nil, nil, fmt.Errorf("无用户列表数据")
	}
	ni := new(interface{})
	if err = json.Unmarshal([]byte(c), &ni); err != nil {
		return nil, nil, err
	}

	return nmap, ni, err
}

//Delete 缓存用户列表信息删除
func (l *CacheUser) Delete() error {
	cache := l.c.GetRegularCache()
	key := cacheUserListFormat + "*"
	return cache.Delete(key)
}

//SaveUser 缓存用户信息
func (l *CacheUser) SaveUser(userID int, data db.QueryRow) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := cacheUserFormat + strconv.Itoa(userID)
	return cache.Set(key, string(buff), l.cacheTime)
}

//QueryUser 获取用户数据
func (l *CacheUser) QueryUser(userID int) (data db.QueryRow, err error) {
	//从缓存中查询用户数据
	cache := l.c.GetRegularCache()
	key := cacheUserFormat + strconv.Itoa(userID)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, fmt.Errorf("无用户数据")
	}
	nmap := make(db.QueryRow)
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, err
	}
	return nmap, err
}

//DeleteUser 缓存用户信息删除
func (l *CacheUser) DeleteUser() error {
	cache := l.c.GetRegularCache()
	key := cacheUserFormat + "*"
	return cache.Delete(key)
}
