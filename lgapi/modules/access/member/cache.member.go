package member

import (
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"

	cachekey "github.com/micro-plat/sso/lgapi/modules/const/cache"
)

type ICacheMember interface {
	CreateUserInfoByCode(code string, userId int64) error
	SetLoginFail(userName string) (int, error)
	GetLoginFailCnt(userName string) (int, error)
	SetLoginSuccess(userName string) error
	SetUnLockTime(userName string, expire int) error
	ExistsUnLockTime(userName string) bool
}

//CacheMember 控制用户登录
type CacheMember struct {
	c          component.IContainer
	maxFailCnt int
	cacheTime  int
}

//NewCacheMember 创建登录对象
func NewCacheMember(c component.IContainer) *CacheMember {
	return &CacheMember{
		c: c,
	}
}

// CreateUserInfoByCode 通过key取缓存的登录用户
func (l *CacheMember) CreateUserInfoByCode(code string, userId int64) error {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cachekey.CacheLoginUser, "key", code)
	return cache.Set(cachekey, types.GetString(userId), 300)
}

//SetLoginFail 设置登录失败次数
func (l *CacheMember) SetLoginFail(userName string) (int, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailCount, "user_name", userName)
	v, err := cache.Increment(key, 1)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

//GetLoginFailCnt 获取登录失败次数
func (l *CacheMember) GetLoginFailCnt(userName string) (int, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailCount, "user_name", userName)
	s, err := cache.Get(key)
	if err != nil {
		return 0, err
	}
	if s == "" {
		return 0, nil
	}
	return types.GetInt(s, 0), nil
}

//SetLoginSuccess 设置为登录成功
func (l *CacheMember) SetLoginSuccess(u string) error {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailCount, "user_name", u)
	return cache.Delete(key)
}

//SetUnLockTime 设置解锁过期时间
func (l *CacheMember) SetUnLockTime(userName string, expire int) error {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailUnLockTime, "user_name", userName)
	return cache.Set(key, time.Now().Add(time.Second*time.Duration(expire)).Format("2006-01-02 15:04:05"), expire)
}

//ExistsUnLockTime 解锁时间是否过期
func (l *CacheMember) ExistsUnLockTime(userName string) bool {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailUnLockTime, "user_name", userName)
	return cache.Exists(key)
}
