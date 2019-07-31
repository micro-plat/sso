package member

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/lgapi/modules/const/enum"
	"github.com/micro-plat/sso/lgapi/modules/const/sqls"
	"github.com/micro-plat/sso/lgapi/modules/model"
)

type IDBMember interface {
	Query(u, p, ident string) (s *model.MemberState, err error)
	ChangePwd(userID int, expassword string, newpassword string) (err error)
	QueryByID(uid int64) (db.QueryRow, error)
	CheckUserHasAuth(ident string, userID int64) error
	QueryByOpenID(openID, ident string) (s *model.MemberState, err error)
	ExistsOpenId(content string) error
	QueryByName(userName, ident string) (s *model.MemberState, err error)

	QueryByUserName(u string, ident string) (info db.QueryRow, err error)
	GetUserInfo(u string) (db.QueryRow, error)

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
func (l *DBMember) Query(u, p, ident string) (s *model.MemberState, err error) {
	db := l.c.GetRegularDB()
	data, _, _, errt := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
		"user_name": u,
	})
	if errt != nil {
		return nil, context.NewError(context.ERR_SERVICE_UNAVAILABLE, "暂时无法登录系统")
	}
	if data.IsEmpty() {
		return nil, context.NewError(context.ERR_BAD_REQUEST, "用户名或密码错误")
	}

	row := data.Get(0)
	s = &model.MemberState{
		UserID:    row.GetInt64("user_id", -1),
		Password:  row.GetString("password"),
		UserName:  row.GetString("user_name"),
		ExtParams: row.GetString("ext_params"),
		Status:    row.GetInt("status"),
	}

	//处理如果是子系统传系统编号登录就要判断权限
	params := map[string]interface{}{
		"user_id": data.Get(0).GetInt64("user_id", -1),
		"ident":   " and 1=1 ",
	}
	if ident != "" {
		params["ident"] = " and s.ident='" + ident + "' "
	}

	roles, _, _, erro := db.Query(sqls.QueryUserRole, params)
	if erro != nil || roles.IsEmpty() {
		return nil, context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, "用户没有相关系统权限,请联系管理员")
	}

	if ident != "" {
		s.SysIdent = ident
		s.SystemID = roles.Get(0).GetInt("sys_id")
		s.RoleName = roles.Get(0).GetString("role_name")
		s.IndexURL = roles.Get(0).GetString("index_url")
		s.LoginURL = roles.Get(0).GetString("login_url")
	}

	return s, err
}

// ChangePwd 修改密码
func (l *DBMember) ChangePwd(userID int, expassword string, newpassword string) (err error) {
	db := l.c.GetRegularDB()

	data, _, _, err := db.Query(sqls.QueryOldPwd, map[string]interface{}{
		"user_id": userID,
		"ident":   " and 1=1 ",
	})

	//data.Get(0).GetInt("changepwd_times") >= 3
	if err != nil {
		return context.NewError(context.ERR_BAD_REQUEST, "用户不存在, 修改失败")
	}

	if strings.ToLower(md5.Encrypt(expassword)) != strings.ToLower(data.Get(0).GetString("password")) {
		return context.NewError(context.ERR_BAD_REQUEST, "原密码错误")
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
	params := map[string]interface{}{"user_id": userID, "ident": " and 1=1 "}
	if ident != "" {
		params["ident"] = " and s.ident='" + ident + "' "
	}
	fmt.Println(params)

	db := l.c.GetRegularDB()
	count, _, _, err := db.Scalar(sqls.QueryUserRoleCount, params)
	if err != nil {
		return context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, fmt.Sprintf("出现错误，等会在登录: %s", err))
	}
	if types.GetInt(count, 0) <= 0 {
		return context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, "没有相应权限，请联系管理员")
	}
	return nil
}

