package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

//GetHandler 获取用户信息
type GetHandler struct {
	c component.IContainer
	m member.IMember
}

//NewGetHandler 创建用户查询操作
func NewGetHandler(container component.IContainer) (u *GetHandler) {
	return &GetHandler{
		c: container,
		m: member.NewMember(container),
	}
}

//Handle 根据登录成功的jwt信息获取用户信息
func (u *GetHandler) Handle(ctx *context.Context) (r interface{}) {
	uid := member.Get(ctx).UserID
	data, err := u.m.Query(uid)
	if err != nil {
		return err
	}
	return data
}
