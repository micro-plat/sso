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

type SystemFuncEditInput struct {
	Id string `form:"id" valid:"required"`
	Name string `form:"name" valid:"required"`
	Icon string `form:"icon" valid:"required"`
	Path string `form:"path" valid:"required"`
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
	var inputData SystemFuncEditInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	input := map[string]interface{}{
		"id": inputData.Id,
		"name": inputData.Name,
		"icon": inputData.Icon,
		"path": inputData.Path,
	}
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.Edit(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.返回数据。")
	return "success"
}


