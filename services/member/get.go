package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

type GetHandler struct {
	container component.IContainer
	m         member.IMember
}

func NewGetHandler(container component.IContainer) (u *GetHandler) {
	return &GetHandler{
		c: container,
		m: member.NewMember(container),
	}
}

func (u *GetHandler) Handle(ctx *context.Context) (r interface{}) {
	uid := member.Get(ctx).UserID
	data, err := u.m.Query(uid)
	if err != nil {
		return err
	}
	return data
}
