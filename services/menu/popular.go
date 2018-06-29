package menu

import (
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/menu"
)

type PopularHandler struct {
	container component.IContainer
	m         menu.IPopular
}

func NewPopularHandler(container component.IContainer) (u *PopularHandler) {
	return &PopularHandler{
		container: container,
		m:         menu.NewPopular(container),
	}
}

//GetHandle 查询常用菜单
func (u *PopularHandler) GetHandle(ctx *context.Context) (r interface{}) {
	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	data, err := u.m.Query(uid, sysid)
	if err != nil {
		return err
	}
	return data
}

//PostHandle 添加常用菜单
func (u *PopularHandler) PostHandle(ctx *context.Context) (r interface{}) {
	if err := ctx.Request.Check("menu_ids", "pids"); err != nil {
		return err
	}
	menuIds := strings.Split(ctx.Request.GetString("menu_ids"), ",")
	pids := strings.Split(ctx.Request.GetString("pids"), ",")
	if len(menuIds) != len(pids) || len(menuIds) > 20 {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "菜单个数错误")
	}

	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	err := u.m.Save(uid, sysid, pids, menuIds)
	if err != nil {
		return err
	}
	return "success"
}
