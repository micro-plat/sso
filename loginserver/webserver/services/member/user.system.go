package member

import (
	"github.com/micro-plat/hydra"
	m "github.com/micro-plat/sso/loginserver/webserver/modules/access/member"
	"github.com/micro-plat/sso/loginserver/webserver/modules/logic"
)

//UserSysHandler 用户对象
type UserSysHandler struct {
	sys logic.ISystemLogic
}

//NewUserSysHandler 用户
func NewUserSysHandler() (u *UserSysHandler) {
	return &UserSysHandler{
		sys: logic.NewSystemLogic(),
	}
}

//Handle 返回用户可以访问的系统
func (u *UserSysHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------返回用户可以访问的子系统---------")
	data, err := u.sys.QueryUserSystem(m.Get(ctx).UserID)
	if err != nil {
		return err
	}
	return data
}
