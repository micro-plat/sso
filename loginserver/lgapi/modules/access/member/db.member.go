package member

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/const/sqls"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/model"
)

type IDBMember interface {
	Query(u, p, ident string) (s *model.MemberState, err error)
	ChangePwd(userID int, expassword string, newpassword string) (err error)
	QueryByID(uid int64) (db.QueryRow, error)
	CheckUserHasAuth(ident string, userID int64) error
	GetUserInfo(u string) (db.QueryRows, error)
	UpdateUserStatus(userID int64, status int) error
	UnLock(userName string) error

	UpdateUserOpenID(data map[string]string) error
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

//UnLock 解锁被锁定的用户
func (l *DBMember) UnLock(userName string) error {
	db := l.c.GetRegularDB()
	_, _, _, err := db.Execute(sqls.UnLockMember, map[string]interface{}{"user_name": userName})
	return err
}

//Query 用户登录时从数据库获取信息
func (l *DBMember) Query(u, p, ident string) (s *model.MemberState, err error) {
	db := l.c.GetRegularDB()
	data, _, _, errt := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
		"user_name": u,
	})
	if errt != nil {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, context.NewError(model.ERR_USER_PWDWRONG, "用户名或密码错误")
	}

	row := data.Get(0)
	s = &model.MemberState{
		UserID:    row.GetInt64("user_id", -1),
		Password:  row.GetString("password"),
		UserName:  row.GetString("user_name"),
		ExtParams: row.GetString("ext_params"),
		Status:    row.GetInt("status"),
	}

	roles, _, _, erro := db.Query(sqls.QueryUserRole, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})

	if erro != nil || roles.IsEmpty() {
		return nil, context.NewError(model.ERR_USER_HASNOROLES, "用户没有相关系统权限,请联系管理员")
	}

	s.SysIdent = ident
	s.SystemID = roles.Get(0).GetInt("sys_id")
	s.RoleName = roles.Get(0).GetString("role_name")
	s.IndexURL = roles.Get(0).GetString("index_url")

	return s, err
}

// ChangePwd 修改密码
func (l *DBMember) ChangePwd(userID int, expassword string, newpassword string) (err error) {
	db := l.c.GetRegularDB()

	data, _, _, err := db.Query(sqls.QueryOldPwd, map[string]interface{}{
		"user_id": userID,
	})

	if err != nil {
		return context.NewError(context.ERR_BAD_REQUEST, "用户不存在, 修改失败")
	}

	if strings.ToLower(md5.Encrypt(expassword)) != strings.ToLower(data.Get(0).GetString("password")) {
		return context.NewError(model.ERR_USER_OLDPWDWRONG, "原密码错误")
	}
	_, _, _, err = db.Execute(sqls.SetNewPwd, map[string]interface{}{
		"user_id":  userID,
		"password": md5.Encrypt(newpassword),
	})
	if err != nil {
		context.NewError(context.ERR_SERVER_ERROR, err)
	}
	return nil
}

// CheckUserHasAuth xx
func (l *DBMember) CheckUserHasAuth(ident string, userID int64) error {
	db := l.c.GetRegularDB()
	count, _, _, err := db.Scalar(sqls.QueryUserRoleCount, map[string]interface{}{
		"user_id": userID,
		"ident":   ident,
	})
	if err != nil {
		return context.NewError(context.ERR_SERVER_ERROR, fmt.Sprintf("出现错误，等会在登录: %s", err))
	}
	if types.GetInt(count, 0) <= 0 {
		return context.NewError(model.ERR_USER_HASNOROLES, "没有相应权限，请联系管理员")
	}
	return nil
}

//GetUserInfo 根据用户名获取用户信息
func (l *DBMember) GetUserInfo(u string) (db.QueryRows, error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
		"user_name": u,
	})
	return data, err
}

//QueryByID 根据用户编号获取用户信息
func (l *DBMember) QueryByID(uid int64) (db.QueryRow, error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QueryUserInfoByUID, map[string]interface{}{
		"user_id": uid,
	})
	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, context.NewError(model.ERR_USER_NOTEXISTS, "用户不存在")
	}
	return data.Get(0), nil
}

//UpdateUserStatus 更新用户状态
func (l *DBMember) UpdateUserStatus(userID int64, status int) error {
	db := l.c.GetRegularDB()
	_, _, _, err := db.Execute(sqls.UpdateUserStatus, map[string]interface{}{
		"user_id": userID,
		"status":  status,
	})
	return err
}

//UpdateUserOpenID 保存用户的openId信息
func (l *DBMember) UpdateUserOpenID(data map[string]string) error {
	db := l.c.GetRegularDB()
	exists, _, _, err1 := db.Scalar(sqls.OpenIDIsExists, map[string]interface{}{
		"openid": data["openid"],
	})
	if err1 != nil {
		return err1
	}
	if types.GetInt(exists, 0) == 1 {
		return context.NewError(model.ERR_OPENID_ONLY_BIND_Once, "一个微信只能绑定一个账户")
	}
	_, _, _, err := db.Execute(sqls.AddUserOpenID, map[string]interface{}{
		"user_id": data["userid"],
		"openid":  data["openid"],
	})
	return err
}
