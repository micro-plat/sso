package function

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/function"
)

type SystemFuncDeleteHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
}


func NewSystemFuncDeleteHandler(container component.IContainer) (u *SystemFuncDeleteHandler) {
	return &SystemFuncDeleteHandler{
		container: container,
		subLib:    sub.NewSystemFunc(container),
	}
}

func (u *SystemFuncDeleteHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------删除系统功能------")
	ctx.Log.Info("1. 参数检查")
	if err := ctx.Request.Check("id"); err !=nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	id := ctx.Request.GetInt("id")
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.Delete(id)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回数据。")
	return"success"
}

