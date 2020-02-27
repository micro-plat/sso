package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/common/service"
	m "github.com/micro-plat/sso/loginserver/lgapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/logic"
)

//ChangePwdHandler 用户对象
type ChangePwdHandler struct {
	c   component.IContainer
	mem logic.IMemberLogic
}

//NewChangePwdHandler 用户
func NewChangePwdHandler(container component.IContainer) (u *ChangePwdHandler) {
	return &ChangePwdHandler{
		c:   container,
		mem: logic.NewMemberLogic(container),
	}
}

//Handle 修改用户密码
func (u *ChangePwdHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------修改用户密码---------")

	if err := ctx.Request.Check("expassword", "newpassword"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	err := service.ChangePwd(u.c, int(m.Get(ctx).UserID), ctx.Request.GetString("expassword"),
		ctx.Request.GetString("newpassword"))
	if err != nil {
		return err
	}

	ctx.Log.Info("修改成功")
	return "success"
}
