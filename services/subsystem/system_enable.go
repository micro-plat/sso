package subsystem

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/subsystem"
)

type SystemEnableHandler struct {
	container component.IContainer
	subLib sub.ISystem
}


func NewSystemEnableHandler(container component.IContainer) (u *SystemEnableHandler) {
	return &SystemEnableHandler{
		container: container,
		subLib:   sub.NewSystem(container),
	}
}


func (u *SystemEnableHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------修改系统管理状态------")
	ctx.Log.Info("1. 参数检查")
	id := ctx.Request.GetInt("id")
	status := ctx.Request.GetInt("status")
	dbInput := map[string]interface{}{
		"id":    id,
		"status":  status,
	}
	ctx.Log.Info("2.更新数据库查询--------")
	err := u.subLib.UpdateEnable(dbInput)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.返回数据。")
	return map[string]interface{}{
		"msg": "success",
	}
}










