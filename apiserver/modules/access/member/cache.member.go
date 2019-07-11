package member

import (
	"encoding/json"

	"github.com/micro-plat/sso/apiserver/modules/model"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"
)

type ICacheMember interface {
	Save(s *model.MemberState) error
	SetLoginSuccess(u string) error
	SetLoginFail(u string) (int, error)
	GetUserInfoByKey(key string) (string, error)
	DeleteInfoByKey(key string)
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
	cacheLoginUser  = "{sso}:login:state-user:{@key}"

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

// GetUserInfoByKey 通过key取缓存的登录用户
func (l *CacheMember) GetUserInfoByKey(key string) (info string, err error) {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cacheLoginUser, "key", key)
	info, err = cache.Get(cachekey)
	return
}

// DeleteInfoByKey key
func (l *CacheMember) DeleteInfoByKey(key string) {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cacheLoginUser, "key", key)
	cache.Delete(cachekey)
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
