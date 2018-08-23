package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/system"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c      component.IContainer
	m      member.IMember
	code   member.ICodeMember
	wxcode member.IWxcode
	sys    system.ISystem
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:      container,
		m:      member.NewMember(container),
		code:   member.NewCodeMember(container),
		wxcode: member.NewWxcode(container),
		sys:    system.NewSystem(container),
	}
}

//Handle 处理用户登录，登录成功后转跳到指定的系统
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {

	//检查系统是否设置需要微信登录
	b, err := u.isWechatLogin(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	if b == true {
		if err := ctx.Request.Check("wxcode"); err != nil {
			return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
		}
	}
	//检查输入参数
	if err := ctx.Request.Check("username", "password", "ident"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	if b == true {
		if err := u.wxcode.Check(ctx.Request.GetString("username"),
			ctx.Request.GetString("wxcode")); err != nil {
			return err
		}
	}
	//处理用户登录
	member, err := u.m.Login(ctx.Request.GetString("username"),
		md5.Encrypt(ctx.Request.GetString("password")),
		ctx.Request.GetString("ident"))
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
	return map[string]interface{}{
		"code":  code,
		"ident": ctx.Request.GetString("ident"),
	}
}

func (u *LoginHandler) isWechatLogin(ident string) (bool, error) {
	if ident == "" {
		return false, context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return false, err
	}
	if data.GetInt("wechat_status") == 1 {
		return true, nil
	}
	return false, nil
}
