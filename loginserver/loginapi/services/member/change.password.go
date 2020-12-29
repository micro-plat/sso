package member

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/common/service"
	m "github.com/micro-plat/sso/loginserver/loginapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/logic"
)

//ChangePwdHandler 用户对象
type ChangePwdHandler struct {
	mem logic.IMemberLogic
}

//NewChangePwdHandler 用户
func NewChangePwdHandler() (u *ChangePwdHandler) {
	return &ChangePwdHandler{
		mem: logic.NewMemberLogic(),
	}
}

//Handle 修改用户密码
func (u *ChangePwdHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------修改用户密码---------")

	if err := ctx.Request().Check("expassword", "newpassword"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	err := service.ChangePwd(int(m.Get(ctx).UserID), ctx.Request().GetString("expassword"),
		ctx.Request().GetString("newpassword"))
	if err != nil {
		return err
	}

	ctx.Log().Info("修改成功")
	return "success"
}
