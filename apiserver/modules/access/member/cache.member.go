package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/transform"
)

type ICacheMember interface {
	//Save(s *model.MemberState) error
	GetUserInfoByCode(code string) (string, error)
	DeleteInfoByCode(code string)
}

//CacheMember 控制用户登录
type CacheMember struct {
	c component.IContainer
}

const (
	//cacheFormat    = "{sso}:login:state-info:{@userName}-{@ident}"
	cacheLoginUser = "{sso}:login:state-user:{@code}"
)

//NewCacheMember 创建登录对象
func NewCacheMember(c component.IContainer) *CacheMember {
	return &CacheMember{
		c: c,
	}
}

// GetUserInfoByCode 通过key取缓存的登录用户
func (l *CacheMember) GetUserInfoByCode(code string) (info string, err error) {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cacheLoginUser, "code", code)
	info, err = cache.Get(cachekey)
	return
}

// DeleteInfoByCode code
func (l *CacheMember) DeleteInfoByCode(code string) {
	cache := l.c.GetRegularCache()
	cachekey := transform.Translate(cacheLoginUser, "code", code)
	cache.Delete(cachekey)
}
