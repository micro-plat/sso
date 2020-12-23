package user

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	cacheConst "github.com/micro-plat/sso/mgrserver/modules/const/cache"
	"github.com/micro-plat/sso/mgrserver/modules/model"
)

type ICacheUser interface {
	Query(s *model.QueryUserInput) (data db.QueryRows, total int, err error)
	Save(s *model.QueryUserInput, data db.QueryRows, total int) error
	Delete() error
	SaveUser(userID int, data db.QueryRow) error
	QueryUser(userID int) (data db.QueryRow, err error)
	DeleteUser() error
	QueryUserBySys(sysID, pi, ps int) (data db.QueryRows, counr int, err error)
	SaveUserBySys(sysID, pi, ps int, data db.QueryRows, count int) (err error)

	DeleteLockUserInfo(userName string) error
}

//CacheUser 控制用户登录
type CacheUser struct {
	cacheTime int
}

//NewCacheUser 创建对象
func NewCacheUser() *CacheUser {
	return &CacheUser{
		cacheTime: 3600 * 24,
	}
}

//DeleteLockUserInfo 解锁用户(删除缓存key)
func (l *CacheUser) DeleteLockUserInfo(userName string) error {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cacheConst.CacheLoginFailCount, "user_name", userName)
	return cache.Delete(key)
}

func (l *CacheUser) QueryUserBySys(sysID, pi, ps int) (data db.QueryRows, counr int, err error) {
	//从缓存中查询用户列表数据
	cache := components.Def.Cache().GetRegularCache()
	keyData := types.Translate(cacheConst.CacheUserSysFormat, "sysID", sysID, "pi", pi, "ps", ps)
	keyCount := types.Translate(cacheConst.CacheUserSysCountFormat, "sysID", sysID)
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
	return nmap, types.GetInt(c, 0), err
}

func (l *CacheUser) SaveUserBySys(sysID, pi, ps int, data db.QueryRows, count int) (err error) {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}

	cache := components.Def.Cache().GetRegularCache()
	keyData := types.Translate(cacheConst.CacheUserSysFormat, "sysID", sysID, "pi", pi, "ps", ps)
	keyCount := types.Translate(cacheConst.CacheUserSysCountFormat, "sysID", sysID)
	if err := cache.Set(keyData, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(keyCount, fmt.Sprint(count), l.cacheTime)
}

//Save 缓存用户列表信息
func (l *CacheUser) Save(s *model.QueryUserInput, data db.QueryRows, count int) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}

	cache := components.Def.Cache().GetRegularCache()
	keyData := types.Translate(cacheConst.CacheUserListFormat, "userName", s.UserName, "roleID", s.RoleID, "pageSize", s.PageSize, "pageIndex", s.PageIndex)
	keyCount := types.Translate(cacheConst.CacheUserListCountFormat, "userName", s.UserName, "roleID", s.RoleID)
	if err := cache.Set(keyData, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(keyCount, fmt.Sprint(count), l.cacheTime)
}

//Query 获取用户列表数据
func (l *CacheUser) Query(s *model.QueryUserInput) (data db.QueryRows, total int, err error) {
	//从缓存中查询用户列表数据
	cache := components.Def.Cache().GetRegularCache()
	keyData := types.Translate(cacheConst.CacheUserListFormat, "userName", s.UserName, "roleID", s.RoleID, "pageSize", s.PageSize, "pageIndex", s.PageIndex)
	keyCount := types.Translate(cacheConst.CacheUserListCountFormat, "userName", s.UserName, "roleID", s.RoleID)
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
	return nmap, types.GetInt(c, 0), err
}

//Delete 缓存用户列表信息删除
func (l *CacheUser) Delete() error {
	cache := components.Def.Cache().GetRegularCache()
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
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cacheConst.CacheUserFormat, "userID", userID)
	return cache.Set(key, string(buff), l.cacheTime)
}

//QueryUser 获取用户数据
func (l *CacheUser) QueryUser(userID int) (data db.QueryRow, err error) {
	//从缓存中查询用户数据
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cacheConst.CacheUserFormat, "userID", userID)
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
	cache := components.Def.Cache().GetRegularCache()
	return cache.Delete(cacheConst.CacheUserDeleteFormat)
}
