package member

// import (
// 	"encoding/json"

// 	"github.com/micro-plat/lib4go/db"
// 	"github.com/micro-plat/sso/mgrapi/modules/model"

// 	"github.com/micro-plat/hydra/component"
// 	"github.com/micro-plat/hydra/context"
// 	"github.com/micro-plat/lib4go/transform"
// )

// type ICacheMember interface {
// 	Save(s *model.MemberState) error
// 	QueryAuth(sysID, userID int64) (err error)
// 	SaveAuth(sysID, userID int64, data db.QueryRows) (err error)
// }

// //CacheMember 控制用户登录
// type CacheMember struct {
// 	c          component.IContainer
// 	maxFailCnt int
// 	cacheTime  int
// }

// const (
// 	cacheFormat     = "{sso}:login:state-info:{@userName}-{@ident}"
// 	cacheCodeFormat = "{sso}:login:state-code:{@userName}-{@ident}"
// 	lockFormat      = "{sso}:login:state-locker:{@userName}"

// 	cacheSysAuth = "{sso}:sys:auth:{@sysID}-{@userID}"
// )

// //NewCacheMember 创建登录对象
// func NewCacheMember(c component.IContainer) *CacheMember {
// 	return &CacheMember{
// 		c:          c,
// 		maxFailCnt: 5,
// 		cacheTime:  3600 * 24,
// 	}
// }

// // 缓存系统权限信息
// func (l *CacheMember) QueryAuth(sysID, userID int64) (err error) {
// 	//从缓存中获取系统数据
// 	cache := l.c.GetRegularCache()
// 	key := transform.Translate(cacheSysAuth, "sysID", sysID, "userID", userID)
// 	v, err := cache.Get(key)
// 	if err != nil {
// 		return err
// 	}
// 	if v == "" {
// 		return context.NewError(context.ERR_FORBIDDEN, "无数据")
// 	}
// 	return nil
// }

// // 查询系统权限信息
// func (l *CacheMember) SaveAuth(sysID, userID int64, data db.QueryRows) (err error) {
// 	buff, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	cache := l.c.GetRegularCache()
// 	key := transform.Translate(cacheSysAuth, "sysID", sysID, "userID", userID)
// 	return cache.Set(key, string(buff), l.cacheTime)
// }

// //Save 缓存用户信息
// func (l *CacheMember) Save(s *model.MemberState) error {
// 	s.ReflushCode()
// 	buff, err := json.Marshal(s)
// 	if err != nil {
// 		return err
// 	}
// 	cache := l.c.GetRegularCache()
// 	key := transform.Translate(cacheFormat, "userName", s.UserName, "ident", s.SysIdent)
// 	return cache.Set(key, string(buff), l.cacheTime)
// }
