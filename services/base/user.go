package base

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/base"
)

type BaseUserHandler struct {
	container component.IContainer
	baseLib   base.IBase
}

func NewBaseUserHandler(container component.IContainer) (u *BaseUserHandler) {
	return &BaseUserHandler{
		container: container,
		baseLib:   base.NewBase(container),
	}
}

func (u *BaseUserHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询用户角色列表--------")
	ctx.Log.Info("1.获取数据")
	rows, err := u.baseLib.QueryUserRoleList()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	fmt.Println("rolelist:", rows)
	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"rolelist": rows,
	}
}
