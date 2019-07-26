package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"
	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/logic"
	"github.com/micro-plat/sso/lgapi/modules/model"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c component.IContainer
	m logic.IMemberLogic
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c: container,
		m: logic.NewMemberLogic(container),
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
	var code = ""
	var err error
	if err = u.m.CheckHasRoles(m.UserID, ctx.Request.GetString("ident")); err != nil {
		ctx.Log.Errorf("验证权限出错: %v", err)
		return err
	}

	if ctx.Request.GetInt("containkey", 1) == 1 {
		ctx.Log.Info("3:已登录返回code")
		code, err = u.m.CreateLoginUserCode(m.UserID)
		if err != nil {
			return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
		}
	}

	ctx.Log.Info("4: 设置jwt数据")
	ctx.Response.SetJWT(m)

	return code
}

//PostHandle sso用户登录
func (u *LoginHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 用户登录---------")

	ctx.Log.Info("1:参数验证")
	if err := ctx.Request.Check("username", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("用户名和密码不能为空"))
	}

	//当有ident时没有权限就跳转错误页面
	ctx.Log.Info("2:处理用户登录")
	member, err := u.m.Login(
		ctx.Request.GetString("username"),
		ctx.Request.GetString("password"),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	var code = ""
	if ctx.Request.GetInt("containkey", 1) == 1 {
		ctx.Log.Info("3: 设置已登录code")
		code, err = u.m.CreateLoginUserCode(member.UserID)
		if err != nil {
			return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
		}
	}
	ctx.Log.Info("4: 设置jwt数据")
	ctx.Response.SetJWT(member)

	return code
}

//RefreshHandle 刷新token 这个接口只是为了刷新sso登录用户的jwt, jwt刷新在框架就做了
func (u *LoginHandler) RefreshHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 刷新token---------")

	return "success"
}

//WxConfHandle weixin登录取配置
func (u *LoginHandler) WxConfHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi weixin登录---------")

	config := model.GetConf(u.c)
	stateCode := utility.GetGUID()

	ctx.Log.Info("1: 将stateCode存到缓存中,wx会将这个还回,用于判断是否伪造")
	if err := u.m.SaveWxLoginStateCode(stateCode); err != nil {
		return context.NewError(context.ERR_SERVER_ERROR, "系统繁忙，等会在登录")
	}

	return map[string]interface{}{
		"wxlogin_url": config.WxLoginUrl,
		"appid":       config.Appid,
		"state":       stateCode,
	}
}

//WxCheckHandle 验证用户微信登录
//这里面有两个code, 一个是wx返回的code, 还有给子系统生成的code
func (u *LoginHandler) WxCheckHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 用户跳转登录---------")

	ctx.Log.Info("1:参数验证")
	ctx.Log.Infof("参数为：containkey:%d, ident:%s, code:%s, state:%s",
		ctx.Request.GetInt("containkey"), ctx.Request.GetString("ident"),
		ctx.Request.GetString("code"), ctx.Request.GetString("state"))

	if err := ctx.Request.Check("code", "state"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("微信登录过程中有些参数丢失,请正常登录"))
	}

	ctx.Log.Info("2:验证state code是否存在, 防止伪造")
	if flag, _ := u.m.ExistsWxLoginStateCode(ctx.Request.GetString("state")); !flag {
		return context.NewError(context.ERR_REQUEST_TIMEOUT, fmt.Errorf("微信登录标识过期,请重新登录"))
	}

	ctx.Log.Info("3:调用wx接口,获取用户openid")
	config := model.GetConf(u.c)
	url := config.WxTokenUrl + "?appid=" + config.Appid + "&secret=" + config.Secret + "&code=" + ctx.Request.GetString("code") + "&grant_type=authorization_code"
	ctx.Log.Infof("获取用户openid的url: %s", url)

	opID, err := u.m.GetWxUserOpID(url)
	if err != nil {
		ctx.Log.Errorf("调用wx api出错: %v+", err)
		return err
	}
	if opID == "" {
		return context.NewError(context.ERR_NOT_EXTENDED, "调用微信失败，稍后再登录")
	}

	ctx.Log.Info("4: 通过opid查询是否有相关用户")
	userInfo, err := u.m.GetUserInfoByOpID(opID, ctx.Request.GetString("ident"))
	if err != nil {
		ctx.Log.Errorf("通过openid:%s, 查询用户信息出错: %v+", opID, err)
		return err
	}

	var code string
	if ctx.Request.GetInt("containkey", 1) == 1 {
		ctx.Log.Info("5:设置返回code")
		code, err = u.m.CreateLoginUserCode(userInfo.UserID)
		if err != nil {
			return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
		}
	}

	ctx.Log.Info("4: 设置jwt数据")
	ctx.Response.SetJWT(userInfo)

	return code
}
