package sysfunc

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/subsystem/sysfunc"
	"github.com/micro-plat/lib4go/types"
)

type SystemFuncEnableHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
}

type SystemFuncEnableInput struct {
	Id string `form:"id"`
	Status string `form:"status"`
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
	var inputData SystemFuncEnableInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	input, err := types.Struct2Map(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info(input)
	ctx.Log.Info("2.更新数据库数据--------")
	err = u.subLib.Enable(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.返回数据。")
	return map[string]interface{}{
		"msg": "success",
	}
}
