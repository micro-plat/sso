package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/subsystem"
)

type SystemEditHandler struct {
	container component.IContainer
	subLib sub.ISystem
}

type SystemEditInput struct {
	Enable	string `form:"enable"`
	Id	string `form:"id"`
	Index_url string `form:"index_url"`
	Login_timeout string `form:"login_timeout"`
	Logo string `form:"logo"`
	Name string `form:"name"`
	Theme string `form:"theme"`
	Layout string `form:"layout"`
}


func NewSystemEditHandler(container component.IContainer) (u *SystemEditHandler) {
	return &SystemEditHandler{
		container: container,
		subLib:   sub.NewSystem(container),
	}
}


func (u *SystemEditHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------编辑系统管理数据------")
	ctx.Log.Info("1. 参数检查")
	var input SystemEditInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	dbInput := map[string]interface{}{
		"enable": input.Enable,
		"id": input.Id,
		"index_url": input.Index_url,
		"login_timeout": input.Login_timeout,
		"logo": input.Logo,
		"name": input.Name,
		"layout": input.Layout,
		"theme": input.Theme,
	}
	ctx.Log.Info("2.更新数据库--------")
	err := u.subLib.UpdateEdit(dbInput)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.返回数据。")
	return "success"
}
