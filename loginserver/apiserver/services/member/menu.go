package member

import (
	"fmt"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/logic"
)

//MenuHandler 菜单查询对象
type MenuHandler struct {
	m logic.IMenuLogic
}

//NewMenuHandler 创建菜单查询对象
func NewMenuHandler() (u *MenuHandler) {
	return &MenuHandler{
		m: logic.NewMenuLogic(),
	}
}

/*
* Handle: 查询用户在某个系统下的菜单数据
* ident:子系统标识
* user_id:用户标识
 */
func (u *MenuHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取用户的菜单列表------")

	if err := ctx.Request().Check("user_id"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	ctx.Log().Info("1. 获取用户在指定系统的菜单列表数据")
	data, err := u.m.Query(ctx.Request().GetInt64("user_id"), ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 返回菜单数据")
	return data
}
