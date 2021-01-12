package member

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/logic"
)

//MemberSysHandler 用户->系统
type MemberSysHandler struct {
	m logic.IMemberLogic
}

//NewMemberSysHandler xx
func NewMemberSysHandler() (u *MemberSysHandler) {
	return &MemberSysHandler{
		m: logic.NewMemberLogic(),
	}
}

/*
* GetHandle: 查询某个用户有权限的子系统(当前子系统除外)
* user_id:用户标识
 */
func (u *MemberSysHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------查询某个用户有权限的子系统(当前子系统除外)---------")

	ctx.Log().Info("1. 检查参数")
	if err := ctx.Request().Check("user_id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2. 查询用户可用的子系统")
	data, err := u.m.QueryUserSystem(ctx.Request().GetInt("user_id"), ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}
	ctx.Log().Info("3. 返回可用子系统信息")
	return data
}
