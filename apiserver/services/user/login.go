package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/apiserver/modules/logic"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c    component.IContainer
	m    logic.IMemberLogic
	code logic.ICodeMemberLogic
	sys  logic.ISystemLogic
	op   logic.IOperateLogic
}

//NewLoginHandler 用户登录
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:    container,
		m:    logic.NewMemberLogic(container),
		code: logic.NewCodeMemberLogic(container),
		sys:  logic.NewSystemLogic(container),
		op:   logic.NewOperateLogic(container),
	}
}

/*
* Handle: 子系统账号登录
* ident:子系统标识
* username:用户名
* password:用户密码
* timestamp:时间戳
* sign:签名字符
 */
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统用户远程登录---------")

	if err := ctx.Request.Check("username", "password", "ident", "timestamp", "sign"); err != nil {
		return context.NewError(context.ERR_REQUEST_URI_TOO_LONG, err)
	}

	//处理用户登录
	member, err := u.m.Login(ctx.Request.GetString("username"),
		md5.Encrypt(ctx.Request.GetString("password")),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//保存用户信息
	_, err = u.code.Save(member)
	if err != nil {
		return err
	}

	//记录登录行为
	if err := u.op.LoginOperate(member); err != nil {
		return err
	}
	ctx.Log.Infof("%+v", member)
	return member
}
