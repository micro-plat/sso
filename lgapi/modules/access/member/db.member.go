package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/lgapi/modules/const/sqls"
	"github.com/micro-plat/sso/lgapi/modules/model"
)

type IDBMember interface {
	Query(u string, p string) (s *model.MemberState, err error)

	QueryByID(uid int64) (db.QueryRow, error)
	QueryByUserName(u string, ident string) (info db.QueryRow, err error)
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

//Query 用户登录时从数据库获取信息
func (l *DBMember) Query(u string, p string) (s *model.MemberState, err error) {
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
	row := data.Get(0)

	s = &model.MemberState{
		UserID:    row.GetInt64("user_id", -1),
		Password:  row.GetString("password"),
		UserName:  row.GetString("user_name"),
		ExtParams: row.GetString("ext_params"),
		Status:    row.GetInt("status"),
	}
	return s, err
}

/////////////////////////////////////////////////

func (l *DBMember) QueryAuth(sysID, userID int64) (data db.QueryRows, err error) {
	db := l.c.GetRegularDB()
	//查询当前系统下是否有此用户
	data, _, _, err = db.Query(sqls.QuerySysAuth, map[string]interface{}{
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
	data, _, _, err := db.Query(sqls.QueryUserInfoByOpenID, map[string]interface{}{
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
	data, _, _, err := db.Query(sqls.QueryUserInfoByUID, map[string]interface{}{
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
	data, _, _, err := db.Query(sqls.QueryUserByLogin, map[string]interface{}{
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
