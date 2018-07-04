
package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type RoleHandler struct {
	container component.IContainer
}

func NewRoleHandler(container component.IContainer) (u *RoleHandler) {
	return &RoleHandler{container: container}
}


func (u *RoleHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}





