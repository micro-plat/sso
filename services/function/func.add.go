package function

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/function"
)

type SystemFuncAddHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
}

func NewSystemFuncAddHandler(container component.IContainer) (u *SystemFuncAddHandler) {
	return &SystemFuncAddHandler{
		container: container,
		subLib:    sub.NewSystemFunc(container),
	}
}

func (u *SystemFuncAddHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------添加系统功能------")
	ctx.Log.Info("1. 参数检查")
	var input sub.SystemFuncAddInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.Add(&input)
	if err != nil {
		return  err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}



