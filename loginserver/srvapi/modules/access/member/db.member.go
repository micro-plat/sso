package member

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	commsqls "github.com/micro-plat/sso/loginserver/loginapi/modules/const/sqls"
    "github.com/micro-plat/sso/loginserver/loginapi/modules/const/errorcode"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/const/sqls"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/model"
)

type IDBMember interface {
	QueryByUserName(u string, ident string) (info db.QueryRow, err error)
	QueryByID(uid int, ident string) (s *model.MemberState, err error)
	QueryUserSystem(userID int, ident string) (s db.QueryRows, err error)
	QueryAllUserInfo(source string, sourceID string) (s db.QueryRows, err error)
	GetAllUserInfoByUserRole(userID int, ident string) (string, error)
}

//DBMember 控制用户登录
type DBMember struct {
}

//NewDBMember 创建登录对象
func NewDBMember() *DBMember {
	return &DBMember{}
}

//QueryAllUserInfo 获取全部用户
func (l *DBMember) QueryAllUserInfo(source string, sourceID string) (s db.QueryRows, err error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryAllUserInfo, map[string]interface{}{
		"source":    source,
		"source_id": sourceID,
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

// QueryByUserName 根据用户名查询用户信息
func (l *DBMember) QueryByUserName(u string, ident string) (info db.QueryRow, err error) {
	//根据用户名，查询用户信息
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(commsqls.QueryUserByUserName, map[string]interface{}{
		"user_name": u,
	})
	if err != nil {
		return nil, errs.NewError(http.StatusServiceUnavailable, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, errs.NewError(http.StatusForbidden, "用户不存在")
	}
	//查询用户所在系统的登录地址及角色编号
	roles, err := db.Query(commsqls.QueryUserRole, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})
	if roles.IsEmpty() {
		return nil, errs.NewError(http.StatusUnsupportedMediaType, "不允许登录系统")
	}

	userData := data.Get(0)
	userData["ident"] = ident

	return userData, err
}

//QueryUserSystem 查询用户可用的子系统
func (l *DBMember) QueryUserSystem(userID int, ident string) (s db.QueryRows, err error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryUserSystem, map[string]interface{}{
		"user_id": userID,
		"ident":   ident,
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

// QueryByID 根据userid查询用户信息
func (l *DBMember) QueryByID(uid int, ident string) (s *model.MemberState, err error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(
		commsqls.QueryUserInfoByUID, map[string]interface{}{
			"user_id": uid,
		})
	if err != nil {
		return nil, errs.NewError(http.StatusServiceUnavailable, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, errs.NewError(errorcode.ERR_USER_NOTEXISTS, "用户不存在")
	}

	s = &model.MemberState{}

	//查询用户所在系统的登录地址及角色编号
	roles, err := db.Query(commsqls.QueryUserRole, map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   ident,
	})

	if err != nil {
		return nil, errs.NewError(http.StatusUnsupportedMediaType, "查询权限出错")
	}

	if roles.IsEmpty() {
		return nil, errs.NewError(errorcode.ERR_USER_HASNOROLES, "用户没有相关系统权限")
	}

	s.RoleID = roles.Get(0).GetInt("role_id")
	s.RoleName = roles.Get(0).GetString("role_name")
	s.IndexURL = roles.Get(0).GetString("index_url")
	s.SystemID = roles.Get(0).GetInt("sys_id")
	s.UserID = data.Get(0).GetInt64("user_id", -1)
	s.Status = data.Get(0).GetInt("status")
	s.Password = data.Get(0).GetString("password")
	s.FullName = data.Get(0).GetString("full_name")
	s.UserName = data.Get(0).GetString("user_name")
	s.ExtParams = data.Get(0).GetString("ext_params")
	s.SysIdent = ident
	s.LastLoginTime = data.Get(0).GetString("last_login_time")

	return
}

//GetAllUserInfoByUserRole 获取和当前用户同一个角色的用户ids
func (l *DBMember) GetAllUserInfoByUserRole(userID int, ident string) (string, error) {
	db := components.Def.DB().GetRegularDB()
	userInfo, err := db.Query(sqls.GetAllUserInfoByUserRole, map[string]interface{}{
		"user_id": userID,
		"ident":   ident,
	})
	if err != nil {
		return "", fmt.Errorf("GetAllUserInfoByUserRole出错: err:%+v", err)
	}
	var userIDArray []string
	for _, item := range userInfo {
		userIDArray = append(userIDArray, item.GetString("user_id"))
	}
	return strings.Join(userIDArray, ","), nil
}
