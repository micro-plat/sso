package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/flowserver/modules/member"
	"github.com/micro-plat/sso/flowserver/modules/operate"
	"github.com/micro-plat/sso/flowserver/modules/system"
	"github.com/micro-plat/sso/flowserver/modules/util"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c      component.IContainer
	m      member.IMember
	code   member.ICodeMember
	sys    system.ISystem
	op     operate.IOperate
	member member.IDBMember
	cache  member.ICacheMember
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c:      container,
		m:      member.NewMember(container),
		code:   member.NewCodeMember(container),
		sys:    system.NewSystem(container),
		op:     operate.NewOperate(container),
		member: member.NewDBMember(container),
		cache:  member.NewCacheMember(container),
	}
}

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

//SysHandle 子系统远程登录
func (u *LoginHandler) SysHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统用户远程登录---------")
	//检查输入参数
	if err := ctx.Request.Check("username", "password", "ident", "timestamp", "sign"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	//检查系统是否设置需要微信登录
	b, secret, err := u.isWechatLogin(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	if b == true {
		if err := ctx.Request.Check("wxcode"); err != nil {
			return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
		}
	}
	ctx.Log.Info("1.检查参数")

	//校验签名
	d := map[string]interface{}{}
	d["username"] = ctx.Request.GetString("username")
	d["password"] = ctx.Request.GetString("password")
	d["ident"] = ctx.Request.GetString("ident")
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	ctx.Log.Info("2.执行操作")
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
	_, err = u.code.Save(member)
	if err != nil {
		return err
	}

	//记录登录行为
	if err := u.op.LoginOperate(member); err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return member

}

//PostHandle 根据登录get获取用户信息，jwt信息获取用户信息
func (u *LoginHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------根据登录get获取用户信息，jwt信息获取用户信息---------")
	if err := ctx.Request.Check("code"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("code不能为空"))
	}
	code := ctx.Request.GetString("code")
	state, err := u.code.Query(code)
	if err != nil {
		return err
	}
	ctx.Response.SetJWT(state)
	// jwtConf, err := ctx.Request.GetJWTConfig() //获取jwt配置
	// if err != nil {
	// 	return err
	// }
	// jwtToken, err := jwt.Encrypt(jwtConf.Secret, jwtConf.Mode, state, jwtConf.ExpireAt)
	// if err != nil {
	// 	return err
	// }
	// return map[string] interface{}{
	// 	"state":state,
	// 	"jwt":jwtToken,
	// }
	return state
}

//CodeHandle  切换系统，用旧code换取新code
func (u *LoginHandler) CodeHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------换取code-------")
	ctx.Log.Info("1.检查参数")
	if err := ctx.Request.Check("code", "ident", "username"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.检查传入code是否有效")
	codeMember := member.NewCodeMember(u.c)
	loginState, err := codeMember.Query(ctx.Request.GetString("code"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3.获取新系统用户数据")
	m, err := u.member.Query(ctx.Request.GetString("username"), loginState.Password, ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	ctx.Log.Info("4.生成新code和新的系统数据")
	newCode, err := u.code.ExchangeCode(ctx.Request.GetString("code"), (*member.LoginState)(m))
	if err != nil {
		return err
	}
	ctx.Log.Info("5.缓存用户数据")
	if err := u.cache.Save(m); err != nil {
		return err
	}
	ctx.Log.Info("6.返回数据")
	// 设置jwt数据
	ctx.Response.SetJWT((*member.LoginState)(m))
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
