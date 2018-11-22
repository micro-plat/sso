package menu

import (
	"fmt"
	"strings"

	"github.com/micro-plat/sso/flowserver/modules/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/flowserver/modules/menu"
)

//MenuHandler 菜单查询对象
type MenuHandler struct {
	c component.IContainer
	m menu.IMenu
	p menu.IPopular
}

//NewMenuHandler 创建菜单查询对象
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		c: container,
		m: menu.NewMenu(container),
		p: menu.NewPopular(container),
	}
}

//GetHandle 查询指定用户在指定系统的菜单列表
func (u *MenuHandler) GetHandle(ctx *context.Context) (r interface{}) {
	l := member.Query(ctx, u.c)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	data, err := u.m.Query(l.UserID, l.SystemID)
	if err != nil {
		return err
	}
	return data
}

//QueryHandle 查询指定用户在指定系统的菜单列表
func (u *MenuHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统调用，根据系统名称获取系统的所有用户------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_id", "system_id"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	ctx.Log.Info("2. 执行操作")
	data, err := u.m.Query(ctx.Request.GetInt64("user_id"), ctx.Request.GetInt("system_id"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3. 返回数据")
	return data
}

//PostHandle 查询常用菜单
func (u *MenuHandler) PostHandle(ctx *context.Context) (r interface{}) {
	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	data, err := u.m.Query(uid, sysid)
	if err != nil {
		return err
	}
	return data
}

//PutHandle 添加常用菜单
func (u *MenuHandler) PutHandle(ctx *context.Context) (r interface{}) {
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
	err := u.p.Save(uid, sysid, pids, menuIds)
	if err != nil {
		return err
	}
	return "success"
}

//VerifyHandle 查询用户在指定系统的页面是否有权限
func (u *MenuHandler) VerifyHandle(ctx *context.Context) (r interface{}) {
	path := ctx.Request.GetString("path")
	method := ctx.Request.GetString("method", "get")
	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	err := u.m.Verify(uid, sysid, path, method)
	if err != nil {
		return err
	}
	return ""
}
