package user

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/logic"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/model"
	"github.com/micro-plat/sso/sdk/sso"
)

//UserHandler is
type UserHandler struct {
	userLib logic.IUserLogic
	sys     logic.ISystemLogic
	op      logic.IOperateLogic
}

//NewUserHandler is
func NewUserHandler() (u *UserHandler) {
	return &UserHandler{
		userLib: logic.NewUserLogic(),
		sys:     logic.NewSystemLogic(),
		op:      logic.NewOperateLogic(),
	}
}

//GetAllHandle 查询用户信息数据
func (u *UserHandler) GetAllHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------查询用户信息数据--------")

	ctx.Log().Info("1.参数校验")
	var inputData model.QueryUserInput
	if err := ctx.Request().Bind(&inputData); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	rows, count, err := u.userLib.Query(&inputData)
	if err != nil {
		return errs.NewError(http.StatusNotImplemented, err)
	}

	// sql, err := sso.GetDataPermission(
	// 	sso.GetMember(ctx).UserID,
	// 	"pis_inbound_info",
	// 	sso.WithAlias("p"),
	// 	sso.WithCustomParams(map[string]interface{}{
	// 		"seller_id": 10,
	// 	}))

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(sql)

	ctx.Log().Info("2.返回数据。")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}

//ChangeStatusHandle 修改用户状态
func (u *UserHandler) ChangeStatusHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------修改用户状态--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("user_id", "status", "user_name"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.执行操作")
	if err := u.userLib.ChangeStatus(ctx.Request().GetInt("user_id"), ctx.Request().GetInt("status"), ctx.Request().GetString("user_name")); err != nil {
		return errs.NewError(http.StatusNotImplemented, err)
	}

	ctx.Log().Info("3.记录行为")
	if err := u.op.UserOperate(sso.GetMember(ctx), "修改用户状态", "user_id", ctx.Request().GetInt("user_id"), "status", ctx.Request().GetInt("status")); err != nil {
		return err
	}

	ctx.Log().Info("4.返回结果")
	return "success"
}

//DelHandle 删除用户
func (u *UserHandler) DelHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------删除用户--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("user_id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.执行操作")
	if err := u.userLib.Delete(ctx.Request().GetInt("user_id")); err != nil {
		return errs.NewError(http.StatusNotImplemented, err)
	}
	ctx.Log().Info("3.记录行为")
	if err := u.op.UserOperate(sso.GetMember(ctx), "删除用户", "user_id", ctx.Request().GetInt("user_id")); err != nil {
		return err
	}

	ctx.Log().Info("4.返回结果")
	return "success"
}

//EditHandle 编辑用户详细资料（包括系统数据）
func (u *UserHandler) EditHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("------编辑用户详细资料（包括系统数据）--------")

	ctx.Log().Info("1.参数校验")
	var input model.UserInputNew
	if err := ctx.Request().Bind(&input); err != nil && input.UserID == 0 {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("3.执行操作")
	if err := u.userLib.Save(&input); err != nil {
		return err
	}

	ctx.Log().Info("4.返回结果")
	return "success"
}

//AddHandle 添加用户信息
func (u *UserHandler) AddHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------添加用户信息--------")

	ctx.Log().Info("1.参数校验")
	var inputData model.UserInputNew
	if err := ctx.Request().Bind(&inputData); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2.执行操作")
	if err := u.userLib.Add(&inputData); err != nil {
		return err
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

//SetPwdHandle 重置用户密码(1qaz2wsx)
func (u *UserHandler) SetPwdHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------重置用户密码--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("user_id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2: 重置密码")
	if err := u.userLib.SetDefaultPwd(ctx.Request().GetInt("user_id")); err != nil {
		return err
	}
	return "success"
}

//GenerateQrcodeHandle 生成绑定微信信息
func (u *UserHandler) GenerateQrcodeHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------生成绑定微信信息--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().Check("user_id"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2: 生成二维码信息")
	data, err := u.userLib.GenerateQrcodeInfo(ctx.Request().GetInt("user_id"))
	if err != nil {
		return err
	}
	return data
}

//GenerateUserNameHandle 根据汉字生成拼音用户名
func (u *UserHandler) GenerateUserNameHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------根据汉字生成拼音用户名--------")

	ctx.Log().Info("1: 参数验证")
	if err := ctx.Request().Check("full_name"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2: 生成登录用户名并返回")
	return u.userLib.GenerateUserNameByFullName(ctx.Request().GetString("full_name"))
}
