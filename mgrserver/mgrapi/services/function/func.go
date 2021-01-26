package function

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/const/errorcode"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"

	"github.com/micro-plat/sso/sso"
)

//SystemFuncHandler is
type SystemFuncHandler struct {
	subLib logic.ISystemFuncLogic
	op     logic.IOperateLogic
}

//NewSystemFuncHandler is
func NewSystemFuncHandler() (u *SystemFuncHandler) {
	return &SystemFuncHandler{
		subLib: logic.NewSystemFuncLogic(),
		op:     logic.NewOperateLogic(),
	}
}

//GetHandle 查询系统模块数据
func (u *SystemFuncHandler) FuncGetHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------查询系统模块数据------")

	ctx.Log().Info("1.参数检查")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.获取数据")
	data, err := u.subLib.Get(ctx.Request().GetInt("id"))
	if err != nil {
		return err
	}
	ctx.Log().Info("3.返回数据")
	return data
}

//AddHandle 添加系统功能
func (u *SystemFuncHandler) FuncAddHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------添加系统功能------")

	ctx.Log().Info("1. 参数检查")
	var input model.SystemFuncAddInput
	if err := ctx.Request().Bind(&input); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	/*验证当没有根节点时，不能增加子节点*/
	if input.Parentid == 0 && input.ParentLevel != 0 {
		return errs.NewError(errorcode.ERR_SYSFUNC_ROOTNOTEXISTS, "请先保存根节点")
	}

	ctx.Log().Info("2.更新数据库数据--------")
	err := u.subLib.Add(&input)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.MenuOperate(sso.GetMember(ctx), "添加菜单", data); err != nil {
		return err
	}

	ctx.Log().Info("3.返回数据")
	return "success"
}

//EditHandle 编辑系统功能
func (u *SystemFuncHandler) FuncEditHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------编辑系统功能------")

	ctx.Log().Info("1. 参数检查")
	var input model.SystemFuncEditInput
	if err := ctx.Request().Bind(&input); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.更新数据库数据--------")
	err := u.subLib.Edit(&input)
	if err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	data, _ := types.Struct2Map(&input)
	if err := u.op.MenuOperate(sso.GetMember(ctx), "编辑菜单", data); err != nil {
		return err
	}

	ctx.Log().Info("3.返回数据")
	return "success"
}

//DelHandle 删除系统功能
func (u *SystemFuncHandler) FuncDelHandle(ctx hydra.IContext) (r interface{}) {
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
	if err := u.op.MenuOperate(sso.GetMember(ctx), "删除菜单", "id", ctx.Request().GetInt("id")); err != nil {
		return err
	}

	ctx.Log().Info("3.返回数据。")
	return "success"
}

//ChangeStatusHandle 修改系统功能状态
func (u *SystemFuncHandler) FuncChangeStatusHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------修改系统功能状态------")

	ctx.Log().Info("1. 参数检查")
	if err := ctx.Request().Check("id", "status"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.更新数据库数据--------")
	err := u.subLib.ChangeStatus(ctx.Request().GetInt("id"), ctx.Request().GetInt("status"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	if err := u.op.MenuOperate(sso.GetMember(ctx), "修改菜单状态", "id", ctx.Request().GetInt("id"), "status", ctx.Request().GetInt("status")); err != nil {
		return err
	}

	ctx.Log().Info("3.返回数据")
	return "success"
}
