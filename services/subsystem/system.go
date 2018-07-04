
package subsystem

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type SystemHandler struct {
	container component.IContainer
}

func NewSystemHandler(container component.IContainer) (u *SystemHandler) {
	return &SystemHandler{container: container}
}


func (u *SystemHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}





