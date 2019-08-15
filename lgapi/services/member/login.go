package member

import (
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/logic"
	"github.com/micro-plat/sso/lgapi/modules/model"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c   component.IContainer
	m   logic.IMemberLogic
	sys logic.ISystemLogic
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:   container,
		m:   logic.NewMemberLogic(container),
		sys: logic.NewSystemLogic(container),
	}
}

//TypeConfHandle 取配置，显示验证码登录还是扫码登录
func (u *LoginHandler) TypeConfHandle(ctx *context.Context) (r interface{}) {
	ident := ctx.Request.GetString("ident")
	sysName := ""
	if !strings.EqualFold(ident, "") {
		data, err := u.sys.QuerySysInfoByIdent(ident)
		if err == nil && data != nil {
			sysName = data.GetString("name")
		}
	}
	return map[string]interface{}{
		"sysname": sysName,
	}
}

//CheckHandle 验证用户是否已经登录
func (u *LoginHandler) CheckHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 用户跳转登录---------")

	ctx.Log.Info("1: 获取登录用户信息")
	m := member.Get(ctx)
	if m == nil {
		return context.NewError(context.ERR_FORBIDDEN, "请重新登录")
	}
	ctx.Log.Infof("用户信息:%v", m)

	ctx.Log.Info("2:判断当前用户是否有这个子系统的权限")

	ident := ctx.Request.GetString("ident")
	var err error
	if err = u.m.CheckHasRoles(m.UserID, ident); err != nil {
		ctx.Log.Errorf("验证权限出错: %v", err)
		return err
	}
	result := map[string]string{
		"code":     "",
		"callback": "",
	}

	//是否直接调转回子系统
	if ident != "" {
		ctx.Log.Info("3:已登录返回code")
		code, err := u.m.CreateLoginUserCode(m.UserID)
		if err != nil {
			return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
		}
		result["code"] = code
		sysInfo, err := u.sys.QuerySysInfoByIdent(ident)
		if err != nil {
			ctx.Log.Errorf("查询系统信息出错: %v+", err)
		}
		if err == nil && sysInfo != nil && sysInfo.GetString("index_url") != "" {
			result["callback"] = sysInfo.GetString("index_url")
		}
	}

	ctx.Log.Info("4: 设置jwt数据")
	ctx.Response.SetJWT(m)

	return result
}

//PostHandle sso用户账号登录
func (u *LoginHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 用户登录---------")

	ctx.Log.Info("1:参数验证")
	if err := ctx.Request.Check("username", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "用户名和密码不能为空")
	}

	//验证码这个是不是可以不用密码验证了
	config := model.GetConf(u.c)
	if config.RequireCode {
		validatecode := ctx.Request.GetString("validatecode")
		if strings.EqualFold(validatecode, "") {
			return context.NewError(context.ERR_BAD_REQUEST, "请输入微信验证码")
		}
		//验证通过公众号发的验证码
		isValid, err := u.m.ValidVerifyCode(ctx.Request.GetString("username"), validatecode)
		if err != nil {
			return err
		}
		if !isValid {
			return context.NewError(context.ERR_BAD_REQUEST, "微信验证码错误")
		}
	}

	ident := ctx.Request.GetString("ident")
	//当有ident时没有权限就跳转错误页面
	ctx.Log.Info("2:处理用户账号登录")
	member, err := u.m.Login(
		ctx.Request.GetString("username"),
		ctx.Request.GetString("password"),
		ident)
	if err != nil {
		return err
	}

	result := map[string]string{
		"code":     "",
		"callback": "",
	}

	if ctx.Request.GetString("ident") != "" {
		ctx.Log.Info("3: 设置已登录code")
		code, err := u.m.CreateLoginUserCode(member.UserID)
		if err != nil {
			return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
		}
		result["code"] = code
		sysInfo, err := u.sys.QuerySysInfoByIdent(ident)
		if err != nil {
			ctx.Log.Errorf("查询系统信息出错: %v+", err)
		}
		if err == nil && sysInfo != nil && sysInfo.GetString("index_url") != "" {
			result["callback"] = sysInfo.GetString("index_url")
		}
	}
	ctx.Log.Info("4: 设置jwt数据")
	ctx.Response.SetJWT(member)

	return result
}

//RefreshHandle 刷新token 这个接口只是为了刷新sso登录用户的jwt, jwt刷新在框架就做了
func (u *LoginHandler) RefreshHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 刷新token---------")

	return "success"
}
