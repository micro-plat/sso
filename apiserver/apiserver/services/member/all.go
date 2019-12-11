package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
)

//MemberGetAllHandler MemberGetAllHandler
type MemberGetAllHandler struct {
	container component.IContainer
	m         logic.IMemberLogic
}

//NewMemberGetAllHandler new
func NewMemberGetAllHandler(container component.IContainer) (u *MemberGetAllHandler) {
	return &MemberGetAllHandler{
		container: container,
		m:         logic.NewMemberLogic(container),
	}
}

/*
* Handle: 获取所有用户信息
 */
func (u *MemberGetAllHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------获取所有用户信息---------")

	ctx.Log.Info("1. 查询数据")
	members, err := u.m.QueryAllUserInfo()
	if err != nil {
		return err
	}
	ctx.Log.Info("2. 返回用户信息")
	return members
}
