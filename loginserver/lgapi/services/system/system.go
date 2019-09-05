package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/logic"
)

type SystemHandler struct {
	c   component.IContainer
	sys logic.ISystemLogic
}

//NewSystemHandler 系统信息
func NewSystemHandler(container component.IContainer) (u *SystemHandler) {
	return &SystemHandler{
		c:   container,
		sys: logic.NewSystemLogic(container),
	}
}

//Handle 取系统信息配置
func (u *SystemHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------获取系统信息配置---------")

	ctx.Log.Info("1: 获取系统信息配置参数")
	data, err := u.sys.GetSystemConfig(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log.Info("2: 返回数据")
	return data
}
