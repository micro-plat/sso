package logic

import (
	"github.com/micro-plat/sso/loginserver/sdkapi/modules/access/user"
	"github.com/micro-plat/sso/loginserver/sdkapi/modules/model"
)

//IUserLogic 用户管理
type IUserLogic interface {
	AddUser(req model.UserInputNew) error
}

//UserLogic 用户管理
type UserLogic struct {
	db user.IDBUser
}

//NewUserLogic 创建登录对象
func NewUserLogic() *UserLogic {
	return &UserLogic{
		db: user.NewDBUser(),
	}
}

//AddUser 新增用户
func (m *UserLogic) AddUser(req model.UserInputNew) error {
	return m.db.AddUser(req)
}
