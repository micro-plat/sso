package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/flowserver/modules/member"
)

//QueryHandler 查询用户信息
type QueryHandler struct {
	c component.IContainer
	m member.ICacheMember
}

func NewQueryHandler(container component.IContainer) (u *QueryHandler) {
	return &QueryHandler{
		c: container,
		m: member.NewCacheMember(container),
	}
}

func (u *QueryHandler) Handle(ctx *context.Context) (r interface{}) {
	userName := member.Get(ctx).UserName
	ident := member.Get(ctx).SysIdent
	data, err := u.m.Query(userName, ident)
	if err != nil {
		return err
	}
	return (*member.LoginState)(data)
}
