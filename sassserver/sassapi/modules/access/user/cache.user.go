package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/transform"
	cacheConst "github.com/micro-plat/sso/sassserver/sassapi/modules/const/cache"
)

type ICacheUser interface {
	// Query(s *model.QueryUserInput) (data db.QueryRows, total int, err error)
	// Save(s *model.QueryUserInput, data db.QueryRows, total int) error
	// Delete() error
	// SaveUser(userID int, data db.QueryRow) error
	// QueryUser(userID int) (data db.QueryRow, err error)
	// DeleteUser() error
	// QueryUserBySys(sysID, pi, ps int) (data db.QueryRows, counr int, err error)
	// SaveUserBySys(sysID, pi, ps int, data db.QueryRows, count int) (err error)

	DeleteLockUserInfo(userName string) error
}

//CacheUser 控制用户登录
type CacheUser struct {
	c         component.IContainer
	cacheTime int
}

//NewCacheUser 创建对象
func NewCacheUser(c component.IContainer) *CacheUser {
	return &CacheUser{
		c:         c,
		cacheTime: 3600 * 24,
	}
}

//DeleteLockUserInfo 解锁用户(删除缓存key)
func (l *CacheUser) DeleteLockUserInfo(userName string) error {
	cache := l.c.GetRegularCache()
	key := transform.Translate(cacheConst.CacheLoginFailCount, "user_name", userName)
	return cache.Delete(key)
}
