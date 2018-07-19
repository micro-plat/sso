package user

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/const/enum"
	"github.com/micro-plat/sso/modules/const/sql"
)

var _ IDbUser = &DbUser{}

type IDbUser interface {
	Query(input *QueryUserInput) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(userID int, status int) (err error)
	Get(userID int) (data db.QueryRow, err error)
	Delete(userID int) (err error)
	Edit(input *UserEditInput) (err error)
	Add(input *UserEditInput) (err error)
	CheckPWD(oldPwd string, userID int64) (err error)
}

//UserEditInput 编辑用户 输入参数
type UserEditInput struct {
	UserName string `form:"user_name" json:"user_name" valid:"ascii,required"`
	UserID   int64  `form:"user_id" json:"user_id"`
	RoleID   int64  `form:"role_id" json:"role_id" `
	Mobile   int64  `form:"mobile" json:"mobile" valid:"length(11|11),required"`
	Status   int    `form:"status" json:"status" valid:"required"`
	IsAdd    int    `form:"is_add" json:"is_add" valid:"required"`
	Auth     string `form:"auth" json:"auth" valid:"required"`
	Email    string `form:"email" json:"email" valid:"email,required"`
}

//QueryUserInput 查询用户列表输入参数
type QueryUserInput struct {
	PageIndex int    `form:"pi" json:"pi" valid:"required"`
	PageSize  int    `form:"ps" json:"ps" valid:"required"`
	UserName  string `form:"username" json:"username"`
	RoleID    string `form:"role_id" json:"role_id"`
}

type DbUser struct {
	c component.IContainer
}

func NewDbUser(c component.IContainer) *DbUser {
	return &DbUser{
		c: c,
	}
}

//Query 获取用户信息列表
func (u *DbUser) Query(input *QueryUserInput) (data db.QueryRows, count interface{}, err error) {
	db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"role_id":   input.RoleID,
		"user_name": " and t.user_name like '%" + input.UserName + "%'",
		"pi":        input.PageIndex,
		"ps":        input.PageSize,
	}
	count, q, a, err := db.Scalar(sql.QueryUserInfoListCount, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取用户信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	data, q, a, err = db.Query(sql.QueryUserInfoList, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取用户信息列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	sysRoles, q, a, err := db.Query(sql.QueryUserRoleList, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取用户信息列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	roles := make(map[string][]map[string]string)
	rolestr := make(map[string]string)
	for _, sysRole := range sysRoles {
		uid := sysRole.GetString("user_id")
		if _, ok := roles[uid]; !ok {
			roles[uid] = make([]map[string]string, 0, 2)
			rolestr[uid] = ""
		}
		roles[uid] = append(roles[uid], map[string]string{
			"sys_name":  sysRole.GetString("sys_name"),
			"role_name": sysRole.GetString("role_name"),
			"sys_id":    sysRole.GetString("sys_id"),
			"role_id":   sysRole.GetString("role_id"),
		})
		rolestr[uid] = rolestr[uid] + sysRole.GetString("sys_name") + "/" + sysRole.GetString("role_name") + ";"
	}
	for _, user := range data {
		uid := user.GetString("user_id")
		user["roles"] = roles[uid]
		user["rolestr"] = rolestr[uid]
	}
	return data, count, nil
}

//ChangeStatus 修改用户状态
func (u *DbUser) ChangeStatus(userID int, status int) (err error) {
	db := u.c.GetRegularDB()
	input := map[string]interface{}{
		"user_id": userID,
	}
	switch status {
	case enum.Disabled:
		input["status"] = enum.Disabled
	case enum.Normal, enum.Unlock:
		input["status"] = enum.Normal
	}
	_, q, a, err := db.Execute(sql.UpdateUserStatus, input)
	if err != nil {
		return fmt.Errorf("修改用户状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}

//Delete 删除用户
func (u *DbUser) Delete(userID int) (err error) {
	db := u.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.DeleteUser, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	_, q, a, err = dbTrans.Execute(sql.DelUserRole, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//Get 查询用户信息
func (u *DbUser) Get(userID int) (data db.QueryRow, err error) {
	db := u.c.GetRegularDB()
	result, q, a, err := db.Query(sql.QueryUserInfo, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return nil, fmt.Errorf("查询用户信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	return result.Get(0), nil
}

//Edit 编辑用户信息
func (u *DbUser) Edit(input *UserEditInput) (err error) {
	db := u.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}
	params, err := types.Struct2Map(input)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.EditUserInfo, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("编辑用户信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	_, q, a, err = dbTrans.Execute(sql.DelUserRole, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户原角色信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	as := strings.Split(input.Auth, "|")
	for i := 0; i < len(as)-1; i++ {
		as1 := strings.Split(as[i], ",")
		params["sys_id"] = as1[0]
		params["role_id"] = as1[1]
		_, q, a, err = dbTrans.Execute(sql.AddUserRole, params)
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	dbTrans.Commit()
	return nil
}

//Add 添加用户
func (u *DbUser) Add(input *UserEditInput) (err error) {
	db := u.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}
	params, err := types.Struct2Map(input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	n, _, _, err := dbTrans.Scalar(sql.GetNewUserID, map[string]interface{}{})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("获取新用户ID发生错误(err:%v)", err)
	}
	params["user_id"] = n.(string)
	params["password"] = md5.Encrypt(enum.UserDefaultPassword)

	_, q, a, err := dbTrans.Execute(sql.AddUserInfo, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加用户发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	as := strings.Split(input.Auth, "|")
	for i := 0; i < len(as)-1; i++ {
		as1 := strings.Split(as[i], ",")
		params["sys_id"] = as1[0]
		params["role_id"] = as1[1]
		_, q, a, err = dbTrans.Execute(sql.AddUserRole, params)
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	dbTrans.Commit()
	return nil
}

//CheckPWD 检查用户原密码是否匹配
func (u *DbUser) CheckPWD(oldPWD string, userID int64) (err error) {
	db := u.c.GetRegularDB()
	row, q, a, err := db.Scalar(sql.QueryUserPswd, map[string]interface{}{
		"password": md5.Encrypt(oldPWD),
		"user_id":  userID,
	})
	if err != nil {
		return context.NewError(406, fmt.Errorf("查询用户信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a))
	}
	if fmt.Sprint(row) != "1" {
		return context.NewError(403, fmt.Errorf("输入的原密码不正确"))
	}
	return nil
}
