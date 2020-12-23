package system

import (
	"encoding/json"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/mgrserver/modules/logic"
	"github.com/micro-plat/sso/mgrserver/modules/model"
	"github.com/micro-plat/sso/sdk/sso"
)

//SystemMenuHandler 系统菜单信息
type SystemMenuHandler struct {
	subLib logic.ISystemMenuLogic
	op     logic.IOperateLogic
}

//NewSystemMenuHandler new
func NewSystemMenuHandler() (u *SystemMenuHandler) {
	return &SystemMenuHandler{
		subLib: logic.NewSystemMenuLogic(),
		op:     logic.NewOperateLogic(),
	}
}

//ExportHandle 导出菜单数据
func (u *SystemMenuHandler) ExportHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------导出系统菜单数据------")

	ctx.Log().Info("1. 参数检查")
	if err := ctx.Request().Check("id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2:获取数据")
	data, err := u.subLib.Export(ctx.Request().GetInt("id"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	if err := u.op.SysOperate(sso.GetMember(ctx), "导出菜单", "sys_id", ctx.Request().GetInt("id")); err != nil {
		ctx.Log().Errorf("导出菜单->记录日志出错: %+v", err)
	}

	return data
}

//ImportHandle 导入菜单数据
func (u *SystemMenuHandler) ImportHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------导出系统菜单数据------")

	ctx.Log().Info("1. 反序列化json对象")
	var req model.ImportReq
	json.Unmarshal([]byte(ctx.Request().GetString("data")), &req)
	ctx.Log().Infof("导入的菜单数据为: %+v", req)

	ctx.Log().Info("2. 导入菜单数据")
	if err := u.subLib.Import(&req); err != nil {
		return err
	}

	ctx.Log().Info("3.记录行为")
	if err := u.op.SysOperate(sso.GetMember(ctx), "导入菜单", "sys_id", req.Id); err != nil {
		ctx.Log().Errorf("导入菜单->记录日志出错: %+v", err)
	}

	return "success"
}
