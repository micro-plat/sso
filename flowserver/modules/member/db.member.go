package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/flowserver/modules/const/sql"
)

type IDBMember interface {
	QueryByID(uid int64) (db.QueryRow, error)
	Query(u string, p string, ident string) (s *MemberState, err error)
	GetUserInfo(u string) (db.QueryRow, error)
	QueryByOpenID(string) (db.QueryRow, error)
	QueryAuth(sysID, userID int64) (data db.QueryRows, err error)
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
func (l *DBMember) QueryAuth(sysID, userID int64) (data db.QueryRows, err error) {
	db := l.c.GetRegularDB()
	//查询当前系统下是否有此用户
	data, _, _, err = db.Query(sql.QuerySysAuth, map[string]interface{}{
		"sys_id":  sysID,
		"user_id": userID,
	})
	if err != nil || data.IsEmpty() {
		return nil, fmt.Errorf("没有权限：err:%v,data:%v", err, data)
	}
	return data, nil
}

//QueryByOpenID 根据openid 查询用户信息
func (l *DBMember) QueryByOpenID(open_id string) (db.QueryRow, error) {
	db := l.c.GetRegularDB()

	//根据用户名密码，查询用户信息
	data, _, _, err := db.Query(sql.QueryUserInfoByOpenID, map[string]interface{}{
		"open_id": open_id,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil

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

//GetUserInfo 根据用户名获取用户信息
func (l *DBMember) GetUserInfo(u string) (db.QueryRow, error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sql.QueryUserByLogin, map[string]interface{}{
		"user_name": u,
	})
	if err != nil {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, err)
	}
	if data.IsEmpty() {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, fmt.Sprintf("用户(%s)不存在", u))
	}
	return data.Get(0), nil
}

//Query 用户登录时从数据库获取信息
func (l *DBMember) Query(u string, p string, ident string) (s *MemberState, err error) {
	//根据用户名密码，查询用户信息
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sql.QueryUserByLogin, map[string]interface{}{
		"user_name": u,
	})
	if err != nil {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, context.NewError(context.ERR_FORBIDDEN, "用户名或密码错误")
	}
	s = &MemberState{}
	// if err = data.Get(0).ToStruct(s); err != nil {
	// 	return nil, err
	// }
	//查询用户所在系统的登录地址及角色编号
	roles, _, _, err := db.Query(sql.QueryUserRole, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})
	if roles.IsEmpty() {
		return nil, context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, "不允许登录系统")
	}
	s.UserID = data.Get(0).GetInt64("user_id", -1)
	s.Password = data.Get(0).GetString("password")
	s.UserName = data.Get(0).GetString("user_name")
	s.RoleID = roles.Get(0).GetInt("role_id")
	s.RoleName = roles.Get(0).GetString("role_name")
	s.IndexURL = roles.Get(0).GetString("index_url")
	s.LoginURL = roles.Get(0).GetString("login_url")
	s.SystemID = roles.Get(0).GetInt("sys_id")
	s.SysIdent = ident
	return s, err
}
