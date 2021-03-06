package system

import (
	"fmt"
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	"github.com/micro-plat/sso/sdk/sso"
)

// SystemHandler 子系统信息
type SystemHandler struct {
	subLib logic.ISystemLogic
	op     logic.IOperateLogic
}

//NewSystemHandler xx
func NewSystemHandler() (u *SystemHandler) {
	return &SystemHandler{
		subLib: logic.NewSystemLogic(),
		op:     logic.NewOperateLogic(),
	}
}

//GetAllHandle 分页获取系统管理列表
func (u *SystemHandler) GetAllHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------查询系统管理数据--------")
	rows, count, err := u.subLib.Query(
		ctx.Request().GetString("name"), ctx.Request().GetString("status"),
		ctx.Request().GetInt("pi", 1), ctx.Request().GetInt("ps", 10))
	if err != nil {
		return errs.NewError(http.StatusNotImplemented, err)
	}

	ctx.Log().Info("2.返回数据")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}

//AddHandle 添加系统
func (u *SystemHandler) AddHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------添加系统管理数据------")

	ctx.Log().Info("1. 参数检查")
	var input model.AddSystemInput
	if err := ctx.Request().Bind(&input); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.添加系统--------")
	err := u.subLib.Add(&input)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.SysOperate(sso.GetMember(ctx), "添加系统", data); err != nil {
		return err
	}

	return "success"
}

//DelHandle 删除系统管理ByID
func (u *SystemHandler) DelHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------删除系统管理数据-----")
	ctx.Log().Info("1.参数检查")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}
	ctx.Log().Info("2.从数据库删除数据")
	if ctx.Request().GetInt("id") == 0 {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("不能删除当前系统，系统编号：%v", ctx.Request().GetInt("id")))
	}
	err := u.subLib.Delete(ctx.Request().GetInt("id"))
	if err != nil {
		return err
	}
	ctx.Log().Info("3.记录行为")
	if err := u.op.SysOperate(sso.GetMember(ctx), "删除系统", "id", ctx.Request().GetInt("id")); err != nil {
		return err
	}

	return "success"
}

//ChangeStatusHandle 更新系统状态
func (u *SystemHandler) ChangeStatusHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------修改系统管理状态------")

	ctx.Log().Info("1. 参数检查")
	if err := ctx.Request().Check("id", "status"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}
	ctx.Log().Info("2.更新系统状态--------")
	err := u.subLib.ChangeStatus(ctx.Request().GetInt("id"), ctx.Request().GetInt("status"))
	if err != nil {
		return err
	}
	ctx.Log().Info("3.记录行为")
	if err := u.op.SysOperate(sso.GetMember(ctx), "修改系统状态", "id", ctx.Request().GetInt("id"), "status", ctx.Request().GetInt("status")); err != nil {
		return err
	}

	return "success"
}

//EditHandle 编辑系统数据
func (u *SystemHandler) EditHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------编辑系统管理数据------")
	ctx.Log().Info("1. 参数检查")

	var input model.SystemEditInput
	if err := ctx.Request().Bind(&input); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}
	ctx.Log().Info("2.更新系统数据--------")
	err := u.subLib.Edit(&input)
	if err != nil {
		return err
	}
	ctx.Log().Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.SysOperate(sso.GetMember(ctx), "编辑系统数据", data); err != nil {
		return err
	}
	return "success"
}

// ExchangeHandle 排序功能菜单
func (u *SystemHandler) ExchangeHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------互换两个功能菜单------")

	ctx.Log().Info("1.参数校验")
	request := ctx.Request()
	if err := request.Check("sys_id", "sortrank", "level_id", "id", "is_up"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.执行操作")
	err := u.subLib.Sort(
		request.GetInt("sys_id"), request.GetInt("sortrank"), request.GetInt("level_id"),
		request.GetInt("id"), request.GetInt("parent"), request.GetInt("is_up") == 2)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	if err := u.op.SysOperate(sso.GetMember(ctx), "菜单移动", "sys_id", request.GetInt("sys_id"), "sortrank", request.GetInt("sortrank"), "level_id", request.GetInt("level_id"), "id", request.GetInt("id")); err != nil {
		return err
	}

	return "success"
}

//ChangeSecretHandle 修改秘钥
func (u *SystemHandler) ChangeSecretHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------修改系统秘钥------")

	ctx.Log().Info("1.参数校验")
	request := ctx.Request()
	if err := request.Check("id", "secret"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.更新秘钥")
	err := u.subLib.ChangeSecret(request.GetInt("id"), request.GetString("secret"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	if err := u.op.SysOperate(sso.GetMember(ctx), "更新秘钥", "sys_id", request.GetInt("id"), "secret", request.GetString("secret")); err != nil {
		return err
	}

	return "success"
}
