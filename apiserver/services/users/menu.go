package users

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
	"github.com/micro-plat/sso/apiserver/modules/util"
)

//MenuHandler 菜单查询对象
type MenuHandler struct {
	c component.IContainer
	m logic.IMenuLogic
	//p   menu.IPopular
	sys logic.ISystemLogic
}

//NewMenuHandler 创建菜单查询对象
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		c: container,
		m: logic.NewMenuLogic(container),
		//p:   menu.NewPopular(container),
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

	if err := ctx.Request.Check("user_id", "system_id", "ident", "timestamp", "sign"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	d := map[string]interface{}{
		"user_id":   ctx.Request.GetString("user_id"),
		"system_id": ctx.Request.GetString("system_id"),
		"ident":     ctx.Request.GetString("ident"),
		"timestamp": ctx.Request.GetString("timestamp"),
	}

	ctx.Log.Info("请求菜单数据：", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	data, err := u.m.Query(ctx.Request.GetInt64("user_id"), ctx.Request.GetInt("system_id"))
	if err != nil {
		return err
	}

	return data
}

func (u *MenuHandler) getSecret(ident string) (string, error) {
	if ident == "" {
		return "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return "", err
	}

	return data.GetString("secret"), nil
}
