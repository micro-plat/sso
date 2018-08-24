package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	//"github.com/micro-plat/lib4go/security/jwt"
)

//CodeHandler 获取用户信息
type CodeHandler struct {
	c      component.IContainer
	m      member.ICodeMember
	member member.IDBMember
	cache  member.ICacheMember
}

//NewCodeHandler 创建用户查询操作
func NewCodeHandler(container component.IContainer) (u *CodeHandler) {
	return &CodeHandler{
		c:      container,
		m:      member.NewCodeMember(container),
		member: member.NewDBMember(container),
		cache:  member.NewCacheMember(container),
	}
}

//Handle 根据登录get获取用户信息，jwt信息获取用户信息
func (u *CodeHandler) GetHandle(ctx *context.Context) (r interface{}) {
	if err := ctx.Request.Check("code"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("code不能为空"))
	}
	code := ctx.Request.GetString("code")
	state, err := u.m.Query(code)
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

// 切换系统，用旧code换取新code
func (u *CodeHandler) PostHandle(ctx *context.Context) (r interface{}) {
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
	newCode, err := u.m.ExchangeCode(ctx.Request.GetString("code"), (*member.LoginState)(m))
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
