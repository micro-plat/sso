package login

import (
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/logic"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/config"
)

//LoginHandler 用户登录
type LoginHandler struct {
	container component.IContainer
	subLib    logic.ILoginLogic
}

//NewLoginHandler new
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		container: container,
		subLib:    logic.NewLoginLogic(container),
	}
}

//Handle 用户登录
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------登录验证(用户名密码)---------")

	ctx.Log.Info("1:参数验证")
	if err := ctx.Request.Check("mobile", "password", "verify_code"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "手机号、密码、验证码不能为空")
	}

	ctx.Log.Info("2:判断验证是否正确")
	// if err := u.subLib.CheckVerifyCode(ctx.Request.GetString("mobile"), ctx.Request.GetString("verify_code")); err != nil {
	// 	return err
	// }

	ctx.Log.Info("3: 判断用户是否被锁定, 锁定时间过期后要解锁")
	if err := u.subLib.CheckUserIsLocked(ctx.Request.GetString("mobile")); err != nil {
		return err
	}

	ctx.Log.Info("4:处理用户账号登录")
	member, err := u.subLib.Login(ctx.Request.GetString("mobile"), ctx.Request.GetString("password"), config.Ident)
	if err != nil {
		return err
	}

	ctx.Log.Info("5: 设置jwt数据")
	ctx.Response.SetJWT(member)

	ctx.Log.Info("7: 返回用户数据")
	return map[string]interface{}{
		"user_name":      member.UserName,
		"role_name":      member.RoleName,
		"is_first_login": strings.EqualFold(member.LastLoginTime, ""),
	}
}
