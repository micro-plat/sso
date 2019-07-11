package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
)

//UserInfoHandler is
type UserInfoHandler struct {
	container component.IContainer
	sys       logic.ISystemLogic
	m         logic.IMemberLogic
}

//NewUserInfoHandler is
func NewUserInfoHandler(container component.IContainer) (u *UserInfoHandler) {
	return &UserInfoHandler{
		container: container,
		sys:       logic.NewSystemLogic(container),
		m:         logic.NewMemberLogic(container),
	}
}

/*
* Handle: 根据用户名查询用户的相关信息
* ident:子系统标识
* username:用户名称
* timestamp:时间戳
* sign:签名
 */
func (u *UserInfoHandler) InfoHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统用户远程登录---------")

	if err := ctx.Request.Check("username"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	member, err := u.m.QueryUserInfo(ctx.Request.GetString("username"),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	return member
}

/*
* Handle: 根据登录后给子系统的key还回用户信息
* key:guid
* ident:子系统标识
* timestamp:时间戳
* sign:签名
 */
func (u *UserInfoHandler) KeyHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统远程通过key来拿取用户user_id,user_name---------")

	if err := ctx.Request.Check("key"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	info, err := u.m.GetUserInfoByKey(ctx.Request.GetString("key"))
	if err != nil {
		return err
	}

	return info
}
