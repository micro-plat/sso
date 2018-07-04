
package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type UserHandler struct {
	container component.IContainer
}

func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{container: container}
}


func (u *UserHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}





