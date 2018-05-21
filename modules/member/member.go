package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/sql"
)

type IMember interface {
	Login(u string, p string, sys int) (*LoginState, string, error)
	Query(uid int64) (db.QueryRow, error)
}

const (
	UserNormal int = iota
	UserLock
	UserDisable
)

//Member 控制用户登录
type Member struct {
	c component.IContainer
}

//NewMember 创建登录对象
func NewMember(c component.IContainer) *Member {
	return &Member{
		c: c,
	}
}

//Query 用户登录
func (l *Member) Query(uid int64) (db.QueryRow, error) {
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

//Login 用户登录
func (l *Member) Login(u string, p string, sys int) (*LoginState, string, error) {
	db := l.c.GetRegularDB()

	//根据用户名密码，查询用户信息
	data, _, _, err := db.Query(sql.QueryUserLoginInfo, map[string]interface{}{
		"user_name": u,
		"password":  p,
	})
	if err != nil {
		return nil, "", context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}
	if data.Len() == 0 {
		return nil, "", context.NewError(context.ERR_UNAUTHORIZED, "用户名密码错误")
	}
	if data.Get(0).GetInt("status", -1) == UserLock {
		return nil, "", context.NewError(context.ERR_LOCKED, "用户被锁定")
	}
	if data.Get(0).GetInt("status", -1) == UserDisable {
		return nil, "", context.NewError(context.ERR_FORBIDDEN, "用户被禁用")
	}
	var member LoginState
	if err := data.Get(0).ToStruct(&u); err != nil {
		return nil, "", context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}

	//查询用户所在系统的登录地址及角色编号
	roles, _, _, err := db.Query(sql.QueryRoles, map[string]interface{}{
		"user_id": member.UserID,
		"sys_id":  sys,
	})
	if roles.IsEmpty() {
		return nil, "", context.NewError(4031, "不允许登录系统")
	}
	member.SystemID = sys
	member.RoleID = roles.Get(0).GetInt("role_id")
	indexURL := roles.Get(0).GetString("index_url")
	return &member, indexURL, nil
}
