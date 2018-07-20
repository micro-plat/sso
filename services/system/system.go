package system

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/system"
)

type SystemHandler struct {
	container component.IContainer
	subLib    sub.ISystem
}

func NewSystemHandler(container component.IContainer) (u *SystemHandler) {
	return &SystemHandler{
		container: container,
		subLib:    sub.NewSystem(container),
	}
}

//GetHandle 分页获取系统管理列表
func (u *SystemHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询系统管理数据--------")
	ctx.Log.Info("1.从数据库查询数据--------")
	pi := ctx.Request.GetInt("pi", 1)
	ps := ctx.Request.GetInt("ps", 10)
	name := ctx.Request.GetString("name")
	status := ctx.Request.GetString("status")
	rows, count, err := u.subLib.Query(name, status, pi, ps)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}

//PostHandle 添加系统
func (u *SystemHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------添加系统管理数据------")
	ctx.Log.Info("1. 参数检查")
	var input sub.AddSystemInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.添加数据库查询--------")
	err := u.subLib.Add(&input)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}

//DeleteHandle 删除系统管理ByID
func (u *SystemHandler) DeleteHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------删除系统管理数据-----")
	ctx.Log.Info("1.参数检查")
	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.从数据库删除数据")
	if ctx.Request.GetInt("id") == 1 {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("不能删除当前系统，系统编号：%v", ctx.Request.GetInt("id")))
	}
	err := u.subLib.Delete(ctx.Request.GetInt("id"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}

//PutHandle 更新系统状态
func (u *SystemHandler) PutHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------修改系统管理状态------")
	ctx.Log.Info("1. 参数检查")
	if err := ctx.Request.Check("id", "status"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.更新数据库查询--------")
	err := u.subLib.ChangeStatus(ctx.Request.GetInt("id"), ctx.Request.GetInt("status"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}
