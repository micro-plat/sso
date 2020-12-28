package login

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/common/service"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/logic"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	m logic.IMemberLogic
}

//NewLoginHandler 创建登录对象
func NewLoginHandler() (u *LoginHandler) {
	return &LoginHandler{
		m: logic.NewMemberLogic(),
	}
}

//Handle 账号登录
func (u *LoginHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------用户账号登录---------")
	ctx.Log().Info("1:参数验证")
	if err := ctx.Request().Check("username", "password"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, "用户名和密码不能为空")
	}
	member, err := service.Login(ctx.Log(), model.LoginReq{
		UserName: ctx.Request().GetString("username"),
		Password: ctx.Request().GetString("password"),
		Ident:    ctx.Request().GetString("ident"),
		Wxcode:   ctx.Request().GetString("wxcode"),
	})
	if err != nil {
		return err
	}

	ctx.Log().Info("6:生成返回给子系统的Code")
	result, err := u.m.GenerateCodeAndSysInfo(ctx.Request().GetString("ident"), member.UserID)
	if err != nil {
		return err
	}

	ctx.Log().Info("7: 设置jwt数据")
	ctx.User().Auth().Response(member)
	return result
}
