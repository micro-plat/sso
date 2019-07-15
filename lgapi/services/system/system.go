package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/logic"
)

//SystemHandler 用户登录对象
type SystemHandler struct {
	c   component.IContainer
	sys logic.ISystemLogic
}

//NewSystemHandler 创建登录对象
func NewSystemHandler(container component.IContainer) (u *SystemHandler) {
	return &SystemHandler{
		c:   container,
		sys: logic.NewSystemLogic(container),
	}
}

//UserSystemHandle 返回用户可以访问的系统
func (u *SystemHandler) SystemHandle(ctx *context.Context) (r interface{}) {
	user := member.Get(ctx)
	if user == nil {
		return context.NewError(context.ERR_BAD_REQUEST, "登录信息出错,请重新登录")
	}

	data, err := u.sys.QueryUserSystem(user.UserID)
	if err != nil {
		return err
	}

	return data
}
