package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IUser interface {
	Query(input *QueryUserInput) (data db.QueryRows, count interface{}, err error)
	ChangeStatus(userID int, status int) (err error)
	Delete(userID int) (err error)
	Get(userID int) (data db.QueryRow, err error)
	Save(input *UserEditInput) (err error)
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
func (u *User) Query(input *QueryUserInput) (data db.QueryRows, count interface{}, err error) {
	data, count, err = u.db.Query(input)
	if err != nil {
		return nil, nil, err
	}
	return data, count, nil
}

//ChangeStatus 修改用户状态
func (u *User) ChangeStatus(userID int, status int) (err error) {
	return u.db.ChangeStatus(userID, status)
}

//Delete 删除用户
func (u *User) Delete(userID int) (err error) {
	return u.db.Delete(userID)
}

//Get 查询用户信息
func (u *User) Get(userID int) (data db.QueryRow, err error) {
	return u.db.Get(userID)
}

//Save 保存用户信息
func (u *User) Save(input *UserEditInput) (err error) {
	if input.IsAdd == 1 {
		return u.db.Add(input)
	}
	return u.db.Edit(input)
}
