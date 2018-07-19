package function

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/function"

)

type SystemFuncEditHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
}



func NewSystemFuncEditHandler(container component.IContainer) (u *SystemFuncEditHandler) {
	return &SystemFuncEditHandler{
		container: container,
		subLib:    sub.NewSystemFunc(container),
	}
}

func (u *SystemFuncEditHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------编辑系统功能------")
	ctx.Log.Info("1. 参数检查")
	var input sub.SystemFuncEditInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.Edit(&input)
	if err != nil {
		return  err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}


