package user

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"
	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/logic"
	"github.com/micro-plat/sso/lgapi/modules/model"
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

//CheckHandle 微信绑定先期验证(无需登录)
func (u *UserHandler) CheckHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info(": ---------用户绑定微信账户前，验证用户名密码等----------------")

	if err := ctx.Request.Check("username", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "用户名和密码不能为空")
	}

	ctx.Log.Info("1: 验证用户名密码,用户的状态以及openid")
	if err := u.mem.CheckUerInfo(
		ctx.Request.GetString("username"),
		ctx.Request.GetString("password")); err != nil {
		return err
	}

	stateCode := utility.GetGUID()
	ctx.Log.Info("1: 将stateCode存到缓存中,wx会将这个还回,用于判断是否伪造")
	if err := u.mem.SaveWxStateCode(stateCode, ctx.Request.GetString("username")); err != nil {
		ctx.Log.Errorf("保存 statecode失败：%v+", err)
		return context.NewError(context.ERR_SERVER_ERROR, "系统繁忙，等会在绑定")
	}

	ctx.Log.Info("3: 生成statecode")
	config := model.GetConf(u.c)
	return map[string]interface{}{
		"wxlogin_url": config.WxPhoneLoginUrl,
		"appid":       config.Appid,
		"state":       stateCode,
	}
}

//WxBindHandle 微信绑定用户账号
func (u *UserHandler) WxBindHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info(": ---------用户绑定微信账户----------------")

	ctx.Log.Info("1:参数验证")
	ctx.Log.Infof("参数为：code:%s, state:%s", ctx.Request.GetString("code"), ctx.Request.GetString("state"))

	if err := ctx.Request.Check("code", "state"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("微信登录过程中有些参数丢失,请正常登录"))
	}

	content, err := u.mem.ValidStateAndGetOpenID(
		ctx.Request.GetString("state"),
		ctx.Request.GetString("code"),
		ctx.Log)

	if err != nil {
		return err
	}

	if err := u.mem.SaveUserOpenID(content, ctx.Request.GetString("state"), ctx.Log); err != nil {
		return err
	}
	return "success"
}
