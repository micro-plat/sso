package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/access/user"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

//IUserLogic 用户管理
type IUserLogic interface {
	AddUser(req model.UserInputNew) error
}

//UserLogic 用户管理
type UserLogic struct {
	//cache member.ICacheMember
	db user.IDBUser
}

//NewUserLogic 创建登录对象
func NewUserLogic(c component.IContainer) *UserLogic {
	return &UserLogic{
		//cache: member.NewCacheMember(c),
		db: user.NewDBUser(c),
	}
}

//AddUser 新增用户
func (m *UserLogic) AddUser(req model.UserInputNew) error {
	return m.db.AddUser(req)
}
