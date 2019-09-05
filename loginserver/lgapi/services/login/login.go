package login

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/logic"
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

	ctx.Log.Info("2: 判断系统是否被禁用")
	ident := ctx.Request.GetString("ident")
	if err := u.m.CheckSystemStatus(ident); err != nil {
		return err
	}

	ctx.Log.Info("3: 判断用户是否被锁定, 锁定时间过期后要解锁")
	userName := ctx.Request.GetString("username")
	if err := u.m.CheckUserIsLocked(userName); err != nil {
		return err
	}

	ctx.Log.Info("4: 判断用户输入的验证码")
	if err := u.m.CheckWxValidCode(userName, ctx.Request.GetString("wxcode")); err != nil {
		return err
	}

	ctx.Log.Info("5:处理用户账号登录")
	member, err := u.m.Login(userName, ctx.Request.GetString("password"), ident)
	if err != nil {
		return err
	}

	ctx.Log.Info("6:生成返回给子系统的Code")
	result, err := u.m.GenerateCodeAndSysInfo(ident, member.UserID)
	if err != nil {
		return err
	}

	ctx.Log.Info("7: 设置jwt数据")
	ctx.Response.SetJWT(member)

	return result
}
