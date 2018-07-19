package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/system"
)

type SystemIdentHandler struct {
	container component.IContainer
	sys       system.ISystem
}

func NewSystemIdentHandler(container component.IContainer) (u *SystemIdentHandler) {
	return &SystemIdentHandler{
		container: container,
		sys:       system.NewSystem(container),
	}
}

func (u *SystemIdentHandler) Handle(ctx *context.Context) (r interface{}) {
	ident := ctx.Request.GetString("ident", "sso")
	data, err := u.sys.Get(ident)
	if err != nil {
		return err
	}
	return data
}
