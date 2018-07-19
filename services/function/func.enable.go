package function

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/function"
)

type SystemFuncEnableHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
}

func NewSystemFuncEnableHandler(container component.IContainer) (u *SystemFuncEnableHandler) {
	return &SystemFuncEnableHandler{
		container: container,
		subLib:    sub.NewSystemFunc(container),
	}
}

func (u *SystemFuncEnableHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------修改系统功能状态------")
	ctx.Log.Info("1. 参数检查")
	if err := ctx.Request.Check("id","status");err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.ChangeStatus(ctx.Request.GetInt("id"),ctx.Request.GetInt("status"))
	if err != nil {
		return  err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}
