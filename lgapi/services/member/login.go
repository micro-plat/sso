package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/logic"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c  component.IContainer
	m  logic.IMemberLogic
	op logic.IOperateLogic
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:  container,
		m:  logic.NewMemberLogic(container),
		op: logic.NewOperateLogic(container),
	}
}

//CheckHandle 验证用户是否已经登录
func (u *LoginHandler) CheckHandle(ctx *context.Context) (r interface{}) {

	//1: 获取登录用户信息
	m := member.Get(ctx)
	if m == nil {
		return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
	}
	m.Password = ""

	//2:已登录返回key
	code, err := u.m.SetLoginUserCode(m.UserID)
	if err != nil {
		return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
	}

	//3: 设置jwt数据
	ctx.Response.SetJWT(m)

	//4:记录登录行为
	m.SystemID = ctx.Request.GetInt("sysid")
	u.op.LoginOperate(m)

	return code
}

//PostHandle sso用户登录
func (u *LoginHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 用户登录---------")

	//1:参数验证
	if err := ctx.Request.Check("username", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("用户名和密码不能为空"))
	}

	//2:处理用户登录
	member, err := u.m.Login(
		ctx.Request.GetString("username"),
		md5.Encrypt(ctx.Request.GetString("password")))
	if err != nil {
		return err
	}
	member.Password = ""

	//3: 设置已登录key
	code, err := u.m.SetLoginUserCode(member.UserID)
	if err != nil {
		return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
	}

	//3: 设置jwt数据
	ctx.Response.SetJWT(member)

	//4:记录登录行为
	member.SystemID = ctx.Request.GetInt("sysid")
	u.op.LoginOperate(member)

	return code
}

//RefreshHandle 刷新token
func (u *LoginHandler) RefreshHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 刷新token---------")
	return "success"
}
