package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	m "github.com/micro-plat/sso/loginserver/lgapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/logic"
)

//UserSysHandler 用户对象
type UserSysHandler struct {
	c   component.IContainer
	sys logic.ISystemLogic
}

//NewUserSysHandler 用户
func NewUserSysHandler(container component.IContainer) (u *UserSysHandler) {
	return &UserSysHandler{
		c:   container,
		sys: logic.NewSystemLogic(container),
	}
}

//Handle 返回用户可以访问的系统
func (u *UserSysHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------返回用户可以访问的子系统---------")

	data, err := u.sys.QueryUserSystem(m.Get(ctx).UserID)
	if err != nil {
		return err
	}
	return data
}
