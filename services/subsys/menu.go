package subsys

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/menu"
	"github.com/micro-plat/sso/modules/system"
	"github.com/micro-plat/sso/modules/util"
)

//MenuHandler 菜单查询对象
type MenuHandler struct {
	c   component.IContainer
	m   menu.IMenu
	p   menu.IPopular
	sys system.ISystem
}

//NewMenuHandler 创建菜单查询对象
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		c:   container,
		m:   menu.NewMenu(container),
		p:   menu.NewPopular(container),
		sys: system.NewSystem(container),
	}
}

//GetHandle 查询指定用户在指定系统的菜单列表
func (u *MenuHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统调用，查询指定用户在指定系统的菜单列表------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_id", "system_id", "ident", "timestamp", "sign"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	//获取secret
	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//校验签名
	d := map[string]interface{}{}
	d["user_id"] = ctx.Request.GetString("user_id")
	d["system_id"] = ctx.Request.GetString("system_id")
	d["timestamp"] = ctx.Request.GetString("timestamp")
	if ok := util.VerifySign(d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}
	ctx.Log.Info("2. 执行操作")
	data, err := u.m.Query(ctx.Request.GetInt64("user_id"), ctx.Request.GetInt("system_id"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3. 返回数据")
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
