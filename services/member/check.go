package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

//CheckHandler 检查用户jwt是否有效
type CheckHandler struct {
	container component.IContainer
}

//NewCheckHandler 创建用户检查对象
func NewCheckHandler(container component.IContainer) (u *CheckHandler) {
	return &CheckHandler{container: container}
}

//Handle 检查jwt信息是否有效
func (u *CheckHandler) Handle(ctx *context.Context) (r interface{}) {
	var m member.Member
	if err := ctx.Request.GetJWT(&m); err != nil {
		return context.NewError(context.ERR_FORBIDDEN, err)
	}
	return nil
}
