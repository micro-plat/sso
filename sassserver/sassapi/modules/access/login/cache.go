package login

import (
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"
	cachekey "github.com/micro-plat/sso/sassserver/sassapi/modules/const/cache"
)

type ICacheLogin interface {
	SetLoginFail(mobile string) (int, error)
	SetLoginSuccess(mobile string) error
	SetUnLockTime(mobile string, expire int) error
	ExistsUnLockTime(mobile string) bool
	GetLoginFailCnt(mobile string) (int, error)

	SaveLoginVerifyCode(mobile, verfiyCode string, expire int) error
	GetLoginVerifyCode(mobile string) (string, error)
}

//CacheLogin 控制用户登录
type CacheLogin struct {
	c          component.IContainer
	maxFailCnt int
	cacheTime  int
}

//NewCacheLogin 创建登录对象
func NewCacheLogin(c component.IContainer) *CacheLogin {
	return &CacheLogin{
		c: c,
	}
}

//GetLoginFailCnt 获取登录失败次数
func (l *CacheLogin) GetLoginFailCnt(mobile string) (int, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailCount, "mobile", mobile)
	s, err := cache.Get(key)
	if err != nil {
		return 0, err
	}
	if s == "" {
		return 0, nil
	}
	return types.GetInt(s, 0), nil
}

//ExistsUnLockTime 解锁时间是否过期
func (l *CacheLogin) ExistsUnLockTime(mobile string) bool {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailUnLockTime, "mobile", mobile)
	return cache.Exists(key)
}

//SetLoginFail 设置登录失败次数
func (l *CacheLogin) SetLoginFail(mobile string) (int, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailCount, "mobile", mobile)
	v, err := cache.Increment(key, 1)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

//SetLoginSuccess 设置为登录成功
func (l *CacheLogin) SetLoginSuccess(mobile string) error {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailCount, "mobile", mobile)
	return cache.Delete(key)
}

//SetUnLockTime 设置解锁过期时间
func (l *CacheLogin) SetUnLockTime(mobile string, expire int) error {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginFailUnLockTime, "mobile", mobile)
	return cache.Set(key, time.Now().Add(time.Second*time.Duration(expire)).Format("2006-01-02 15:04:05"), expire)
}

//SaveLoginVerifyCode 保存登录验证码
func (l *CacheLogin) SaveLoginVerifyCode(mobile, verfiyCode string, expire int) error {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginVerifyCode, "mobile", mobile)
	return cache.Set(key, verfiyCode, expire)
}

//GetLoginVerifyCode 获取登录验证码
func (l *CacheLogin) GetLoginVerifyCode(mobile string) (string, error) {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cachekey.CacheLoginVerifyCode, "mobile", mobile)
	code, err := cache.Get(key)
	if err != nil {
		return "", err
	}
	return code, nil
}
