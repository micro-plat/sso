package login

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/logic"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/config"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/sso"
)

//MenuHandler 获取用户菜单信息
type MenuHandler struct {
	container component.IContainer
	subLib    logic.ILoginLogic
}

//NewMenuHandler new
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		container: container,
		subLib:    logic.NewLoginLogic(container),
	}
}

//Handle 用户菜单信息
func (u *MenuHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------获取用户的菜单列表------")

	ctx.Log.Info("1. 获取用户在指定系统的菜单列表数据")
	data, err := u.subLib.QueryUserMenu(sso.GetMember(ctx).UserID, config.Ident)
	if err != nil {
		return err
	}

	ctx.Log.Info("2. 返回菜单数据")
	return data
}
