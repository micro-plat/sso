package function

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	sub "github.com/micro-plat/sso/flowserver/modules/function"
	"github.com/micro-plat/sso/flowserver/modules/member"
	"github.com/micro-plat/sso/flowserver/modules/operate"
)

//SystemFuncHandler is
type SystemFuncHandler struct {
	container component.IContainer
	subLib    sub.ISystemFunc
	op        operate.IOperate
}

//NewSystemFuncHandler is
func NewSystemFuncHandler(container component.IContainer) (u *SystemFuncHandler) {
	return &SystemFuncHandler{
		container: container,
		subLib:    sub.NewSystemFunc(container),
		op:        operate.NewOperate(container),
	}
}

//GetHandle 查询系统功能数据
func (u *SystemFuncHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------查询系统功能数据------")
	ctx.Log.Info("1. 参数检查")
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	}
	sysid := ctx.Request.GetInt("id")
	ctx.Log.Info("2.丛数据库获取数据")
	data, err := u.subLib.Get(sysid)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return data
}

//PostHandle 添加系统功能
func (u *SystemFuncHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------添加系统功能------")
	ctx.Log.Info("1. 参数检查")
	var input sub.SystemFuncAddInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.Add(&input)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.MenuOperate(
		member.Query(ctx, u.container),
		"添加菜单",
		data,
	); err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}

//PutHandle 编辑系统功能
func (u *SystemFuncHandler) PutHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------编辑系统功能------")
	ctx.Log.Info("1. 参数检查")
	var input sub.SystemFuncEditInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.Edit(&input)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.MenuOperate(
		member.Query(ctx, u.container),
		"编辑菜单",
		data,
	); err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}

//DeleteHandle 删除系统功能
func (u *SystemFuncHandler) DeleteHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------删除系统功能------")
	ctx.Log.Info("1. 参数检查")
	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	id := ctx.Request.GetInt("id")
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.Delete(id)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.MenuOperate(
		member.Query(ctx, u.container),
		"删除菜单",
		"id", ctx.Request.GetInt("id"),
	); err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据。")
	return "success"
}

//EnableHandle 修改系统功能状态
func (u *SystemFuncHandler) EnableHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------修改系统功能状态------")
	ctx.Log.Info("1. 参数检查")
	if err := ctx.Request.Check("id", "status"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.ChangeStatus(ctx.Request.GetInt("id"), ctx.Request.GetInt("status"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.MenuOperate(
		member.Query(ctx, u.container),
		"修改菜单状态",
		"id", ctx.Request.GetInt("id"), "status", ctx.Request.GetInt("status"),
	); err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}
