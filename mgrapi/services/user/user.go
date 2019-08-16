package user

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/access/member"
	"github.com/micro-plat/sso/mgrapi/modules/logic"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

//UserHandler is
type UserHandler struct {
	container component.IContainer
	userLib   logic.IUserLogic
	sys       logic.ISystemLogic
	op        logic.IOperateLogic
	member    logic.IMemberLogic
}

//NewUserHandler is
func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		container: container,
		userLib:   logic.NewUserLogic(container),
		sys:       logic.NewSystemLogic(container),
		op:        logic.NewOperateLogic(container),
		member:    logic.NewMemberLogic(container),
	}
}

//GetHandle 获取用户所拥有的系统列表
func (u *UserHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------查询用户所有系统列表---------")
	userID := member.Get(ctx).UserID
	datas, err := u.sys.GetAll(userID)
	if err != nil {
		return err
	}
	ctx.Log.Info("返回数据")
	return datas
}

//PostHandle 查询用户信息数据
func (u *UserHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询用户信息数据--------")
	ctx.Log.Info("1.参数校验")
	var inputData model.QueryUserInput
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	rows, count, err := u.userLib.Query(&inputData)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"count": count,
		"list":  rows,
	}
}

//PutHandle 修改用户状态
func (u *UserHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------修改用户状态--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_id", "status"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.userLib.ChangeStatus(ctx.Request.GetInt("user_id"), ctx.Request.GetInt("status")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.记录行为")
	if err := u.op.UserOperate(member.Query(ctx, u.container), "修改用户状态", "user_id", ctx.Request.GetInt("user_id"), "status", ctx.Request.GetInt("status")); err != nil {
		return err
	}

	ctx.Log.Info("4.返回结果")
	return "success"
}

//DeleteHandle 删除用户
func (u *UserHandler) DeleteHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------删除用户--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	if err := u.userLib.Delete(ctx.Request.GetInt("user_id")); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.记录行为")
	if err := u.op.UserOperate(member.Query(ctx, u.container), "删除用户", "user_id", ctx.Request.GetInt("user_id")); err != nil {
		return err
	}

	ctx.Log.Info("4.返回结果")
	return "success"
}

//EditHandle 编辑用户详细资料（包括系统数据）
func (u *UserHandler) EditHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("------编辑用户详细资料（包括系统数据）--------")
	ctx.Log.Info("1.参数校验")
	// l := member.Query(ctx, u.container)
	// if l == nil {
	// 	return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	// }

	var input model.UserInputNew
	if err := ctx.Request.Bind(&input); err != nil && input.UserID == 0 {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("3.执行操作")
	if err := u.userLib.Save(&input); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("4.返回结果")
	return "success"
}

//SaveHandle 添加用户信息
func (u *UserHandler) SaveHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加用户信息--------")
	ctx.Log.Info("1.参数校验")
	// l := member.Query(ctx, u.container)
	// if l == nil {
	// 	return context.NewError(context.ERR_FORBIDDEN, "code not be null")
	// }
	ctx.Log.Info("2.参数校验")
	var inputData model.UserInputNew
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("3.执行操作")
	if err := u.userLib.Add(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("4.返回结果")
	return "success"
}

//GetAllHandle 查询当前系统下用户列表
// func (u *UserHandler) GetAllHandle(ctx *context.Context) (r interface{}) {

// 	ctx.Log.Info("--------查询当前系统下用户列表--------")
// 	ctx.Log.Info("1.参数校验")
// 	l := member.Query(ctx, u.container)
// 	if l == nil {
// 		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
// 	}
// 	if err := ctx.Request.Check("pi", "ps"); err != nil {
// 		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
// 	}
// 	rows, count, err := u.userLib.GetAll(l.SystemID, ctx.Request.GetInt("pi"), ctx.Request.GetInt("ps"))
// 	if err != nil {
// 		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
// 	}

// 	ctx.Log.Info("2.返回数据。")
// 	return map[string]interface{}{
// 		"count": count,
// 		"list":  rows,
// 	}
// }

//InfoHandle 查询用户信息
// func (u *UserHandler) InfoHandle(ctx *context.Context) (r interface{}) {
// 	ctx.Log.Info("--------查询用户信息--------")
// 	ctx.Log.Info("1.参数校验")
// 	l := member.Query(ctx, u.container)
// 	if l == nil {
// 		return context.NewError(context.ERR_FORBIDDEN, "code not be null")
// 	}
// 	var uid int64
// 	err := ctx.Request.Check("user_id")
// 	if err != nil {
// 		uid = member.Get(ctx).UserID
// 	} else {
// 		uid = ctx.Request.GetInt64("user_id")
// 	}
// 	ctx.Log.Info("2.验证权限")
// 	if err = u.member.QueryAuth(int64(l.SystemID), uid); err != nil {
// 		return context.NewError(context.ERR_FORBIDDEN, err)
// 	}

// 	ctx.Log.Info("3.执行操作")
// 	data, err := u.userLib.Get(int(uid))
// 	if err != nil {
// 		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
// 	}

// 	ctx.Log.Info("4.返回结果")
// 	return map[string]interface{}{
// 		"userinfo": data,
// 	}
// }
