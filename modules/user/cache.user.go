package user

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"
	cacheConst "github.com/micro-plat/sso/modules/const/cache"
)

type ICacheUser interface {
	Query(s *QueryUserInput) (data db.QueryRows, total int, err error)
	Save(s *QueryUserInput, data db.QueryRows, total int) error
	Delete() error
	SaveUser(userID int, data db.QueryRow) error
	QueryUser(userID int) (data db.QueryRow, err error)
	DeleteUser() error
	SetEmail(Guid string, email string) (err error)
	GetEmail(Guid string) (email string, err error)
	QueryUserBySys(sysID, pi, ps int) (data db.QueryRows, counr int, err error)
	SaveUserBySys(sysID, pi, ps int, data db.QueryRows, count int) (err error)
}

//CacheUser 控制用户登录
type CacheUser struct {
	c         component.IContainer
	cacheTime int
}

//NewCacheUser 创建对象
func NewCacheUser(c component.IContainer) *CacheUser {
	return &CacheUser{
		c:         c,
		cacheTime: 3600 * 24,
	}
}

func (l *CacheUser) QueryUserBySys(sysID, pi, ps int) (data db.QueryRows, counr int, err error) {
	//从缓存中查询用户列表数据
	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheConst.CacheUserSysFormat, "sysID", sysID, "pi", pi, "ps", ps)
	keyCount := transform.Translate(cacheConst.CacheUserSysCountFormat, "sysID", sysID)
	v, err := cache.Get(keyData)
	if err != nil {
		return nil, 0, err
	}
	if v == "" {
		return nil, 0, fmt.Errorf("无用户列表数据")
	}
	nmap := make(db.QueryRows, 0)
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, 0, err
	}

	c, err := cache.Get(keyCount)
	if err != nil {
		return nil, 0, err
	}
	return nmap, types.ToInt(c, 0), err
}

func (l *CacheUser) SaveUserBySys(sysID, pi, ps int, data db.QueryRows, count int) (err error) {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}

	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheConst.CacheUserSysFormat, "sysID", sysID, "pi", pi, "ps", ps)
	keyCount := transform.Translate(cacheConst.CacheUserSysCountFormat, "sysID", sysID)
	if err := cache.Set(keyData, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(keyCount, fmt.Sprint(count), l.cacheTime)
}

//Save 缓存用户列表信息
func (l *CacheUser) Save(s *QueryUserInput, data db.QueryRows, count int) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}

	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheConst.CacheUserListFormat, "userName", s.UserName, "roleID", s.RoleID, "pageSize", s.PageSize, "pageIndex", s.PageIndex)
	keyCount := transform.Translate(cacheConst.CacheUserListCountFormat, "userName", s.UserName, "roleID", s.RoleID)
	if err := cache.Set(keyData, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(keyCount, fmt.Sprint(count), l.cacheTime)
}

//Query 获取用户列表数据
func (l *CacheUser) Query(s *QueryUserInput) (data db.QueryRows, total int, err error) {
	//从缓存中查询用户列表数据
	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheConst.CacheUserListFormat, "userName", s.UserName, "roleID", s.RoleID, "pageSize", s.PageSize, "pageIndex", s.PageIndex)
	keyCount := transform.Translate(cacheConst.CacheUserListCountFormat, "userName", s.UserName, "roleID", s.RoleID)
	v, err := cache.Get(keyData)
	if err != nil {
		return nil, 0, err
	}
	if v == "" {
		return nil, 0, fmt.Errorf("无用户列表数据")
	}
	nmap := make(db.QueryRows, 0)
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, 0, err
	}

	c, err := cache.Get(keyCount)
	if err != nil {
		return nil, 0, err
	}
	return nmap, types.ToInt(c, 0), err
}

//Delete 缓存用户列表信息删除
func (l *CacheUser) Delete() error {
	cache := l.c.GetRegularCache()
	if err := cache.Delete(cacheConst.CacheUserListAll); err != nil {
		return err
	}
	return cache.Delete(cacheConst.CacheUserListCountAll)
}

//SaveUser 缓存用户信息
func (l *CacheUser) SaveUser(userID int, data db.QueryRow) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheConst.CacheUserFormat, "userID", userID)
	return cache.Set(key, string(buff), l.cacheTime)
}

//QueryUser 获取用户数据
func (l *CacheUser) QueryUser(userID int) (data db.QueryRow, err error) {
	//从缓存中查询用户数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheConst.CacheUserFormat, "userID", userID)
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
	return cache.Delete(cacheConst.CacheUserDeleteFormat)
}

func (l *CacheUser) SetEmail(Guid string, email string) (err error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheConst.CacheEmail, "guid", Guid)
	return cache.Set(key, email, cacheConst.CacheEamilOutTime)
}

func (l *CacheUser) GetEmail(Guid string) (email string, err error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheConst.CacheEmail, "guid", Guid)
	email, err = cache.Get(key)
	return
}
