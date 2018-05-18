package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

type CheckHandler struct {
	container component.IContainer
}

func NewCheckHandler(container component.IContainer) (u *CheckHandler) {
	return &CheckHandler{container: container}
}

func (u *CheckHandler) Handle(ctx *context.Context) (r interface{}) {
	var m member.Member
	if err := ctx.Request.GetJWT(&m); err != nil {
		return context.NewError(context.ERR_FORBIDDEN, err)
	}
	return nil
}
