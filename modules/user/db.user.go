package user

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/const/sql"
)

type IDbUser interface {
	Query(params map[string]interface{}) (data db.QueryRows, count interface{}, err error)
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
		"role_id":   input["roleid"],
		"user_name": input["username"],
		"pi":        input["pi"],
		"ps":        input["ps"],
	}
	count, q, a, err := db.Scalar(sql.QueryUserInfoListCount, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取用户信息列表条数发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sql.QueryUserInfoList, params)
	if err != nil {
		return nil, nil, fmt.Errorf("获取用户信息列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, count, nil
}
