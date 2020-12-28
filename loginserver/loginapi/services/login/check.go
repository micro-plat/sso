package login

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/access/member"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/logic"
)

//LoginCheckHandler 验证用户是否已登录
type LoginCheckHandler struct {
	m logic.IMemberLogic
}

//NewLoginCheckHandler 创建登录对象
func NewLoginCheckHandler() (u *LoginCheckHandler) {
	return &LoginCheckHandler{
		m: logic.NewMemberLogic(),
	}
}

//Handle 验证用户是否已经登录
func (u *LoginCheckHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------检查已登录用户是否有相应系统的权限---------")

	ctx.Log().Info("1: 获取登录用户信息")
	mem := member.Get(ctx)

	ctx.Log().Info("2:判断系统是否被禁用")
	ident := ctx.Request().GetString("ident")
	if err := u.m.CheckSystemStatus(ident); err != nil {
		return err
	}

	ctx.Log().Info("3:判断当前用户是否有子系统的权限")
	if err := u.m.CheckHasRoles(mem.UserID, ident); err != nil {
		return err
	}

	ctx.Log().Info("4:生成返回给子系统的Code")
	result, err := u.m.GenerateCodeAndSysInfo(ident, mem.UserID)
	if err != nil {
		return err
	}

	ctx.Log().Info("5:返回结果")
	return result
}
