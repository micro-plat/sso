package user

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
)

//MenuHandler 菜单查询对象
type MenuHandler struct {
	c   component.IContainer
	m   logic.IMenuLogic
	sys logic.ISystemLogic
}

//NewMenuHandler 创建菜单查询对象
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		c:   container,
		m:   logic.NewMenuLogic(container),
		sys: logic.NewSystemLogic(container),
	}
}

/*
* Handle: 查询用户在某个系统下的菜单数据
* system_id:子系统id标识
* ident:子系统标识
* user_id:用户标识
* timestamp:时间戳
* sign:签名
 */
func (u *MenuHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统调用，查询指定用户在指定系统的菜单列表------")

	if err := ctx.Request.Check("user_id"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	data, err := u.m.Query(ctx.Request.GetInt64("user_id"), ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	return data
}
