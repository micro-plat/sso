package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
)

//MemberHandler is
type MemberHandler struct {
	container component.IContainer
	m         logic.IMemberLogic
}

//NewMemberHandler is
func NewMemberHandler(container component.IContainer) (u *MemberHandler) {
	return &MemberHandler{
		container: container,
		m:         logic.NewMemberLogic(container),
	}
}

/*
* GetHandle: 根据用户名查询用户的相关信息
* ident:子系统标识
* username:用户名称
 */
func (u *MemberHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统用户远程登录---------")

	if err := ctx.Request.Check("username", "ident"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("1. 根据用户名获取用户信息")
	member, err := u.m.QueryUserInfo(ctx.Request.GetString("username"),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3. 返回用户信息")
	return member
}
