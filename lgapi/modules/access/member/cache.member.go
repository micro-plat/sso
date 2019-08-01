package member

import (
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"

	cachekey "github.com/micro-plat/sso/lgapi/modules/const/cache"
)

type ICacheMember interface {
	SetLoginFail(u string) (int, error)
	CreateUserInfoByCode(code string, userId int64) error
	SaveWxStateCode(code, content string) error
	GetContentByStateCode(code string) (string, error)

	ExistsWxStateCode(code string) (bool, error)
	SaveWxLoginInfo(state, content string) error
	GetWxLoginInfoByStateCode(stateCode string) (string, error)

	CreateValiCode(userName, code string) error
	VerifyValidCode(userName, validCode string) (isValid bool, err error)
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

// CreateUserInfoByCode 通过key取缓存的登录用户
func (l *CacheMember) CreateUserInfoByCode(code string, userId int64) error {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cachekey.CacheLoginUser, "key", code)
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

// SaveWxStateCode xx
func (l *CacheMember) SaveWxStateCode(code, content string) error {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cachekey.WxStateCode, "code", code)
	contentT := content
	if contentT == "" {
		contentT = "1"
	}
	return cache.Set(cachekey, contentT, 60*5)
}

// GetContentByStateCode xx
func (l *CacheMember) GetContentByStateCode(code string) (string, error) {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cachekey.WxStateCode, "code", code)
	return cache.Get(cachekey)
}

// ExistsWxStateCode xx
func (l *CacheMember) ExistsWxStateCode(code string) (bool, error) {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cachekey.WxStateCode, "code", code)
	return cache.Exists(cachekey), nil
}

// SaveWxLoginInfo xx
func (l *CacheMember) SaveWxLoginInfo(state, content string) error {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cachekey.WxStateCode, "code", state)
	return cache.Set(cachekey, content, 60*5)
}

//GetWxLoginInfoByStateCode xx
func (l *CacheMember) GetWxLoginInfoByStateCode(stateCode string) (string, error) {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cachekey.WxStateCode, "code", stateCode)
	return cache.Get(cachekey)
}

// CreateValiCode xx
func (l *CacheMember) CreateValiCode(userName, code string) error {
	cache := l.c.GetRegularCache()

	key := transform.Translate(cachekey.WechatValidcodeCacheKey, "senduser", userName)
	err := cache.Set(key, code, 60*5)
	if err != nil {
		return context.NewError(context.ERR_SERVER_ERROR, err)
	}
	return nil
}

//VerifyValidCode xx
func (l *CacheMember) VerifyValidCode(userName, validCode string) (isValid bool, err error) {
	cache := l.c.GetRegularCache()

	key := transform.Translate(cachekey.WechatValidcodeCacheKey, "senduser", userName)
	val, err := cache.Get(key)
	if err != nil {
		return false, context.NewError(context.ERR_SERVER_ERROR, err)
	}
	if val == "" {
		return false, nil
	}

	if !strings.EqualFold(val, validCode) {
		cacheCountKey := transform.Translate(cachekey.WechatValidcodeErrorCountCacheKey, "senduser", userName)

		var newval int64
		newval = 1

		if flag := cache.Exists(cacheCountKey); !flag {
			cache.Set(cacheCountKey, "1", 60*5)
		} else {
			newval, err = cache.Increment(cacheCountKey, 1)
			if err != nil {
				return false, context.NewError(context.ERR_SERVER_ERROR, err)
			}
		}
		if newval >= 3 {
			cache.Delete(key)
			cache.Delete(cacheCountKey)
		}
		return false, nil
	}
	cache.Delete(key)
	return true, nil
}
