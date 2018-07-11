package sysfunc

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/subsystem/sysfunc"

)

type SystemFuncAddHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
}

type SystemFuncAddInput struct {
	Parentid int `form:"parentid"`
	ParentLevel int `form:"parentlevel"`
	Sysid int `form:"sysid" `
	Name string `form:"name" valid:"required"`
	Icon string `form:"icon" valid:"required"`
	Path string `form:"path" valid:"required"`
}

func NewSystemFuncAddHandler(container component.IContainer) (u *SystemFuncAddHandler) {
	return &SystemFuncAddHandler{
		container: container,
		subLib:    sub.NewSystemFunc(container),
	}
}

func (u *SystemFuncAddHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------添加系统功能------")
	ctx.Log.Info("1. 参数检查")
	var inputData SystemFuncAddInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	input := map[string]interface{}{
		"sys_id": inputData.Sysid,
		"name": inputData.Name,
		"icon": inputData.Icon,
		"path": inputData.Path,
		"parentid": inputData.Parentid,
		"level_id": inputData.ParentLevel +1,

	}
	ctx.Log.Info("2.更新数据库数据--------")
	ctx.Log.Info("请求参数：",input)
	err := u.subLib.Add(input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.返回数据。")
	return map[string]interface{}{
		"msg": "success",
	}
}



