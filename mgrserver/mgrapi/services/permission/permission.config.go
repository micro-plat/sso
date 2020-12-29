package permission

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	"github.com/micro-plat/sso/sdk/sso"
)

//DataPermissionHandler 数据权限
type DataPermissionHandler struct {
	subLib logic.IDataPermissionLogic
	op     logic.IOperateLogic
}

//NewDataPermissionHandler new
func NewDataPermissionHandler() (u *DataPermissionHandler) {
	return &DataPermissionHandler{
		subLib: logic.NewDataPermissionLogic(),
		op:     logic.NewOperateLogic(),
	}
}

//GetAllHandle 查询数据权限数据
func (u *DataPermissionHandler) GetAllHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------查询数据权限数据------")

	ctx.Log().Info("1.参数检查")
	if err := ctx.Request().Check("sys_id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.获取数据")
	data, count, err := u.subLib.Query(
		ctx.Request().GetString("sys_id"), ctx.Request().GetString("name"), ctx.Request().GetString("table_name"),
		ctx.Request().GetInt("pi", 1), ctx.Request().GetInt("ps", 10))
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回数据")
	return map[string]interface{}{
		"count": count,
		"list":  data,
	}
}

//AddHandle 添加系统功能
func (u *DataPermissionHandler) AddHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------添加数据权限数据------")

	ctx.Log().Info("1. 参数检查")
	var input model.DataPermissionReq
	if err := ctx.Request().Bind(&input); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.新增数据--------")
	err := u.subLib.Add(&input)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回数据")
	return "success"
}

//EditHandle 编辑系统功能
func (u *DataPermissionHandler) EditHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------编辑系统功能------")

	ctx.Log().Info("1. 参数检查")
	var input model.DataPermissionReq
	if err := ctx.Request().Bind(&input); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.更新数据库数据--------")
	err := u.subLib.Edit(&input)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回数据")
	return "success"
}

//DelHandle 删除系统功能
func (u *DataPermissionHandler) DelHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------删除系统功能------")

	ctx.Log().Info("1. 参数检查")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.更新数据库数据--------")
	err := u.subLib.Delete(ctx.Request().GetInt("id"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	if err := u.op.MenuOperate(sso.GetMember(ctx), "删除数据权限数据", "id", ctx.Request().GetInt("id")); err != nil {
		return err
	}

	ctx.Log().Info("4.返回数据。")
	return "success"
}

//EnableHandle 启用数据权限配置信息
func (u *DataPermissionHandler) EnableHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------启用数据权限配置信息--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.保存数据")
	err := u.subLib.ChangePermissionConfigStatus(ctx.Request().GetString("id"), model.Enable)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回结果。")
	return "success"
}

//DisableHandle 禁用数据权限配置信息
func (u *DataPermissionHandler) DisableHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------禁用数据权限配置信息--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.保存数据")
	err := u.subLib.ChangePermissionConfigStatus(ctx.Request().GetString("id"), model.Disable)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回结果。")
	return "success"
}
