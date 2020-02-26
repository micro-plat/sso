package login

import (
	"errors"
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/const/sqls"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/config"
)

//IDbLogin interface
type IDbLogin interface {
	UnLock(mobile string) error
	Query(mobile, p, ident string) (s *model.MemberState, err error)
	UpdateUserStatus(userID int64, status int) error
	QueryUserMenu(uid int64, ident string) ([]map[string]interface{}, error)
	GetSystemInfo(ident string) (s db.QueryRow, err error)
	ChangePwd(userID int64, expassword, newpassword string) error
}

//DbLogin 登录相关功能
type DbLogin struct {
	c component.IContainer
}

//NewDbLogin new
func NewDbLogin(c component.IContainer) *DbLogin {
	return &DbLogin{
		c: c,
	}
}

//UnLock 解锁被锁定的用户
func (l *DbLogin) UnLock(userName string) error {
	db := l.c.GetRegularDB(config.DbName)
	_, _, _, err := db.Execute(sqls.UnLockMember, map[string]interface{}{"user_name": userName})
	return err
}

//Query 用户登录时从数据库获取信息
func (l *DbLogin) Query(mobile, p, ident string) (s *model.MemberState, err error) {
	db := l.c.GetRegularDB(config.DbName)
	data, _, _, errt := db.Query(sqls.QueryUserByMobile, map[string]interface{}{
		"mobile": mobile,
	})
	if errt != nil {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, context.NewError(model.ERR_USER_PWDWRONG, "手机号或密码错误")
	}

	row := data.Get(0)
	s = &model.MemberState{
		BelongID:   row.GetInt("belong_id"),
		BelongType: row.GetInt("belong_type"),
		UserID:     row.GetInt64("user_id", -1),
		Password:   row.GetString("password"),
		UserName:   row.GetString("user_name"),
		ExtParams:  row.GetString("ext_params"),
		Status:     row.GetInt("status"),
	}
	fmt.Print(row)

	roles, _, _, err := db.Query(sqls.QueryUserRole, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})
	if err != nil {
		return nil, err
	}

	if roles.IsEmpty() {
		return nil, context.NewError(model.ERR_USER_HASNOROLES, "用户没有相关系统权限,请联系管理员")
	}

	s.SysIdent = ident
	s.SystemID = roles.Get(0).GetInt("sys_id")
	s.RoleName = roles.Get(0).GetString("role_name")
	s.IndexURL = roles.Get(0).GetString("index_url")

	return s, err
}

//UpdateUserStatus 更新用户状态
func (l *DbLogin) UpdateUserStatus(userID int64, status int) error {
	db := l.c.GetRegularDB(config.DbName)
	_, _, _, err := db.Execute(sqls.UpdateUserStatus, map[string]interface{}{
		"user_id": userID,
		"status":  status,
	})
	return err
}

//QueryUserMenu 获取用户指定系统的菜单信息
func (l *DbLogin) QueryUserMenu(uid int64, ident string) ([]map[string]interface{}, error) {
	db := l.c.GetRegularDB(config.DbName)
	data, _, _, err := db.Query(sqls.QueryUserMenus, map[string]interface{}{
		"user_id": uid,
		"ident":   ident,
	})
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0, 4)
	for _, row1 := range data {
		if row1.GetInt("parent") == 0 && row1.GetInt("level_id") == 1 {
			children1 := make([]map[string]interface{}, 0, 4)
			for _, row2 := range data {
				if row2.GetInt("parent") == row1.GetInt("id") && row2.GetInt("level_id") == 2 {
					children2 := make([]map[string]interface{}, 0, 8)
					for _, row3 := range data {
						if row3.GetInt("parent") == row2.GetInt("id") && row3.GetInt("level_id") == 3 {
							children2 = append(children2, row3)
						}
					}
					children1 = append(children1, row2)
					row2["children"] = children2
				}
			}
			row1["children"] = children1
			result = append(result, row1)
		}
	}
	return result, nil
}

//GetSystemInfo 获取系统信息
func (l *DbLogin) GetSystemInfo(ident string) (s db.QueryRow, err error) {
	db := l.c.GetRegularDB(config.DbName)
	data, _, _, err := db.Query(sqls.QuerySystemInfo, map[string]interface{}{
		"ident": ident,
	})

	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, errors.New("系统不存在或则系统被禁用")
	}
	return data.Get(0), err
}

//ChangePwd 修改密码
func (l *DbLogin) ChangePwd(userID int64, expassword, newpassword string) error {
	db := l.c.GetRegularDB(config.DbName)
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
