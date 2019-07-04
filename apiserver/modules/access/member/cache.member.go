package member

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/apiserver/modules/const/enum"
	"github.com/micro-plat/sso/apiserver/modules/model"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"
)

type ICacheMember interface {
	Query(u string, ident string) (ls *model.MemberState, err error)
	Save(s *model.MemberState) error
	SetLoginSuccess(u string) error
	SetLoginFail(u string) (int, error)
	QueryAuth(sysID, userID int64) (err error)
	SaveAuth(sysID, userID int64, data db.QueryRows) (err error)
}

//CacheMember 控制用户登录
type CacheMember struct {
	c          component.IContainer
	maxFailCnt int
	cacheTime  int
}

const (
	cacheFormat     = "{sso}:login:state-info:{@userName}-{@ident}"
	cacheCodeFormat = "{sso}:login:state-code:{@userName}-{@ident}"
	lockFormat      = "{sso}:login:state-locker:{@userName}"

	cacheSysAuth = "{sso}:sys:auth:{@sysID}-{@userID}"
)

//NewCacheMember 创建登录对象
func NewCacheMember(c component.IContainer) *CacheMember {
	return &CacheMember{
		c:          c,
		maxFailCnt: 5,
		cacheTime:  3600 * 24,
	}
}

// 缓存系统权限信息
func (l *CacheMember) QueryAuth(sysID, userID int64) (err error) {
	//从缓存中获取系统数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheSysAuth, "sysID", sysID, "userID", userID)
	v, err := cache.Get(key)
	if err != nil {
		return err
	}
	if v == "" {
		return context.NewError(context.ERR_FORBIDDEN, "无数据")
	}
	return nil
}

// 查询系统权限信息
func (l *CacheMember) SaveAuth(sysID, userID int64, data db.QueryRows) (err error) {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheSysAuth, "sysID", sysID, "userID", userID)
	return cache.Set(key, string(buff), l.cacheTime)
}

//SetLoginSuccess 设置为登录成功
func (l *CacheMember) SetLoginSuccess(u string) error {
	cache := l.c.GetRegularCache()
	key := transform.Translate(lockFormat, "userName", u)
	return cache.Delete(key)
}

//SetLoginFail 设置登录失败次数
func (l *CacheMember) SetLoginFail(u string) (int, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(lockFormat, "userName", u)
	v, err := cache.Increment(key, 1)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}
func (l *CacheMember) getLoginFailCnt(u string) (int, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(lockFormat, "userName", u)
	s, err := cache.Get(key)
	if err != nil {
		return 0, err
	}
	if s == "" {
		return 0, nil
	}
	return types.GetInt(s, 0), nil
}

//Save 缓存用户信息
func (l *CacheMember) Save(s *model.MemberState) error {
	s.ReflushCode()
	buff, err := json.Marshal(s)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheFormat, "userName", s.UserName, "ident", s.SysIdent)
	return cache.Set(key, string(buff), l.cacheTime)
}

//Query 用户登录
func (l *CacheMember) Query(u string, ident string) (ls *model.MemberState, err error) {
	//从缓存中查询用户数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheFormat, "userName", u, "ident", ident)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, fmt.Errorf("未缓存用户数据")
	}
	if err = json.Unmarshal([]byte(v), &ls); err != nil {
		return nil, err
	}

	//检查用户登录失败次数是否超过限制，超过时标记用户状态为锁定
	c, err := l.getLoginFailCnt(u)
	if err != nil {
		return nil, err
	}
	if c >= l.maxFailCnt {
		ls.Status = enum.UserLock
	}
	return ls, err
}
