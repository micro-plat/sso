package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/logic"
)

//UserHandler 用户对象
type UserHandler struct {
	c   component.IContainer
	sys logic.ISystemLogic
	mem logic.IMemberLogic
}

//NewUserHandler 创建登录对象
func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		c:   container,
		sys: logic.NewSystemLogic(container),
		mem: logic.NewMemberLogic(container),
	}
}

//SystemHandle 返回用户可以访问的系统
func (u *UserHandler) SystemHandle(ctx *context.Context) (r interface{}) {
	user := member.Get(ctx)
	if user == nil {
		return context.NewError(context.ERR_FORBIDDEN, "登录信息出错,请重新登录")
	}

	data, err := u.sys.QueryUserSystem(user.UserID)
	if err != nil {
		return err
	}

	return data
}

//ChangePwdHandle 修改用户密码
func (u *UserHandler) ChangePwdHandle(ctx *context.Context) (r interface{}) {
	user := member.Get(ctx)
	if user == nil {
		return context.NewError(context.ERR_FORBIDDEN, "登录信息出错,请重新登录")
	}

	if err := ctx.Request.Check("expassword", "newpassword"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	err := u.mem.ChangePwd(int(user.UserID), ctx.Request.GetString("expassword"), ctx.Request.GetString("newpassword"))
	if err != nil {
		er, flag := err.(context.Error)
		if flag {
			if er.GetCode() == context.ERR_SERVER_ERROR {
				ctx.Log.Error("修改密码发生错误: %v", er)
				return context.NewError(context.ERR_NOT_ACCEPTABLE, "设置密码错误,等稍后再试")
			}
		}
		return err
	}
	return "success"
}

//CodeHandle 返回用户的身份code(这个是子系统选择页面，返回一个登录标识给子系统)
func (u *UserHandler) CodeHandle(ctx *context.Context) (r interface{}) {
	user := member.Get(ctx)
	if user == nil {
		return context.NewError(context.ERR_FORBIDDEN, "登录信息出错,请重新登录")
	}

	ctx.Log.Info("1: 设置已登录code")
	code, err := u.mem.CreateLoginUserCode(user.UserID)
	if err != nil {
		return context.NewError(context.ERR_BAD_REQUEST, err.Error)
	}

	return code
}
