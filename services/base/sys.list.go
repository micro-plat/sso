package base

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/base"
)

type BaseSysHandler struct {
	container component.IContainer
	baseLib   base.IBase
}

func NewBaseSysHandler(container component.IContainer) (u *BaseSysHandler) {
	return &BaseSysHandler{
		container: container,
		baseLib:   base.NewBase(container),
	}
}

func (u *BaseSysHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询系统列表--------")
	ctx.Log.Info("1.获取数据")
	rows, err := u.baseLib.QuerySysList()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"list": rows,
	}
}
