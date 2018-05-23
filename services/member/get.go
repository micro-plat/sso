package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

//GetHandler 获取用户信息
type GetHandler struct {
	c component.IContainer
	m member.ICodeMember
}

//NewGetHandler 创建用户查询操作
func NewGetHandler(container component.IContainer) (u *GetHandler) {
	return &GetHandler{
		c: container,
		m: member.NewCodeMember(container),
	}
}

//Handle 根据登录get获取用户信息，jwt信息获取用户信息
func (u *GetHandler) Handle(ctx *context.Context) (r interface{}) {
	code := ctx.Request.GetString("code")
	if code == "" {
		context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("code不能为空"))
	}

	state, err := u.m.Query(code)
	if err != nil {
		return err
	}
	ctx.Response.SetJWT(state)
	return state
}
