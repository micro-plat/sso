package member

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/loginserver/sdkapi/modules/logic"
)

//MemberHandler is
type MemberHandler struct {
	m logic.IMemberLogic
}

//NewMemberHandler is
func NewMemberHandler() (u *MemberHandler) {
	return &MemberHandler{
		m: logic.NewMemberLogic(),
	}
}

/*
* GetHandle: 根据用户名查询用户的相关信息
* ident:子系统标识
* username:用户名称
 */
func (u *MemberHandler) GetHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------根据用户名查询用户的相关信息---------")

	if err := ctx.Request().Check("username"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("1. 根据用户名获取用户信息")
	member, err := u.m.QueryUserInfo(ctx.Request().GetString("username"), ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}
	ctx.Log().Info("2. 返回用户信息")
	return member
}
