package member

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type ICacheMember interface {
	Login(u string, p string, sys int) (*LoginState, string, error)
	Query(uid int64) (db.QueryRow, error)
}

//CacheMember 控制用户登录
type CacheMember struct {
	c           component.IContainer
	cacheFormat string
	lockFormat  string
}

//NewCacheMember 创建登录对象
func NewCacheMember(c component.IContainer) *CacheMember {
	return &CacheMember{
		c:           c,
		cacheFormat: "sso:login:state:%s",
	}
}

func (l *CacheMember) save2Cache(u string, s *LoginState) error {
	buff, err := json.Marshal(s)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	return cache.Set(fmt.Sprintf(l.cacheFormat, u), string(buff), 3600*24)
}

//Query 用户登录
func (l *CacheMember) Query(u string, p string, sys int) (*LoginState, string, bool, error) {
	//从缓存中查询用户数据
	cache := l.c.GetRegularCache()
	key := fmt.Sprintf(l.lockFormat, u)
	v, err := cache.Get(key)
	if err == nil && v != "" {
		return nil, "", false, fmt.Errorf("未缓存用户信息")
	}

}
