package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

//CodeHandler 获取用户信息
type CodeHandler struct {
	c component.IContainer
	m member.ICodeMember
}

//NewCodeHandler 创建用户查询操作
func NewCodeHandler(container component.IContainer) (u *CodeHandler) {
	return &CodeHandler{
		c: container,
		m: member.NewCodeMember(container),
	}
}

//Handle 根据登录get获取用户信息，jwt信息获取用户信息
func (u *CodeHandler) Handle(ctx *context.Context) (r interface{}) {
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
