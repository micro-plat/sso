package function

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	sub "github.com/micro-plat/sso/flowserver/modules/function"
	"github.com/micro-plat/sso/flowserver/modules/member"
	"github.com/micro-plat/sso/flowserver/modules/operate"
)

type SystemFuncAddHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
	op        operate.IOperate
}

func NewSystemFuncAddHandler(container component.IContainer) (u *SystemFuncAddHandler) {
	return &SystemFuncAddHandler{
		container: container,
		subLib:    sub.NewSystemFunc(container),
		op:        operate.NewOperate(container),
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
		return err
	}
	ctx.Log.Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.MenuOperate(
		member.Query(ctx, u.container),
		"添加菜单",
		data,
	); err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}
