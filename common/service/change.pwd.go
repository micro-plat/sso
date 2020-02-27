package service

import (
	"github.com/micro-plat/hydra/component"
	l "github.com/micro-plat/sso/common/module/login"
)

//ChangePwd 修改密码
func ChangePwd(c component.IContainer, userID int, expassword string, newpassword string) (err error) {
	return l.NewLoginLogic(c).ChangePwd(userID, expassword, newpassword)
}
