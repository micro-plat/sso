package member

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/const/sqls"
	"github.com/micro-plat/sso/sso/errorcode"

	"github.com/micro-plat/sso/loginserver/loginapi/modules/model"
)

type IDBMember interface {
	Query(u, p, ident string) (s *model.MemberState, err error)
	xChangePwd(userID int, expassword string, newpassword string) (err error)
	//QueryByID(uid int64) (types.IXMap, error)
	//CheckUserHasAuth(ident string, userID int64) error
	GetUserInfo(userName string) (types.IXMap, error)
	//UpdateUserStatus(userID int64, status int) error
	//UnLock(userName string) error

	//UpdateUserOpenID(data map[string]string) error

	QueryUserOldPwd(userID int) (types.XMaps, error)
	ChangePwd(userID int, newpassword string) (err error)
	QueryByID(uid int) (types.IXMap, error)
	CheckUserHasAuth(ident string, userID int64) error
	UpdateUserStatus(userID int, status int) error
	UnLock(userName string) error
	UpdateUserLoginTime(userID int64) error
	UpdateUserOpenID(data map[string]string) error
}

//DBMember 控制用户登录
type DBMember struct {
}

//NewDBMember 创建登录对象
func NewDBMember() *DBMember {
	return &DBMember{}
}

//GetUserInfo 根据用户名获取用户信息
func (l *DBMember) GetUserInfo(userName string) (types.IXMap, error) {
	db := components.Def.DB().GetRegularDB()

	userInfo, err := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
		"user_name": userName,
	})

	if err != nil {
		return nil, fmt.Errorf("根据用户名获取用户信息失败(err:%v)", err)
	}
	if userInfo.Len() <= 0 {
		return nil, fmt.Errorf("用户不存在:%s", userName)
	}

	return userInfo.Get(0), nil
}

// ChangePwd 修改密码
func (l *DBMember) xChangePwd(userID int, expassword string, newpassword string) (err error) {
	db := components.Def.DB().GetRegularDB()

	data, err := db.Query(sqls.QueryOldPwd, map[string]interface{}{
		"user_id": userID,
	})

	if err != nil {
		return errs.NewError(http.StatusBadRequest, "用户不存在, 修改失败")
	}

	if strings.ToLower(md5.Encrypt(expassword)) != strings.ToLower(data.Get(0).GetString("password")) {
		return errs.NewError(errorcode.ERR_USER_OLDPWDWRONG, "原密码错误")
	}
	_, err = db.Execute(sqls.SetNewPwd, map[string]interface{}{
		"user_id":  userID,
		"password": md5.Encrypt(newpassword),
	})
	if err != nil {
		errs.NewError(http.StatusInternalServerError, err)
	}
	return nil
}

//UnLock 解锁被锁定的用户
func (l *DBMember) UnLock(userName string) error {
	db := components.Def.DB().GetRegularDB()
	_, err := db.Execute(sqls.UnLockMember, map[string]interface{}{"user_name": userName})
	return err
}

//Query 用户登录时从数据库获取信息 (这个要改)
func (l *DBMember) Query(u, p, ident string) (s *model.MemberState, err error) {

	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
		"user_name": u,
	})
	if err != nil {
		return nil, errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	if data.IsEmpty() {
		return nil, errs.NewError(errorcode.ERR_USER_PWDWRONG, "用户名或密码错误")
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
		SourceID:      row.GetString("source_id"),
		LastLoginTime: row.GetString("last_login_time"),
	}

	if ident == "" {
		return s, nil
	}
	count, err := db.Scalar(sqls.QueryUserHasRoleMenuCount, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})

	if err != nil {
		return nil, errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}

	if types.GetInt(count) <= 0 {
		return nil, errs.NewError(errorcode.ERR_USER_HASNOROLES, "用户没有相关系统权限,请联系管理员")
	}

	roles, err := db.Query(sqls.QueryUserRole, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})

	if err != nil {
		return nil, errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}

	s.SysIdent = ident
	s.SystemID = roles.Get(0).GetInt("sys_id")
	s.RoleName = roles.Get(0).GetString("role_name")
	s.IndexURL = roles.Get(0).GetString("index_url")

	return s, nil
}

//QueryUserOldPwd 查询原密码
func (l *DBMember) QueryUserOldPwd(userID int) (types.XMaps, error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryOldPwd, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return nil, errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	return data, nil
}

