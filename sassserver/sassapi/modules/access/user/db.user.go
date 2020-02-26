package user

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/const/enum"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/const/sqls"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/config"
)

type IDbUser interface {
	Query(input *model.QueryUserInput) (data db.QueryRows, total int, err error)
	ChangeStatus(userID, status, belongID, belongType int) (err error)
	Get(userID int) (data db.QueryRow, err error)
	GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error)
	Delete(userID, belongID, belongType int) (err error)
	Edit(input *model.UserInputNew) (err error)
	Add(input *model.UserInputNew) (err error)
	ResetPwd(userID, belongID, belongType int) (err error)
	GetUserInfoByName(userName string) (data db.QueryRow, err error)
	GetUserInfoByFullName(fullName string) (data db.QueryRow, err error)
	GetUserInfoByMobile(mobile string) (data db.QueryRow, err error)
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
func (u *DbUser) Query(input *model.QueryUserInput) (data db.QueryRows, total int, err error) {
	db := u.c.GetRegularDB(config.DbName)
	params := map[string]interface{}{
		"role_id":     input.RoleID,
		"status":      input.Status,
		"mobile":      " and (t.mobile like '%" + input.Mobile + "%')",
		"belong_id":   input.BelongID,
		"belong_type": input.BelongType,
		"start":       (input.PageIndex - 1) * input.PageSize,
		"ps":          input.PageSize,
	}
	count, q, a, err := db.Scalar(sqls.QueryUserInfoListCount, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	data, q, a, err = db.Query(sqls.QueryUserInfoList, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户信息列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	params["user_id_string"] = ""
	//查询给定用户的角色数据
	if types.GetInt(count) > 0 {
		userids := make([]string, 0)
		for _, v := range data {
			userId := v.GetString("user_id")
			if userId == "" {
				continue
			}
			userids = append(userids, userId)
		}
		params["user_id_string"] = strings.Join(userids, ",")
	}

	sysRoles, q, a, err := db.Query(sqls.QueryUserRoleList, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户信息列表发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	rolestr := make(map[string]string)
	roleId := make(map[string]string)
	for _, sysRole := range sysRoles {
		uid := sysRole.GetString("user_id")
		rolestr[uid] = sysRole.GetString("role_name")
		roleId[uid] = sysRole.GetString("role_id")
	}
	for _, user := range data {
		uid := user.GetString("user_id")
		user["rolestr"] = rolestr[uid]
		user["role_id"] = roleId[uid]
	}
	return data, types.GetInt(count, 0), nil
}

//ChangeStatus 修改用户状态
func (u *DbUser) ChangeStatus(userID, status, belongID, belongType int) (err error) {
	db := u.c.GetRegularDB(config.DbName)
	_, q, a, err := db.Execute(sqls.UpdateUserStatus, map[string]interface{}{
		"user_id":     userID,
		"status":      status,
		"belong_id":   belongID,
		"belong_type": belongType,
	})

	if err != nil {
		return fmt.Errorf("修改用户状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	return nil
}

//Delete 删除用户
func (u *DbUser) Delete(userID, belongID, belongType int) (err error) {
	db := u.c.GetRegularDB(config.DbName)
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sqls.DeleteUser, map[string]interface{}{
		"user_id":     userID,
		"belong_id":   belongID,
		"belong_type": belongType,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	_, q, a, err = dbTrans.Execute(sqls.DelUserRole, map[string]interface{}{
		"user_id":     userID,
		"belong_id":   belongID,
		"belong_type": belongType,
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
	db := u.c.GetRegularDB(config.DbName)
	result, q, a, err := db.Query(sqls.QueryUserInfo, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return nil, fmt.Errorf("查询用户信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	if result.IsEmpty() {
		return nil, context.NewError(model.ERR_USER_NOTEXISTS, "用户不存在")
	}

	return result.Get(0), nil
}

func (u *DbUser) GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error) {
	db := u.c.GetRegularDB(config.DbName)
	c, q, a, err := db.Scalar(sqls.QueryUserBySysCount, map[string]interface{}{
		"sys_id": sysID,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	data, q, a, err = db.Query(sqls.QueryUserBySysList, map[string]interface{}{
		"sys_id": sysID,
		"pi":     pi,
		"ps":     ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, types.GetInt(c), nil

}

//Edit 编辑用户信息
func (u *DbUser) Edit(input *model.UserInputNew) (err error) {
	db := u.c.GetRegularDB(config.DbName)
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}
	params, err := types.Struct2Map(input)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sqls.EditUserInfo, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("编辑用户信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	_, q, a, err = dbTrans.Execute(sqls.DelUserRole, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户原角色信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	//现在系统只要两个角色(加油站管理员:0, 公司管理员:1), 后面会做成从界面上选择, 传入
	_, q, a, err = dbTrans.Execute(sqls.AddUserRole, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("关联用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//Add 添加用户
func (u *DbUser) Add(input *model.UserInputNew) (err error) {
	params, err := types.Struct2Map(input)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	db := u.c.GetRegularDB(config.DbName)
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	params["password"] = md5.Encrypt(enum.UserDefaultPassword)
	userID, _, q, a, err := dbTrans.Executes(sqls.AddUserInfo, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加用户发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	params["user_id"] = userID
	_, q, a, err = dbTrans.Execute(sqls.AddUserRole, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("关联用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//ResetPwd ResetPwd
func (u *DbUser) ResetPwd(userID, belongID, belongType int) (err error) {
	db := u.c.GetRegularDB(config.DbName)
	_, q, a, err := db.Execute(sqls.SetNewPwd, map[string]interface{}{
		"user_id":     userID,
		"password":    md5.Encrypt(enum.UserDefaultPassword),
		"belong_id":   belongID,
		"belong_type": belongType,
	})
	if err != nil {
		return fmt.Errorf("重置用户密码发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	return nil
}

//GetUserInfoByName 根据用户名查询用户信息
func (u *DbUser) GetUserInfoByName(userName string) (data db.QueryRow, err error) {
	db := u.c.GetRegularDB(config.DbName)
	result, _, _, err := db.Query(sqls.GetUserInfoByName, map[string]interface{}{"user_name": userName})
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() {
		return nil, nil
	}
	return result.Get(0), nil
}

//GetUserInfoByFullName 根据姓名查询用户信息
func (u *DbUser) GetUserInfoByFullName(fullName string) (data db.QueryRow, err error) {
	db := u.c.GetRegularDB(config.DbName)
	result, _, _, err := db.Query(sqls.GetUserInfoByFullName, map[string]interface{}{"full_name": fullName})
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() {
		return nil, nil
	}
	return result.Get(0), nil
}

//GetUserInfoByMobile 根据手机号查询用户信息
func (u *DbUser) GetUserInfoByMobile(mobile string) (data db.QueryRow, err error) {
	db := u.c.GetRegularDB(config.DbName)
	result, _, _, err := db.Query(sqls.GetUserInfoByMobile, map[string]interface{}{"mobile": mobile})
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() {
		return nil, nil
	}
	return result.Get(0), nil
}
