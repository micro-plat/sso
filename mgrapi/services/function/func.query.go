package function

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/mgrapi/modules/function"
	"github.com/micro-plat/sso/mgrapi/modules/member"
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

// Handle 查询系统功能数据
func (u *SystemFuncQueryHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------查询系统功能数据------")

	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}

	sysid := ctx.Request.GetInt("id")
	data, err := u.subLib.Get(sysid)
	if err != nil {
		return err
	}

	return data
}
