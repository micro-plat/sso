package member

import (
	"strings"
	"time"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/const/cachekey"
 	"github.com/micro-plat/sso/loginserver/loginapi/modules/const/errorcode"

)

type ICacheMember interface {
	CreateUserInfoByCode(code string, userId int64) error
	SetLoginFail(userName string) (int, error)
	GetLoginFailCnt(userName string) (int, error)
	SetLoginSuccess(userName string) error
	SetUnLockTime(userName string, expire int) error
	ExistsUnLockTime(userName string) bool

	SaveWxStateCode(stateCode, userid string) error
	GetWxStateCodeUserId(stateCode string) (string, error)
	SetLoginValidateCode(validateCode, userName string) error
	CheckLoginValidateCode(userName, wxCode string) error
}

//CacheMember 控制用户登录
type CacheMember struct {
	maxFailCnt int
	cacheTime  int
}

//NewCacheMember 创建登录对象
func NewCacheMember() *CacheMember {
	return &CacheMember{}
}

// CreateUserInfoByCode 通过key取缓存的登录用户
func (l *CacheMember) CreateUserInfoByCode(code string, userId int64) error {
	cache := components.Def.Cache().GetRegularCache("redis")
	cachekey := types.Translate(cachekey.CacheLoginUser, "key", code)
	return cache.Set(cachekey, types.GetString(userId), 300)
}

//SetLoginFail 设置登录失败次数
func (l *CacheMember) SetLoginFail(userName string) (int, error) {
	cache := components.Def.Cache().GetRegularCache("redis")
	key := types.Translate(cachekey.CacheLoginFailCount, "user_name", userName)
	v, err := cache.Increment(key, 1)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

//GetLoginFailCnt 获取登录失败次数
func (l *CacheMember) GetLoginFailCnt(userName string) (int, error) {
	cache := components.Def.Cache().GetRegularCache("redis")
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

//SetLoginSuccess 设置为登录成功
func (l *CacheMember) SetLoginSuccess(u string) error {
	cache := components.Def.Cache().GetRegularCache("redis")
	key := types.Translate(cachekey.CacheLoginFailCount, "user_name", u)
	return cache.Delete(key)
}

//SetUnLockTime 设置解锁过期时间
func (l *CacheMember) SetUnLockTime(userName string, expire int) error {
	cache := components.Def.Cache().GetRegularCache("redis")
	key := types.Translate(cachekey.CacheLoginFailUnLockTime, "user_name", userName)
	return cache.Set(key, time.Now().Add(time.Second*time.Duration(expire)).Format("2006-01-02 15:04:05"), expire)
}

//ExistsUnLockTime 解锁时间是否过期
func (l *CacheMember) ExistsUnLockTime(userName string) bool {
	cache := components.Def.Cache().GetRegularCache("redis")
	key := types.Translate(cachekey.CacheLoginFailUnLockTime, "user_name", userName)
	return cache.Exists(key)
}

//SaveWxStateCode 保存微信凭证
func (l *CacheMember) SaveWxStateCode(stateCode, userid string) error {
	cache := components.Def.Cache().GetRegularCache("redis")
	cachekey := types.Translate(cachekey.CacheWxStateCode, "code", stateCode)
	return cache.Set(cachekey, userid, 60*5)
}

//GetWxStateCodeUserId 获取wxstatecode中存的user_id
func (l *CacheMember) GetWxStateCodeUserId(stateCode string) (string, error) {
	cache := components.Def.Cache().GetRegularCache("redis")
	cachekey := types.Translate(cachekey.CacheWxStateCode, "code", stateCode)
	value, err := cache.Get(cachekey)
	if err != nil {
		return "", err
	}
	return value, nil
}

//SetLoginValidateCode 保存用户登录验证码
func (l *CacheMember) SetLoginValidateCode(validateCode, userName string) error {
	cache := components.Def.Cache().GetRegularCache("redis")
	cachekey := types.Translate(cachekey.CacheLoginValidateCode, "user_name", userName)
	return cache.Set(cachekey, validateCode, 60*5)
}

//CheckLoginValidateCode 验证用户登录验证码
func (l *CacheMember) CheckLoginValidateCode(userName, wxCode string) error {
	cache := components.Def.Cache().GetRegularCache("redis")
	validateCodeKey := types.Translate(cachekey.CacheLoginValidateCode, "user_name", userName)
	value, err := cache.Get(validateCodeKey)
	if err != nil {
		return err
	}
	if strings.EqualFold(value, "") {
		return errs.NewError(errorcode.ERR_VALIDATECODE_TIMEOUT, "验证码过期,重新发送验证码")
	}

	cacheCountKey := types.Translate(cachekey.CacheLoginValidateCodeFaildCount, "user_name", userName)
	if !strings.EqualFold(value, wxCode) {
		var newval int64 = 1
		if flag := cache.Exists(cacheCountKey); !flag {
			cache.Set(cacheCountKey, types.GetString(newval), 60*5)
		} else {
			newval, err = cache.Increment(cacheCountKey, 1)
			if err != nil {
				return err
			}
		}
		if newval >= 5 {
			cache.Delete(validateCodeKey)
			cache.Delete(cacheCountKey)
		}
		return errs.NewError(errorcode.ERR_VALIDATECODE_WRONG, "验证码错误")
	}
	cache.Delete(validateCodeKey)
	cache.Delete(cacheCountKey)
	return nil
}
