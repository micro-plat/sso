package member

import (
	"time"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/types"
	cachekey "github.com/micro-plat/sso/loginserver/apiserver/modules/const/cache"
)

type ICacheMember interface {
	GetUserInfoByCode(code string) (string, error)
	DeleteInfoByCode(code string)

	SetLoginFail(userName string) (int, error)
	SetLoginSuccess(userName string) error
	SetUnLockTime(userName string, expire int) error
	ExistsUnLockTime(userName string) bool
	GetLoginFailCnt(userName string) (int, error)

	SaveLoginVerifyCode(userName, verfiyCode string, expire int) error
	GetLoginVerifyCode(userName string) (string, error)
}

//CacheMember 控制用户登录
type CacheMember struct {
}

//NewCacheMember 创建登录对象
func NewCacheMember() *CacheMember {
	return &CacheMember{}
}

// GetUserInfoByCode 通过key取缓存的登录用户
func (l *CacheMember) GetUserInfoByCode(code string) (info string, err error) {
	cache := components.Def.Cache().GetRegularCache()
	cachekey := types.Translate(cachekey.CacheLoginUser, "code", code)
	info, err = cache.Get(cachekey)
	return
}

// DeleteInfoByCode code
func (l *CacheMember) DeleteInfoByCode(code string) {
	cache := components.Def.Cache().GetRegularCache()
	cachekey := types.Translate(cachekey.CacheLoginUser, "code", code)
	cache.Delete(cachekey)
}

//GetLoginFailCnt 获取登录失败次数
func (l *CacheMember) GetLoginFailCnt(userName string) (int, error) {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cachekey.CacheLoginFailCount, "user_name", userName)
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
func (l *CacheMember) ExistsUnLockTime(userName string) bool {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cachekey.CacheLoginFailUnLockTime, "user_name", userName)
	return cache.Exists(key)
}

//SetLoginFail 设置登录失败次数
func (l *CacheMember) SetLoginFail(userName string) (int, error) {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cachekey.CacheLoginFailCount, "user_name", userName)
	v, err := cache.Increment(key, 1)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

//SetLoginSuccess 设置为登录成功
func (l *CacheMember) SetLoginSuccess(userName string) error {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cachekey.CacheLoginFailCount, "user_name", userName)
	return cache.Delete(key)
}

//SetUnLockTime 设置解锁过期时间
func (l *CacheMember) SetUnLockTime(userName string, expire int) error {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cachekey.CacheLoginFailUnLockTime, "user_name", userName)
	return cache.Set(key, time.Now().Add(time.Second*time.Duration(expire)).Format("2006-01-02 15:04:05"), expire)
}

//SaveLoginVerifyCode 保存登录验证码
func (l *CacheMember) SaveLoginVerifyCode(userName, verfiyCode string, expire int) error {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cachekey.CacheLoginVerifyCode, "user_name", userName)
	return cache.Set(key, verfiyCode, expire)
}

//GetLoginVerifyCode 获取登录验证码
func (l *CacheMember) GetLoginVerifyCode(userName string) (string, error) {
	cache := components.Def.Cache().GetRegularCache()
	key := types.Translate(cachekey.CacheLoginVerifyCode, "user_name", userName)
	code, err := cache.Get(key)
	if err != nil {
		return "", err
	}
	return code, nil
}
