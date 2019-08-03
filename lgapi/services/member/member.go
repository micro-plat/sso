package member

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"
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
	config := model.GetConf(u.c)
	return map[string]interface{}{
		"requirewxlogin": config.RequireWxLogin,
		"requirecode":    config.RequireCode,
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

//PostHandle sso用户账号登录(微信扫码登录是另一个)
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

//WxConfHandle weixin登录取配置
//pc扫码登录要先生成statecode, 手机微信绑定可以一起就生成
func (u *LoginHandler) WxConfHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 获取weixin配置信息---------")

	config := model.GetConf(u.c)
	return map[string]interface{}{
		"wxlogin_url": config.WxPhoneLoginUrl,
		"appid":       config.Appid,
	}
}

//GetWxStateHandle 生成微信statecode(相当与一个用户标识)
func (u *LoginHandler) GetWxStateHandle(ctx *context.Context) (r interface{}) {
	stateCode := utility.GetGUID()

	ctx.Log.Info("1: 将stateCode存到缓存中,wx会将这个还回,用于判断是否伪造")
	if err := u.m.SaveWxStateCode(stateCode, ""); err != nil {
		return context.NewError(context.ERR_SERVER_ERROR, "系统繁忙，等会在登录")
	}
	return stateCode
}

//WxCheckHandle 验证用户微信登录
func (u *LoginHandler) WxCheckHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 验证手机微信扫码跳转登录---------")

	ctx.Log.Info("1:参数验证")
	ctx.Log.Infof("参数为：code:%s, state:%s", ctx.Request.GetString("code"), ctx.Request.GetString("state"))

	if err := ctx.Request.Check("code", "state"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("微信登录过程中有些参数丢失,请正常登录"))
	}

	content, err := u.m.ValidStateAndGetOpenID(
		ctx.Request.GetString("state"),
		ctx.Request.GetString("code"),
		ctx.Log)
	if err != nil {
		return err
	}
	ctx.Log.Infof("微信返回信息为:%s", content)

	ctx.Log.Info("查询openid是否存在, 不存在表示未绑定")
	err = u.m.ExistsOpenId(content)
	if err != nil {
		return err
	}

	ctx.Log.Info("4:将获取用户openid等信息放到缓存中")
	if err := u.m.SaveWxLoginInfo(ctx.Request.GetString("state"), content); err != nil {
		return context.NewError(context.ERR_NOT_EXTENDED, fmt.Errorf("出现系统错误,等会再登录"))
	}

	return "success"
}

//WxLoginHandle 微信登录，然后跳转  token["openid"].(string)
func (u *LoginHandler) WxLoginHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 用户手机微信扫码跳转登录(5分钟内会一直调)---------")

	ctx.Log.Info("1:参数验证")
	ctx.Log.Infof("参数为：containkey:%d, ident:%s, state:%s",
		ctx.Request.GetInt("containkey"), ctx.Request.GetString("ident"),
		ctx.Request.GetString("state"))

	if err := ctx.Request.Check("state"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("微信登录过程中有些参数丢失,请正常登录"))
	}

	ctx.Log.Info("2:通过statecode看是否返回了openid")
	opID, err := u.m.GetWxLoginInfoByStateCode(ctx.Request.GetString("state"))
	if err != nil {
		return err
	}
	ctx.Log.Infof("openid:%s", opID)
	//用户还没有扫码，要等会
	if opID == "" {
		return "success"
	}

	ctx.Log.Info("3: 通过opid查询是否有相关用户")
	userInfo, err := u.m.GetUserInfoByOpID(opID, ctx.Request.GetString("ident"))
	if err != nil {
		ctx.Log.Errorf("通过openid:%s, 查询用户信息出错: %v+", opID, err)
		return err
	}

	var code string
	if ctx.Request.GetInt("containkey", 1) == 1 {
		ctx.Log.Info("4:设置返回code")
		code, err = u.m.CreateLoginUserCode(userInfo.UserID)
		if err != nil {
			return context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
		}
	}

	ctx.Log.Info("5: 设置jwt数据")
	ctx.Response.SetJWT(userInfo)

	return code
}

//WxValidCodeHandle 获取weixin登录验证码(公众号)
func (u *LoginHandler) WxValidCodeHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------lgapi 用户获取weixin验证码---------")

	if err := ctx.Request.Check("username"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	userName := ctx.Request.GetString("username")
	ident := ctx.Request.GetString("ident")

	ctx.Log.Info("获取要发送验证码的用户名")
	sendUser, err := u.m.GetSendUserByName(userName, ident)
	if err != nil {
		return err
	}

	ctx.Log.Info("调用接口发微信验证码")
	//(userName是用户登录的名字，senduser是发给公众号那边对应的), cachekey是以userName为准, senduser是传给其他接口用的
	err = u.m.SendValidCode(userName, sendUser)
	if err != nil {
		return err
	}
	return "success"
}
