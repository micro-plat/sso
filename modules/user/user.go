package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type IUser interface {
	Query(params map[string]interface{}) (data db.QueryRows, count interface{}, err error)
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
