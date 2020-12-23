package service

import (
	l "github.com/micro-plat/sso/common/module/login"
)

//ChangePwd 修改密码
func ChangePwd(userID int, expassword string, newpassword string) (err error) {
	return l.NewLoginLogic().ChangePwd(userID, expassword, newpassword)
}
