package role

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/logic"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/sso"
)

//RoleHandler is
type RoleHandler struct {
	container component.IContainer
	roleLib   logic.IRoleLogic
	op        logic.IOperateLogic
}

//NewRoleHandler is
func NewRoleHandler(container component.IContainer) (u *RoleHandler) {
	return &RoleHandler{
		container: container,
		roleLib:   logic.NewRoleLogic(container),
		op:        logic.NewOperateLogic(container),
	}
}

//GetAllHandle 查询角色信息数据
func (u *RoleHandler) GetAllHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询角色信息数据--------")

	ctx.Log.Info("1.参数校验")
	var inputData model.QueryRoleInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.查询数据")
	rows, count, err := u.roleLib.Query(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回数据。")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}

//SaveHandle 编辑/添加角色信息
func (u *RoleHandler) SaveHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑/添加角色信息--------")
	ctx.Log.Info("1.参数校验")
	var inputData model.RoleEditInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	member := sso.GetMember(ctx)
	inputData.BelongID = member.BelongID
	inputData.BelongType = member.BelongType
	if err := u.roleLib.Save(&inputData); err != nil {
		return err
	}

	ctx.Log.Info("3.记录行为")
	if err := u.op.RoleOperate(member, "编辑/添加角色", "role_name", &inputData.RoleName, "role_id", &inputData.RoleID, "status", &inputData.Status, "IsAdd", &inputData.IsAdd); err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果。")
	return "success"
}

//ChangeStatusHandle 修改角色状态
func (u *RoleHandler) ChangeStatusHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------修改角色状态--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("role_id", "status"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.roleLib.ChangeStatus(ctx.Request.GetString("role_id"), ctx.Request.GetInt("status")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.RoleOperate(sso.GetMember(ctx), "修改角色状态", "user_id", ctx.Request.GetString("role_id"), "status", ctx.Request.GetInt("status")); err != nil {
		return err
	}
	ctx.Log.Info("3.返回结果")
	return "success"
}

//DelHandle 删除角色
func (u *RoleHandler) DelHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------删除角色--------")

	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("role_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.roleLib.Delete(ctx.Request.GetInt("role_id")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.记录行为")
	if err := u.op.RoleOperate(sso.GetMember(ctx), "删除角色", "role_id", ctx.Request.GetInt("role_id")); err != nil {
		return err
	}

	ctx.Log.Info("4.返回结果。")
	return "success"
}
