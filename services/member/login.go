package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/security/jwt"
	"github.com/micro-plat/sso/modules/member"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c     component.IContainer
	login member.ILogin
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:     container,
		login: member.NewLogin(container),
	}
}

//Handle 处理用户登录，登录成功后转跳到指定的系统
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {

	//检查输入参数
	if err := ctx.Request.Form.Check("username", "password", "sysid"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	//处理用户登录
	member, url, err := u.login.Login(ctx.Request.GetString("username"),
		ctx.Request.GetString("password"),
		ctx.Request.GetInt("sysid"))
	if err != nil {
		return err
	}

	//处理jwt参数
	jwtAuth, err := ctx.Request.GetJWTConfig() //获取jwt配置
	if err != nil {
		return err
	}

	//jwt加密
	token, err := jwt.Encrypt(jwtAuth.Secret, jwtAuth.Mode, member, jwtAuth.ExpireAt)
	if err != nil {
		err = fmt.Errorf("jwt.encrypt:%v", err)
		return
	}
	ctx.Response.Redirect(301, fmt.Sprintf("%s?name=%s&token=%s", url, jwtAuth.Name, token))
	return
}
