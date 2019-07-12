package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"

	cachekey "github.com/micro-plat/sso/lgapi/modules/const/cache"
)

type ICacheMember interface {
	SetLoginFail(u string) (int, error)
	SetUserInfoByKey(key string, userId int64) error
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
		c:          c,
		maxFailCnt: 5,
		cacheTime:  3600 * 24,
	}
}

// SetUserInfoByKey 通过key取缓存的登录用户
func (l *CacheMember) SetUserInfoByKey(key string, userId int64) error {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cachekey.CacheLoginUser, "key", key)
	return cache.Set(cachekey, types.GetString(userId), 60)
}

//SetLoginFail 设置登录失败次数
func (l *CacheMember) SetLoginFail(u string) (int, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.LockFormat, "userName", u)
	v, err := cache.Increment(key, 1)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}
func (l *CacheMember) getLoginFailCnt(u string) (int, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.LockFormat, "userName", u)
	s, err := cache.Get(key)
	if err != nil {
		return 0, err
	}
	if s == "" {
		return 0, nil
	}
	return types.GetInt(s, 0), nil
}
