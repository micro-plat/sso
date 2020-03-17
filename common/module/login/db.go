package login

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/common/module/const/sqls"
	"github.com/micro-plat/sso/common/module/model"
)

type IDBMember interface {
	Query(u, p, ident string) (s *model.MemberState, err error)
	QueryUserOldPwd(userID int) (db.QueryRows, error)
	ChangePwd(userID int, newpassword string) (err error)
	QueryByID(uid int) (db.QueryRow, error)
	CheckUserHasAuth(ident string, userID int64) error
	GetUserInfo(u string) (db.QueryRows, error)
	UpdateUserStatus(userID int, status int) error
	UnLock(userName string) error
	UpdateUserLoginTime(userID int64) error
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

//Query 用户登录时从数据库获取信息 (这个要改)
func (l *DBMember) Query(u, p, ident string) (s *model.MemberState, err error) {
	db := l.c.GetRegularDB()
	data, _, _, errt := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
		"user_name": u,
	})
	if errt != nil {
		return nil, errt
	}
	if data.IsEmpty() {
		return nil, context.NewError(model.ERR_USER_PWDWRONG, "用户名或密码错误")
	}

	row := data.Get(0)
	s = &model.MemberState{
		UserID:        row.GetInt64("user_id", -1),
		Password:      row.GetString("password"),
		UserName:      row.GetString("user_name"),
		FullName:      row.GetString("full_name"),
		ExtParams:     row.GetString("ext_params"),
		Status:        row.GetInt("status"),
		Source:        row.GetString("source"),
		SourceID:      row.GetInt("source_id"),
		LastLoginTime: row.GetString("last_login_time"),
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

//QueryUserOldPwd 查询原密码
func (l *DBMember) QueryUserOldPwd(userID int) (db.QueryRows, error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QueryOldPwd, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ChangePwd 修改密码
func (l *DBMember) ChangePwd(userID int, newpassword string) (err error) {
	db := l.c.GetRegularDB()
	_, _, _, err = db.Execute(sqls.SetNewPwd, map[string]interface{}{
		"user_id":  userID,
		"password": md5.Encrypt(newpassword),
	})
	if err != nil {
		return err
	}
	return nil
}

//UpdateUserLoginTime　记录用户成功登录时间
func (l *DBMember) UpdateUserLoginTime(userID int64) error {
	db := l.c.GetRegularDB()
	_, q, a, err := db.Execute(sqls.UpdateUserLoginTime, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return fmt.Errorf("UpdateUserLoginTime 出错: sql:%s, arg:%+v, err:%+v", q, a, err)
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
		return err
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
func (l *DBMember) QueryByID(uid int) (db.QueryRow, error) {
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
func (l *DBMember) UpdateUserStatus(userID int, status int) error {
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
