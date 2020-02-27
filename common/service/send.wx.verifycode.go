package service

import (
	"github.com/micro-plat/hydra/component"
	l "github.com/micro-plat/sso/common/module/login"
)

//SendWxVerifyCode 修改密码
func SendWxVerifyCode(c component.IContainer, userName, openID, ident string) (err error) {
	return l.NewLoginLogic(c).SendWxValidCode(userName, openID, ident)
}
