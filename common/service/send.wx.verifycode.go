package service

import (
	l "github.com/micro-plat/sso/common/module/login"
)

//SendWxVerifyCode 修改密码
func SendWxVerifyCode(userName, openID, ident string) (err error) {
	return l.NewLoginLogic().SendWxValidCode(userName, openID, ident)
}
