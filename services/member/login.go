package member

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c    component.IContainer
	m    member.IMember
	code member.ICodeMember
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:    container,
		m:    member.NewMember(container),
		code: member.NewCodeMember(container),
	}
}

//Handle 处理用户登录，登录成功后转跳到指定的系统
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {

	//检查输入参数
	if err := ctx.Request.Check("username", "password", "sysid"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	//处理用户登录
	member, err := u.m.Login(ctx.Request.GetString("username"),
		ctx.Request.GetString("password"),
		ctx.Request.GetInt("sysid"))
	if err != nil {
		return err
	}
	url := ctx.Request.GetString("redirect_uri")
	if url == "" {
		url = member.IndexURL
	}

	//保存用户信息
	code, err := u.code.Save(member)
	if err != nil {
		return err
	}
	//设置jwt数据
	ctx.Response.SetJWT(member)
	if strings.Contains(url, "?") {
		ctx.Response.Redirect(301, fmt.Sprintf("%s&code=%s", url, code))
		return
	}
	ctx.Response.Redirect(301, fmt.Sprintf("%s?code=%s", url, code))
	return
}
