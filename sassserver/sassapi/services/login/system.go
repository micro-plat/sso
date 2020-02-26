package login

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/logic"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/config"
)

//SystemHandler 系统信息
type SystemHandler struct {
	container component.IContainer
	subLib    logic.ILoginLogic
	op        logic.IOperateLogic
}

//NewSystemHandler new
func NewSystemHandler(container component.IContainer) (u *SystemHandler) {
	return &SystemHandler{
		container: container,
		subLib:    logic.NewLoginLogic(container),
		op:        logic.NewOperateLogic(container),
	}
}

//Handle 用户登录
func (u *SystemHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------获取系统信息----------")

	ctx.Log.Info("1. 执行操作")
	data, err := u.subLib.GetSystemInfo(config.Ident)
	if err != nil {
		return err
	}

	ctx.Log.Info("2. 返回数据")
	return data
	return nil
}
