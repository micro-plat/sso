package subsystem

import (
"github.com/micro-plat/hydra/component"
"github.com/micro-plat/hydra/context"
sub "github.com/micro-plat/sso/modules/subsystem"
)

type SystemQueryHandler struct {
	container component.IContainer
	subLib sub.ISystem
}

type SystemInput struct {
	Name string `form:"name"`
	Status string `form:"status"`
}


func NewSystemQueryHandler(container component.IContainer) (u *SystemQueryHandler) {
	return &SystemQueryHandler{
		container: container,
		subLib:   sub.NewSystem(container),
	}
}


func (u *SystemQueryHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------查询系统管理数据------")
	ctx.Log.Info("1. 参数检查")
	name := ctx.Request.GetString("name")
	status := ctx.Request.GetInt("status")
	dbInput := map[string]interface{}{
		"name":    name,
		"status":  status,
	}
	ctx.Log.Info("2.查询数据库查询--------")
	rows, err := u.subLib.QueryWithField(dbInput)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.返回数据。")
	return map[string]interface{}{
		"list": rows,
	}
}









