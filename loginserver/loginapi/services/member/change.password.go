package member

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	m "github.com/micro-plat/sso/loginserver/loginapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/logic"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/login"
)

//ChangePwdHandler 用户对象
type ChangePwdHandler struct {
	mem logic.IMemberLogic
	l   *login.LoginLogic
}

//NewChangePwdHandler 用户
func NewChangePwdHandler() (u *ChangePwdHandler) {
	return &ChangePwdHandler{
		mem: logic.NewMemberLogic(),
		l:   login.NewLoginLogic(),
	}
}

//Handle 修改用户密码
func (u *ChangePwdHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------修改用户密码---------")

	ctx.Log().Info("1.检查必须参数")
	if err := ctx.Request().Check("expassword", "newpassword"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}
	userInfo := m.Get(ctx)
	ctx.Log().Info("2.执行密码修改：", userInfo.UserID, userInfo.UserName)

	err := u.l.ChangePwd(int(userInfo.UserID), ctx.Request().GetString("expassword"),
		ctx.Request().GetString("newpassword"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3.修改成功")
	return "success"
}
