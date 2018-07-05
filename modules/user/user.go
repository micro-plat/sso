package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IUser interface {
	QueryUserList(params map[string]interface{}) (data db.QueryRows, count interface{}, err error)
}

type User struct {
	c component.IContainer
}

func NewUser(c component.IContainer) *User {
	return &User{
		c: c,
	}
}

//QueryUserList 获取用户信息列表
func (u *User) QueryUserList(params map[string]interface{}) (data db.QueryRows, count interface{}, err error) {
	params["username_sql"] = " and t.user_name like '%" + params["username"].(string) + "%'"
	data, count, err = u.DbQueryUserList(params)
	if err != nil {
		return nil, nil, err
	}
	return data, count, nil
}
