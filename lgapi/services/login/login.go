package login

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/lgapi/modules/logic"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c component.IContainer
	m logic.IMemberLogic
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c: container,
		m: logic.NewMemberLogic(container),
	}
}

//Handle 账号登录
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------用户账号登录---------")

	ctx.Log.Info("1:参数验证")
	if err := ctx.Request.Check("username", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "用户名和密码不能为空")
	}

	ctx.Log.Info("2: 判断用户是否被锁定, 锁定时间过期后要解锁")
	if err := u.m.CheckUserIsLocked(ctx.Request.GetString("username")); err != nil {
		return err
	}

	ctx.Log.Info("3:处理用户账号登录")
	ident := ctx.Request.GetString("ident")
	member, err := u.m.Login(ctx.Request.GetString("username"), ctx.Request.GetString("password"), ident)
	if err != nil {
		return err
	}

	ctx.Log.Info("4:生成返回给子系统的Code")
	result, err := u.m.GenerateCodeAndSysInfo(ident, member.UserID)
	if err != nil {
		return err
	}

	ctx.Log.Info("5: 设置jwt数据")
	ctx.Response.SetJWT(member)

	return result
}
