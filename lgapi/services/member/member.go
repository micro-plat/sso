package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	m "github.com/micro-plat/sso/lgapi/modules/access/member"
	"github.com/micro-plat/sso/lgapi/modules/logic"
)

//MemberHandler 用户对象
type MemberHandler struct {
	c   component.IContainer
	sys logic.ISystemLogic
	mem logic.IMemberLogic
}

//NewMemberHandler 创建登录对象
func NewMemberHandler(container component.IContainer) (u *MemberHandler) {
	return &MemberHandler{
		c:   container,
		sys: logic.NewSystemLogic(container),
		mem: logic.NewMemberLogic(container),
	}
}

//ChangePwdHandle 修改用户密码
func (u *MemberHandler) ChangePwdHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------修改用户密码---------")

	if err := ctx.Request.Check("expassword", "newpassword"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	err := u.mem.ChangePwd(int(m.Get(ctx).UserID), ctx.Request.GetString("expassword"), ctx.Request.GetString("newpassword"))
	if err != nil {
		return err
	}

	return "success"
}

//RefreshHandle 刷新token 这个接口只是为了刷新sso登录用户的jwt, jwt刷新在框架就做了
func (u *MemberHandler) RefreshHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------刷新token---------")

	return "success"
}

//GetUserSysHandle 返回用户可以访问的系统
func (u *MemberHandler) GetUserSysHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------返回用户可以访问的子系统---------")

	data, err := u.sys.QueryUserSystem(m.Get(ctx).UserID)
	if err != nil {
		return err
	}
	return data
}
