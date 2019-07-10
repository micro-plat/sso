package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/apiserver/modules/const/sqls"
	"github.com/micro-plat/sso/apiserver/modules/model"
)

type IDBMember interface {
	Query(u string, p string, ident string) (s *model.MemberState, err error)
	QueryByUserName(u string, ident string) (info db.QueryRow, err error)
}

//DBMember 控制用户登录
type DBMember struct {
	c component.IContainer
}

//NewDBMember 创建登录对象
func NewDBMember(c component.IContainer) *DBMember {
	return &DBMember{
		c: c,
	}
}

//Query 用户登录时从数据库获取信息
func (l *DBMember) Query(u string, p string, ident string) (s *model.MemberState, err error) {
	//根据用户名密码，查询用户信息
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QueryUserByLogin, map[string]interface{}{
		"user_name": u,
	})
	if err != nil {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, context.NewError(context.ERR_FORBIDDEN, "用户名或密码错误")
	}

	s = &model.MemberState{}

	//查询用户所在系统的登录地址及角色编号
	roles, _, _, err := db.Query(sqls.QueryUserRole, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})
	if roles.IsEmpty() {
		return nil, context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, "不允许登录系统")
	}
	s.UserID = data.Get(0).GetInt64("user_id", -1)
	s.Status = data.Get(0).GetInt("status")
	s.Password = data.Get(0).GetString("password")
	s.UserName = data.Get(0).GetString("user_name")
	s.RoleID = roles.Get(0).GetInt("role_id")
	s.RoleName = roles.Get(0).GetString("role_name")
	s.IndexURL = roles.Get(0).GetString("index_url")
	s.LoginURL = roles.Get(0).GetString("login_url")
	s.SystemID = roles.Get(0).GetInt("sys_id")
	s.ExtParams = data.Get(0).GetString("ext_params")
	s.SysIdent = ident
	return s, err
}

// QueryByUserName 根据用户名查询用户信息
func (l *DBMember) QueryByUserName(u string, ident string) (info db.QueryRow, err error) {
	//根据用户名，查询用户信息
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
		"user_name": u,
	})
	if err != nil {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, context.NewError(context.ERR_FORBIDDEN, "用户不存在")
	}
	//查询用户所在系统的登录地址及角色编号
	roles, _, _, err := db.Query(sqls.QueryUserRole, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})
	if roles.IsEmpty() {
		return nil, context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, "不允许登录系统")
	}

	userData := data.Get(0)
	userData["ident"] = ident

	return userData, err
}
