package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IUser interface {
	Query(input QueryUserInput) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(userID int, status int) (err error)
	Delete(userID int) (err error)
	UserInfo(userID int) (data interface{}, err error)
	UserEdit(input UserEditInput) (err error)
	CheckPswd(oldPwd string, newPwd string, userID int64) (code int, err error)
}

type User struct {
	c  component.IContainer
	db IDbUser
}

func NewUser(c component.IContainer) *User {
	return &User{
		c:  c,
		db: NewDbUser(c),
	}
}

//Query 获取用户信息列表
func (u *User) Query(input QueryUserInput) (data db.QueryRows, count interface{}, err error) {
	data, count, err = u.db.Query(input)
	if err != nil {
		return nil, nil, err
	}
	return data, count, nil
}

//CHangeStatus 修改用户状态
func (u *User) ChangeStatus(userID int, status int) (err error) {
	return u.db.ChangeStatus(userID, status)
}

//Delete 删除用户
func (u *User) Delete(userID int) (err error) {
	return u.db.Delete(userID)
}

//UserInfo 查询用户信息
func (u *User) UserInfo(userID int) (data interface{}, err error) {
	data, err = u.db.UserInfo(userID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//UserEdit 编辑用户信息
func (u *User) UserEdit(input UserEditInput) (err error) {
	if input.IsAdd == 1 {
		return u.db.Add(input)
	} else {
		return u.db.Edit(input)
	}
}

//CheckPswd 检查用户原密码是否匹配
func (u *User) CheckPswd(oldPwd string, newPwd string, userID int64) (code int, err error) {
	code, err = u.db.CheckPswd(oldPwd, newPwd, userID)
	if err != nil {
		return code, err
	}
	return 200, nil
}