//QueryByOpenID 根据openid 查询用户信息
func (l *DBMember) QueryByOpenID(openID, ident string) (s *model.MemberState, err error) {
	db := l.c.GetRegularDB()

	//根据用户名密码，查询用户信息
	data, _, _, err := db.Query(sqls.QueryUserInfoByOpenID, map[string]interface{}{
		"open_id": openID,
	})
	if err != nil {
		return nil, context.NewError(context.ERR_SERVER_ERROR, fmt.Sprintf("用openid查询时出错:%v+", err))
	}
	if data.IsEmpty() {
		return nil, context.NewError(context.ERR_UNAUTHORIZED, "没有关注公众号，先关注公众号")
	}

	row := data.Get(0)

	//检查用户是否已锁定
	if row.GetInt("status") == enum.UserLock || row.GetInt("status") == enum.UserDisable {
		return nil, context.NewError(context.ERR_BAD_REQUEST, "用户被锁定或被禁用，暂时无法登录")
	}

	s = &model.MemberState{
		UserID:    row.GetInt64("user_id", -1),
		Password:  row.GetString("password"),
		UserName:  row.GetString("user_name"),
		ExtParams: row.GetString("ext_params"),
		Status:    row.GetInt("status"),
	}

	//处理如果是子系统传系统编号登录就要判断权限
	params := map[string]interface{}{
		"user_id": row.GetInt64("user_id", -1),
		"ident":   " and 1=1 ",
	}
	if ident != "" {
		params["ident"] = " and s.ident='" + ident + "' "
	}

	roles, _, _, erro := db.Query(sqls.QueryUserRole, params)
	if erro != nil || roles.IsEmpty() {
		return nil, context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, "用户没有相关系统权限,请联系管理员")
	}

	if ident != "" {
		s.SysIdent = ident
		s.SystemID = roles.Get(0).GetInt("sys_id")
		s.RoleName = roles.Get(0).GetString("role_name")
		s.IndexURL = roles.Get(0).GetString("index_url")
		s.LoginURL = roles.Get(0).GetString("login_url")
	}
	return s, err
}

//ExistsOpenId xx
func (l *DBMember) ExistsOpenId(content string) error {
	contentMap := map[string]interface{}{}
	if err := json.Unmarshal([]byte(content), &contentMap); err != nil {
		return context.NewError(context.ERR_SERVER_ERROR, err)
	}
	openID, ok := contentMap["openid"]
	if !ok {
		return context.NewError(context.ERR_SERVER_ERROR, "openid不存在")
	}

	fmt.Printf("opid:%s", openID)
	db := l.c.GetRegularDB()
	count, _, _, err := db.Scalar(sqls.ExistsUserByOpenId, map[string]interface{}{
		"openid": openID,
	})
	fmt.Printf("count:%v", count)
	if err != nil {
		return context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, fmt.Sprintf("出现错误，等会在登录: %s", err))
	}
	if types.GetInt(count, 0) <= 0 {
		return context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, "没有绑定公众号,请先绑定")
	}
	return nil
}

//QueryByName xx
func (l *DBMember) QueryByName(userName, ident string) (s *model.MemberState, err error) {
	db := l.c.GetRegularDB()

	data, _, _, err := db.Query(sqls.QueryUserByUserName,
		map[string]interface{}{"user_name": userName})

	if err != nil {
		return nil, context.NewError(context.ERR_SERVER_ERROR, fmt.Sprintf("QueryByName访问出错: %v+", err))
	}
	if data.IsEmpty() {
		return nil, context.NewError(context.ERR_UNAUTHORIZED, "用户名不存在")
	}
	row := data.Get(0)
	if row.GetInt("status") == enum.UserLock || row.GetInt("status") == enum.UserDisable {
		return nil, context.NewError(context.ERR_UNAUTHORIZED, "用户被锁定或被禁用，暂时无法登录")
	}

	s = &model.MemberState{
		UserID:    row.GetInt64("user_id", -1),
		Password:  row.GetString("password"),
		UserName:  row.GetString("user_name"),
		ExtParams: row.GetString("ext_params"),
		Status:    row.GetInt("status"),
	}

	//处理如果是子系统传系统编号登录就要判断权限
	params := map[string]interface{}{
		"user_id": row.GetInt64("user_id", -1),
		"ident":   " and 1=1 ",
	}
	if ident != "" {
		params["ident"] = " and s.ident='" + ident + "' "
	}

	roles, _, _, erro := db.Query(sqls.QueryUserRole, params)
	if erro != nil || roles.IsEmpty() {
		return nil, context.NewError(context.ERR_UNSUPPORTED_MEDIA_TYPE, "用户没有相关系统权限,请联系管理员")
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
	data, _, _, err := db.Query(sqls.QueryUserByUserName, map[string]interface{}{
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
