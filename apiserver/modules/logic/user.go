package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/modules/access/user"
)

type IUserLogic interface {
	ChangePwd(user_id int, expassword string, newpassword string) (err error)
}

type UserLogic struct {
	c  component.IContainer
	db user.IDbUser
}

// NewUserLogic logic
func NewUserLogic(c component.IContainer) *UserLogic {
	return &UserLogic{
		c:  c,
		db: user.NewDbUser(c),
	}
}

// ChangePwd 修改密码
func (u *UserLogic) ChangePwd(user_id int, expassword string, newpassword string) (err error) {
	return u.db.ChangePwd(user_id, expassword, newpassword)
}
