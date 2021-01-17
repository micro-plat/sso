package user

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/enum"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/sqls"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/errorcode"
	
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

type IDbUser interface {
	Query(input *model.QueryUserInput) (data db.QueryRows, total int, err error)
	ChangeStatus(userID int, status int) (err error)
	Get(userID int) (data db.QueryRow, err error)
	GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error)
	Delete(userID int) (err error)
	Edit(input *model.UserInputNew) (err error)
	Add(input *model.UserInputNew) (err error)
	EditInfo(username string, tel string, email string) (err error)
	ResetPwd(userID int) (err error)
	GetUserInfoByName(userName string) (data db.QueryRow, err error)
	GetUserInfoByFullName(fullName string) (data db.QueryRow, err error)
}

type DbUser struct{}

func NewDbUser() *DbUser {
	return &DbUser{}
}

//Query 获取用户信息列表
func (u *DbUser) Query(input *model.QueryUserInput) (data db.QueryRows, total int, err error) {
	db := components.Def.DB().GetRegularDB()
	params := map[string]interface{}{
		"role_id":   input.RoleID,
		"status":    input.Status,
		"user_name": " and (t.user_name like '%" + input.UserName + "%' or t.full_name like '%" + input.UserName + "%')",
		"start":     (input.PageIndex - 1) * input.PageSize,
		"ps":        input.PageSize,
	}
	count, err := db.Scalar(sqls.QueryUserInfoListCount, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户信息列表条数发生错误(err:%v)", err)
	}
	data, err = db.Query(sqls.QueryUserInfoList, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户信息列表发生错误(err:%v)", err)
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

	sysRoles, err := db.Query(sqls.QueryUserRoleList, params)
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户信息列表发生错误(err:%v)", err)
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
	return data, types.GetInt(count, 0), nil
}

//ChangeStatus 修改用户状态
func (u *DbUser) ChangeStatus(userID int, status int) (err error) {
	db := components.Def.DB().GetRegularDB()
	_, err = db.Execute(sqls.UpdateUserStatus, map[string]interface{}{
		"user_id": userID,
		"status":  status,
	})

	if err != nil {
		return fmt.Errorf("修改用户状态发生错误(err:%v)", err)
	}
	return nil
}

//Delete 删除用户
func (u *DbUser) Delete(userID int) (err error) {
	db := components.Def.DB().GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	_, err = dbTrans.Execute(sqls.DeleteUser, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户发生错误(err:%v)", err)
	}

	_, err = dbTrans.Execute(sqls.DelUserRole, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户角色发生错误(err:%v)", err)
	}

	dbTrans.Commit()
	return nil
}

//Get 查询用户信息
func (u *DbUser) Get(userID int) (data db.QueryRow, err error) {
	db := components.Def.DB().GetRegularDB()
	result, err := db.Query(sqls.QueryUserInfo, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return nil, fmt.Errorf("查询用户信息发生错误(err:%v)", err)
	}

	if result.IsEmpty() {
		return nil, errs.NewError(errorcode.ERR_USER_NOTEXISTS, fmt.Sprintf("用户不存在:%d",userID))
	}

	return result.Get(0), nil
}

func (u *DbUser) GetAll(sysID, pi, ps int) (data db.QueryRows, count int, err error) {
	db := components.Def.DB().GetRegularDB()
	c, err := db.Scalar(sqls.QueryUserBySysCount, map[string]interface{}{
		"sys_id": sysID,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户信息列表条数发生错误(err:%v)", err)
	}
	data, err = db.Query(sqls.QueryUserBySysList, map[string]interface{}{
		"sys_id": sysID,
		"pi":     pi,
		"ps":     ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户列表发生错误(err:%v)", err)
	}
	return data, types.GetInt(c), nil

}

//Edit 编辑用户信息
func (u *DbUser) Edit(input *model.UserInputNew) (err error) {
	db := components.Def.DB().GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}
	params, err := types.Struct2Map(input)
	if err != nil {
		return fmt.Errorf("Struct2Map Error(err:%v)", err)
	}

	_, err = dbTrans.Execute(sqls.EditUserInfo, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("编辑用户信息发生错误(err:%v)", err)
	}

	_, err = dbTrans.Execute(sqls.DelUserRole, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("删除用户原角色信息发生错误(err:%v)", err)
	}

	as := strings.Split(input.Auth, "|")
	for i := 0; i < len(as)-1; i++ {
		as1 := strings.Split(as[i], ",")
		params["sys_id"] = as1[0]
		params["role_id"] = as1[1]
		_, err = dbTrans.Execute(sqls.AddUserRole, params)
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加用户角色发生错误(err:%v)", err)
		}
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

	db := components.Def.DB().GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}

	params["password"] = md5.Encrypt(enum.UserDefaultPassword)
	user_id, _, err := dbTrans.Executes(sqls.AddUserInfo, params)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加用户发生错误(err:%v)", err)
	}

	params["user_id"] = user_id
	as := strings.Split(input.Auth, "|")
	for i := 0; i < len(as)-1; i++ {
		as1 := strings.Split(as[i], ",")
		params["sys_id"] = as1[0]
		params["role_id"] = as1[1]
		_, err = dbTrans.Execute(sqls.AddUserRole, params)
		if err != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加用户角色发生错误(err:%v)", err)
		}
	}

	dbTrans.Commit()
	return nil
}

//EditInfo EditInfo
func (u *DbUser) EditInfo(username string, tel string, email string) (err error) {
	db := components.Def.DB().GetRegularDB()
	params := map[string]interface{}{
		"username": username,
		"tel":      tel,
		"email":    email,
	}
	_, err = db.Execute(sqls.EditInfo, params)
	if err != nil {
		return fmt.Errorf("编辑个人资料发生错误(err:%v)", err)
	}
	return nil

}

//ResetPwd ResetPwd
func (u *DbUser) ResetPwd(userID int) (err error) {
	db := components.Def.DB().GetRegularDB()
	_, err = db.Execute(sqls.SetNewPwd, map[string]interface{}{
		"user_id":  userID,
		"password": md5.Encrypt(enum.UserDefaultPassword),
	})
	if err != nil {
		return fmt.Errorf("重置用户密码发生错误(err:%v)", err)
	}
	return nil
}

//GetUserInfoByName 根据用户名查询用户信息
func (u *DbUser) GetUserInfoByName(userName string) (data db.QueryRow, err error) {
	db := components.Def.DB().GetRegularDB()
	result, err := db.Query(sqls.GetUserInfoByName, map[string]interface{}{"user_name": userName})
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
	db := components.Def.DB().GetRegularDB()
	result, err := db.Query(sqls.GetUserInfoByFullName, map[string]interface{}{"full_name": fullName})
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() {
		return nil, nil
	}
	return result.Get(0), nil
}
