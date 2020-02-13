package permission

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	"github.com/micro-plat/sso/sdk/sso"
)

//DataPermissionHandler 数据权限
type DataPermissionHandler struct {
	container component.IContainer
	subLib    logic.IDataPermissionLogic
	op        logic.IOperateLogic
}

//NewDataPermissionHandler new
func NewDataPermissionHandler(container component.IContainer) (u *DataPermissionHandler) {
	return &DataPermissionHandler{
		container: container,
		subLib:    logic.NewDataPermissionLogic(container),
		op:        logic.NewOperateLogic(container),
	}
}

//GetAllHandle 查询数据权限数据
func (u *DataPermissionHandler) GetAllHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------查询数据权限数据------")

	ctx.Log.Info("1.参数检查")
	if err := ctx.Request.Check("sys_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.获取数据")
	data, count, err := u.subLib.Query(
		ctx.Request.GetString("sys_id"), ctx.Request.GetString("table_name"),
		ctx.Request.GetInt("pi", 1), ctx.Request.GetInt("ps", 10))
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return map[string]interface{}{
		"count": count,
		"list":  data,
	}
}

//AddHandle 添加系统功能
func (u *DataPermissionHandler) AddHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------添加数据权限数据------")

	ctx.Log.Info("1. 参数检查")
	var input model.DataPermissionReq
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.新增数据--------")
	err := u.subLib.Add(&input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.MenuOperate(sso.GetMember(ctx), "添加数据权限数据", data); err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return "success"
}

//EditHandle 编辑系统功能
func (u *DataPermissionHandler) EditHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------编辑系统功能------")

	ctx.Log.Info("1. 参数检查")
	var input model.DataPermissionReq
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
	if err := u.op.MenuOperate(sso.GetMember(ctx), "编辑数据权限数据", data); err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return "success"
}

//DelHandle 删除系统功能
func (u *DataPermissionHandler) DelHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------删除系统功能------")

	ctx.Log.Info("1. 参数检查")
	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.更新数据库数据--------")
	err := u.subLib.Delete(ctx.Request.GetInt("id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3.记录行为")
	if err := u.op.MenuOperate(sso.GetMember(ctx), "删除数据权限数据", "id", ctx.Request.GetInt("id")); err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据。")
	return "success"
}
