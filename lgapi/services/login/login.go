package login

import (
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/logic"
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

//CheckHandle 验证用户是否已经登录
func (u *LoginHandler) CheckHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------检查已登录用户是否有相应系统的权限---------")

	ctx.Log.Info("1: 获取登录用户信息")
	m := member.Get(ctx)

	ctx.Log.Info("2:判断当前用户是否有子系统的权限")
	ident := ctx.Request.GetString("ident")
	if err := u.m.CheckHasRoles(m.UserID, ident); err != nil {
		return err
	}

	ctx.Log.Info("3:生成返回给子系统的Code")
	result, err := u.generateCode(ident, m.UserID)
	if err != nil {
		return err
	}
	return result
}

//Handle 账号登录
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------sso用户账号登录---------")

	ctx.Log.Info("1:参数验证")
	if err := ctx.Request.Check("username", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "用户名和密码不能为空")
	}

	ctx.Log.Info("2:处理用户账号登录")
	ident := ctx.Request.GetString("ident")
	member, err := u.m.Login(ctx.Request.GetString("username"), ctx.Request.GetString("password"), ident)
	if err != nil {
		return err
	}

	ctx.Log.Info("3:生成返回给子系统的Code")
	result, err := u.generateCode(ident, member.UserID)
	if err != nil {
		return err
	}

	ctx.Log.Info("4: 设置jwt数据")
	ctx.Response.SetJWT(member)

	return result
}

//generateCode 生成登录后的Code
func (u *LoginHandler) generateCode(ident string, userID int64) (map[string]string, error) {
	result := map[string]string{"code": "", "callback": ""}
	if strings.EqualFold(ident, "") {
		return result, nil
	}

	code, err := u.m.CreateLoginUserCode(userID)
	if err != nil {
		return nil, context.NewError(context.ERR_BAD_REQUEST, "请重新登录")
	}
	result["code"] = code
	sysInfo, err := u.sys.QuerySysInfoByIdent(ident)
	if err != nil {
		return nil, err
	}
	if sysInfo != nil {
		result["callback"] = sysInfo.GetString("index_url")
	}
	return result, nil
}
