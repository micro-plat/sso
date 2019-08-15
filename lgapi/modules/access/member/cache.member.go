package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"

	cachekey "github.com/micro-plat/sso/lgapi/modules/const/cache"
)

type ICacheMember interface {
	CreateUserInfoByCode(code string, userId int64) error
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
