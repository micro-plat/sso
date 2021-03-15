package member

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/logic"
)

//MemberGetAllHandler MemberGetAllHandler
type MemberGetAllHandler struct {
	m logic.IMemberLogic
}

//NewMemberGetAllHandler new
func NewMemberGetAllHandler() (u *MemberGetAllHandler) {
	return &MemberGetAllHandler{
		m: logic.NewMemberLogic(),
	}
}

/*
* Handle: 获取所有用户信息
 */
func (u *MemberGetAllHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取所有用户信息---------")
	source := ctx.Request().GetString("source")
	sourceID := ctx.Request().GetString("source_id")
	ident := ctx.Request().GetString("ident")
	ctx.Log().Info("1. 查询数据", ident, source, sourceID)
	members, err := u.m.QueryAllUserInfo(ident, source, sourceID)
	if err != nil {
		return err
	}
	ctx.Log().Info("2. 返回用户信息")
	return members
}
