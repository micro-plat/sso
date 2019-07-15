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

/*
//GetHandle 处理用户登录，登录成功后转跳到指定的系统
func (u *LoginHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("-------用户登录---------")
	ctx.Log.Info("1.检查参数")
	//检查输入参数
	if err := ctx.Request.Check("username", "password", "ident"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	//处理用户登录
	member, err := u.m.Login(ctx.Request.GetString("username"),
		md5.Encrypt(ctx.Request.GetString("password")),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	//保存用户信息
	code, err := u.code.Save(member)
	if err != nil {
		return err
	}
	//设置jwt数据
	ctx.Response.SetJWT(member)
	//记录登录行为
	if err := u.op.LoginOperate(member); err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return map[string]interface{}{
		"code":  code,
		"ident": ctx.Request.GetString("ident"),
	}
}
*/

/*
//CodeHandle  切换系统，用旧code换取新code
func (u *LoginHandler) CodeHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------换取code-------")
	ctx.Log.Info("1.检查参数")
	if err := ctx.Request.Check("code", "ident", "username"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.检查传入code是否有效")
	//codeMember := u.code.NewCodeMember(u.c)
	loginState, err := u.code.Query(ctx.Request.GetString("code"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3.获取新系统用户数据")
	m, err := u.m.QueryRoleByNameAndIdent(ctx.Request.GetString("username"), loginState.Password, ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	ctx.Log.Info("4.生成新code和新的系统数据")
	newCode, err := u.code.ExchangeCode(ctx.Request.GetString("code"), (*model.LoginState)(m))
	if err != nil {
		return err
	}
	ctx.Log.Info("5.缓存用户数据")
	if err := u.m.SaveLoginStateToCache(m); err != nil {
		return err
	}
	ctx.Log.Info("6.返回数据")
	// 设置jwt数据
	ctx.Response.SetJWT((*model.LoginState)(m))
	return map[string]interface{}{
		"code":  newCode,
		"ident": ctx.Request.GetString("ident"),
	}
}

func (u *LoginHandler) isWechatLogin(ident string) (bool, string, error) {
	if ident == "" {
		return false, "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return false, "", err
	}
	if data.GetInt("wechat_status") == 1 {
		return false, data.GetString("secret"), nil
	}
	return false, data.GetString("secret"), nil
}
*/