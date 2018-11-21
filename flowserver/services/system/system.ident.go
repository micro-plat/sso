package system

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/system"
)

//SystemIdentHandler is
type SystemIdentHandler struct {
	container component.IContainer
	sys       system.ISystem
}

//NewSystemIdentHandler is
func NewSystemIdentHandler(container component.IContainer) (u *SystemIdentHandler) {
	return &SystemIdentHandler{
		container: container,
		sys:       system.NewSystem(container),
	}
}

//GetHandle is
func (u *SystemIdentHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ident := ctx.Request.GetString("ident", "sso")
	data, err := u.sys.Get(ident)
	if err != nil {
		return err
	}
	return data
}

//PostHandle 根据系统名称获取系统的所有用户
func (u *SystemIdentHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("-------根据系统名称获取系统的所有用户------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("system_name"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	ctx.Log.Info("2. 执行操作")
	data, datas, err := u.sys.GetUsers(ctx.Request.GetString("system_name"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"users":   data,
		"alluser": datas,
	}
}
