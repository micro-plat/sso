
package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type LoginHandler struct {
	container component.IContainer
}

func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{container: container}
}


func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}





