package login

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/logic"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/sso"
)

//MenuHandler 获取用户菜单信息
type ChangePwdHandler struct {
	container component.IContainer
	subLib    logic.ILoginLogic
}

//NewChangePwdHandler new
func NewChangePwdHandler(container component.IContainer) (u *ChangePwdHandler) {
	return &ChangePwdHandler{
		container: container,
		subLib:    logic.NewLoginLogic(container),
	}
}

//Handle 修改密码
func (u *ChangePwdHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------用户修改密码------")

	ctx.Log.Info("1:参数验证")
	if err := ctx.Request.Check("expassword", "newpassword"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2:修改密码")
	err := u.subLib.ChangePwd(
		sso.GetMember(ctx).UserID,
		ctx.Request.GetString("expassword"),
		ctx.Request.GetString("newpassword"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3:返回结果")
	return "success"
}
