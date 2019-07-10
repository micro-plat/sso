package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/modules/access/member"
	"github.com/micro-plat/sso/apiserver/modules/model"
)

//ICodeMemberLogic xx
type ICodeMemberLogic interface {
	Save(s *model.LoginState) (string, error)
}

//CodeMemberLogic 控制用户登录
type CodeMemberLogic struct {
	c    component.IContainer
	code member.ICodeMember
}

//NewCodeMemberLogic 创建登录对象
func NewCodeMemberLogic(c component.IContainer) *CodeMemberLogic {
	return &CodeMemberLogic{
		c:    c,
		code: member.NewCodeMember(c),
	}
}

//Save 缓存用户信息
func (c *CodeMemberLogic) Save(s *model.LoginState) (code string, err error) {
	return c.code.Save(s)
}
