package logic

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrapi/modules/access/member"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

//ICodeMemberLogic xx
type ICodeMemberLogic interface {
	Query(code string) (ls *model.LoginState, err error)
	Save(s *model.LoginState) (string, error) //这个外 api在用
	ExchangeCode(code string, s *model.LoginState) (newCode string, err error)
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

//Query 用户登录
func (c *CodeMemberLogic) Query(code string) (ls *model.LoginState, err error) {
	return c.code.Query(code)
}

// ExchangeCode 删除旧code,生成新code
func (c *CodeMemberLogic) ExchangeCode(code string, s *model.LoginState) (newCode string, err error) {
	return c.code.ExchangeCode(code, s)
}
