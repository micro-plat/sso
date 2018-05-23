package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/sql"
)

type IDBMember interface {
	Login(u string, p string, sys int) (*LoginState, string, error)
	Query(uid int64) (db.QueryRow, error)
}

const (
	UserNormal int = iota
	UserLock
	UserDisable
)

//DBMember 控制用户登录
type DBMember struct {
	c           component.IContainer
	cacheFormat string
}

//NewDBMember 创建登录对象
func NewDBMember(c component.IContainer) *DBMember {
	return &DBMember{
		c:           c,
		cacheFormat: "sso:login:state:%s",
	}
}

//QueryByID 根据用户编号获取用户信息
func (l *DBMember) QueryByID(uid int64) (db.QueryRow, error) {
	db := l.c.GetRegularDB()

	//根据用户名密码，查询用户信息
	data, _, _, err := db.Query(sql.QueryUserInfoByUID, map[string]interface{}{
		"user_id": uid,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}

//QueryByLogin 用户登录时从数据库获取信息
func (l *DBMember) QueryByLogin(u string, p string, sysID int) (s *LoginState, err error) {
	//根据用户名密码，查询用户信息
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sql.QueryUserLoginInfo, map[string]interface{}{
		"user_name": u,
		"password":  p,
		"sys_id":    sysID,
	})
	if err != nil {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, context.NewError(403, "不允许登录系统")
	}
	err = data.Get(0).ToStruct(&s)
	return s, err
}