// ChangePwd 修改密码
func (l *DBMember) ChangePwd(userID int, newpassword string) (err error) {
	db := components.Def.DB().GetRegularDB()
	_, err = db.Execute(sqls.SetNewPwd, map[string]interface{}{
		"user_id":  userID,
		"password": newpassword,
	})
	if err != nil {
		return errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	return nil
}

//UpdateUserLoginTime 记录用户成功登录时间
func (l *DBMember) UpdateUserLoginTime(userID int64) error {
	db := components.Def.DB().GetRegularDB()
	_, err := db.Execute(sqls.UpdateUserLoginTime, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	return nil
}

// CheckUserHasAuth xx
func (l *DBMember) CheckUserHasAuth(ident string, userID int64) error {
	db := components.Def.DB().GetRegularDB()
	count, err := db.Scalar(sqls.QueryUserRoleCount, map[string]interface{}{
		"user_id": userID,
		"ident":   ident,
	})
	if err != nil {
		return errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	if types.GetInt(count, 0) <= 0 {
		return errs.NewError(errorcode.ERR_USER_HASNOROLES, "没有相应权限，请联系管理员")
	}
	return nil
}

//QueryByID 根据用户编号获取用户信息
func (l *DBMember) QueryByID(uid int) (types.IXMap, error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryUserInfoByUID, map[string]interface{}{
		"user_id": uid,
	})
	if err != nil {
		return nil, errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	if data.IsEmpty() {
		return nil, errs.NewError(errorcode.ERR_USER_NOTEXISTS, "用户不存在:"+strconv.Itoa(uid))
	}
	return data.Get(0), nil
}

//UpdateUserStatus 更新用户状态
func (l *DBMember) UpdateUserStatus(userID int, status int) (err error) {
	db := components.Def.DB().GetRegularDB()
	_, err = db.Execute(sqls.UpdateUserStatus, map[string]interface{}{
		"user_id": userID,
		"status":  status,
	})
	if err != nil {
		err = errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	return nil
}

//UpdateUserOpenID 保存用户的openId信息
func (l *DBMember) UpdateUserOpenID(data map[string]string) error {
	db := components.Def.DB().GetRegularDB()
	exists, err := db.Scalar(sqls.OpenIDIsExists, map[string]interface{}{
		"openid": data["openid"],
	})
	if err != nil {
		return errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	if types.GetInt(exists, 0) == 1 {
		return errs.NewError(errorcode.ERR_OPENID_ONLY_BIND_Once, "一个微信只能绑定一个账户")
	}
	_, err = db.Execute(sqls.AddUserOpenID, map[string]interface{}{
		"user_id": data["userid"],
		"openid":  data["openid"],
	})
	if err != nil {
		err = errs.NewError(errorcode.ERR_SYS_ERROR, err)
	}
	return nil
}

/*
//UnLock 解锁被锁定的用户
func (l *DBMember) UnLock(userName string) error {
	db := components.Def.DB().GetRegularDB()
	_, err := db.Execute(sqls.UnLockMember, map[string]interface{}{"user_name": userName})
	return err
}

// CheckUserHasAuth xx
func (l *DBMember) CheckUserHasAuth(ident string, userID int64) error {
	db := components.Def.DB().GetRegularDB()
	count, err := db.Scalar(sqls.QueryUserRoleCount, map[string]interface{}{
		"user_id": userID,
		"ident":   ident,
	})
	if err != nil {
		return errs.NewError(http.StatusInternalServerError, fmt.Sprintf("出现错误，等会在登录: %s", err))
	}
	if types.GetInt(count, 0) <= 0 {
		return errs.NewError(errorcode.ERR_USER_HASNOROLES, "没有相应权限，请联系管理员")
	}
	return nil
}


//QueryByID 根据用户编号获取用户信息
func (l *DBMember) QueryByID(uid int64) (types.IXMap, error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryUserInfoByUID, map[string]interface{}{
		"user_id": uid,
	})
	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, errs.NewError(errorcode.ERR_USER_NOTEXISTS, "用户不存在")
	}
	return data.Get(0), nil
}


//UpdateUserStatus 更新用户状态
func (l *DBMember) UpdateUserStatus(userID int64, status int) error {
	db := components.Def.DB().GetRegularDB()
	_, err := db.Execute(sqls.UpdateUserStatus, map[string]interface{}{
		"user_id": userID,
		"status":  status,
	})
	return err
}

//UpdateUserOpenID 保存用户的openId信息
func (l *DBMember) UpdateUserOpenID(data map[string]string) error {
	db := components.Def.DB().GetRegularDB()
	exists, err1 := db.Scalar(sqls.OpenIDIsExists, map[string]interface{}{
		"openid": data["openid"],
	})
	if err1 != nil {
		return err1
	}
	if types.GetInt(exists, 0) == 1 {
		return errs.NewError(errorcode.ERR_OPENID_ONLY_BIND_Once, "一个微信只能绑定一个账户")
	}
	_, err := db.Execute(sqls.AddUserOpenID, map[string]interface{}{
		"user_id": data["userid"],
		"openid":  data["openid"],
	})
	return err
}
*/
