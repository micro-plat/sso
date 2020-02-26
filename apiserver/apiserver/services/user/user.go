package user

import (
	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

//UserHandler 用户信息
type UserHandler struct {
	c    component.IContainer
	user logic.IUserLogic
}

//NewUserHandler new
func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		c:    container,
		user: logic.NewUserLogic(container),
	}
}

//AddHandle 增加用户
func (u *UserHandler) AddHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------增加用户------")

	var req model.UserInputNew
	if err := ctx.Request.Bind(&req); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	if err := u.user.AddUser(req); err != nil {
		return err
	}
	return "success"
}
