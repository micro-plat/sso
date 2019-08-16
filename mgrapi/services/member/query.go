package member

/*
import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/access/member"
	"github.com/micro-plat/sso/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

//QueryHandler 查询用户信息
type QueryHandler struct {
	c component.IContainer
	m logic.IMemberLogic
}

func NewQueryHandler(container component.IContainer) (u *QueryHandler) {
	return &QueryHandler{
		c: container,
		m: logic.NewMemberLogic(container),
	}
}

func (u *QueryHandler) Handle(ctx *context.Context) (r interface{}) {
	userName := member.Get(ctx).UserName
	ident := member.Get(ctx).SysIdent
	data, err := u.m.CacheQuery(userName, ident)
	if err != nil {
		return err
	}
	return (*model.LoginState)(data)
}
*/
