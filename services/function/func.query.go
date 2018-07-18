package function

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/function"
)

type SystemFuncQueryHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
}

func NewSystemFuncQueryHandler(container component.IContainer) (u *SystemFuncQueryHandler) {
	return &SystemFuncQueryHandler{
		container: container,
		subLib:    sub.NewSystemFunc(container),
	}
}

func (u *SystemFuncQueryHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------查询系统功能数据------")
	ctx.Log.Info("1. 参数检查")
	sysid := ctx.Request.GetInt("id")

	ctx.Log.Info("2.丛数据库获取数据--------")
	data,err := u.subLib.Query(sysid)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.返回数据。")
	return data
}
