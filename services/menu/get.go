
package menu

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type GetHandler struct {
	container component.IContainer
}

func NewGetHandler(container component.IContainer) (u *GetHandler) {
	return &GetHandler{container: container}
}


func (u *GetHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}





