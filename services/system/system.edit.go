package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/operate"
	sub "github.com/micro-plat/sso/modules/system"
)

type SystemEditHandler struct {
	container component.IContainer
	subLib    sub.ISystem
	op        operate.IOperate
}

func NewSystemEditHandler(container component.IContainer) (u *SystemEditHandler) {
	return &SystemEditHandler{
		container: container,
		subLib:    sub.NewSystem(container),
		op:        operate.NewOperate(container),
	}
}

func (u *SystemEditHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------编辑系统管理数据------")
	ctx.Log.Info("1. 参数检查")
	var input sub.SystemEditInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.更新数据库--------")
	err := u.subLib.Edit(&input)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.SysOperate(
		member.Query(ctx, u.container),
		"编辑系统数据",
		data,
	); err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据。")
	return "success"
}
