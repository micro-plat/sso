package user

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/modules/const/sql"
	"github.com/micro-plat/sso/modules/const/util"
)

type IDbUser interface {
	Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error)
	CHangeStatus(input map[string]interface{}) (err error)
	UserInfo(input map[string]interface{}) (data interface{}, err error)
	Delete(input map[string]interface{}) (err error)
	Edit(input map[string]interface{}) (err error)
	Add(input map[string]interface{}) (err error)
	CheckPswd(input map[string]interface{}) (code int, err error)
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
func (u *DbUser) Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error) {
	db := u.c.GetRegularDB()
	params := map[string]interface{}{
		"role_id":   input["role_id"],
		"user_name": " and t.user_name like '%" + input["username"].(string) + "%'",
		"pi":        input["pi"],
		"ps":        input["ps"],
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

//CHangeStatus 修改用户状态
func (u *DbUser) CHangeStatus(input map[string]interface{}) (err error) {
	if input["ex_status"].(float64) == util.UserDisabled || input["ex_status"].(float64) == util.UserLocked {
		input["status"] = util.UserNormal
	} else if input["ex_status"].(float64) == util.UserNormal {
		input["status"] = util.UserDisabled
	}

	db := u.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.UpdateUserStatus, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("修改用户状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//Delete 删除用户
func (u *DbUser) Delete(input map[string]interface{}) (err error) {
	db := u.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.DeleteUser, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	_, q, a, err = dbTrans.Execute(sql.DelUserRole, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	dbTrans.Commit()
	return nil
}

//UserInfo 查询用户信息
func (u *DbUser) UserInfo(input map[string]interface{}) (data interface{}, err error) {
	db := u.c.GetRegularDB()
	data, q, a, err := db.Scalar(sql.QueryUserInfo, input)
	if err != nil {
		return nil, fmt.Errorf("查询用户信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	return data, nil
}

//Edit 编辑用户信息
func (u *DbUser) Edit(input map[string]interface{}) (err error) {
	db := u.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, q, a, err := dbTrans.Execute(sql.EditUserInfo, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("编辑用户信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	_, q, a, err = dbTrans.Execute(sql.DelUserRole, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户原角色信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	as := strings.Split(input["auth"].(string), "|")
	for i := 0; i < len(as)-1; i++ {
		as1 := strings.Split(as[i], ",")
		input["sys_id"] = as1[0]
		input["role_id"] = as1[1]
		_, q, a, err = dbTrans.Execute(sql.AddUserRole, input)
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	dbTrans.Commit()
	return nil
}

//Add 添加用户
func (u *DbUser) Add(input map[string]interface{}) (err error) {
	db := u.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	n, _, _, err := dbTrans.Scalar(sql.GetNewUserID, map[string]interface{}{})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("获取新用户ID发生错误(err:%v)", err)
	}
	input["user_id"] = n.(string)
	input["password"] = md5.Encrypt(util.UserDefaultPassword)

	fmt.Println("Adduser:", input)
	_, q, a, err := dbTrans.Execute(sql.AddUserInfo, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加用户发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	as := strings.Split(input["auth"].(string), "|")
	for i := 0; i < len(as)-1; i++ {
		as1 := strings.Split(as[i], ",")
		input["sys_id"] = as1[0]
		input["role_id"] = as1[1]
		_, q, a, err = dbTrans.Execute(sql.AddUserRole, input)
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加用户角色发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	dbTrans.Commit()
	return nil
}

//CheckPswd 检查用户原密码是否匹配
func (u *DbUser) CheckPswd(input map[string]interface{}) (code int, err error) {
	db := u.c.GetRegularDB()
	row, q, a, err := db.Scalar(sql.QueryUserPswd, input)
	if err != nil {
		return 406, fmt.Errorf("查询用户信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	data := md5.EncryptBytes([]byte(row.(string)))
	if input["expassword"].(string) != data {
		return 403, fmt.Errorf("输入的原密码不正确")
	}
	return 400, nil
}
