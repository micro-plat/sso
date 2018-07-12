package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IUser interface {
	Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error)
	CHangeStatus(input map[string]interface{}) (err error)
	Delete(input map[string]interface{}) (err error)
	UserInfo(input map[string]interface{}) (data interface{}, err error)
	UserEdit(input map[string]interface{}) (err error)
	CheckPswd(input map[string]interface{}) (code int, err error)
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
func (u *User) Query(input map[string]interface{}) (data db.QueryRows, count interface{}, err error) {
	data, count, err = u.db.Query(input)
	if err != nil {
		return nil, nil, err
	}
	return data, count, nil
}

//CHangeStatus 修改用户状态
func (u *User) CHangeStatus(input map[string]interface{}) (err error) {
	err = u.db.CHangeStatus(input)
	if err != nil {
		return err
	}
	return nil
}

//Delete 删除用户
func (u *User) Delete(input map[string]interface{}) (err error) {
	err = u.db.Delete(input)
	if err != nil {
		return err
	}
	return nil
}

//UserInfo 查询用户信息
func (u *User) UserInfo(input map[string]interface{}) (data interface{}, err error) {
	data, err = u.db.UserInfo(input)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//UserEdit 编辑用户信息
func (u *User) UserEdit(input map[string]interface{}) (err error) {
	if input["is_add"].(float64) == 1 {
		err = u.db.Add(input)
		if err != nil {
			return err
		}
	} else {
		err = u.db.Edit(input)
		if err != nil {
			return err
		}
	}
	return nil
}

//CheckPswd 检查用户原密码是否匹配
func (u *User) CheckPswd(input map[string]interface{}) (code int, err error) {
	code, err = u.db.CheckPswd(input)
	if err != nil {
		return code, err
	}
	return 200, nil
}
