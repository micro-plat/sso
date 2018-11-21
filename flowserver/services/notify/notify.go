//消息列表
package notify

import (
	"strconv"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/notify"
)

type NotifyHandler struct {
	container component.IContainer
	Lib       notify.INotify
}

func NewNotifyHandler(container component.IContainer) (u *NotifyHandler) {
	return &NotifyHandler{
		container: container,
		Lib:       notify.NewNotify(container),
	}
}

//GetHandle 获取报警消息
func (u *NotifyHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询报警消息数据--------")
	ctx.Log.Info("1.参数校验")
	var input notify.UserNotifyInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "没有权限访问")
	}
	input.UserID = strconv.Itoa(int(l.UserID))
	input.SysID = strconv.Itoa(l.SystemID)
	ctx.Log.Info("2.执行操作")
	data, count, err := u.Lib.Query(&input)
	if err != nil {
		return err
	}
	ctx.Log.Info("2.返回数据")
	return map[string]interface{}{
		"list":  data,
		"count": count,
	}
}

//DeleteHandle 删除消息
func (u *NotifyHandler) DeleteHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------删除报警消息数据--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	err := u.Lib.Delete(ctx.Request.GetInt64("id"), member.Get(ctx).UserID)
	if err != nil {
		return err
	}
	ctx.Log.Info("2.返回数据")
	return "success"
}

//接入系统上传消息  "/sso/notify/info"  [post]
func (u *NotifyHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------上报系统消息---------")
	ctx.Log.Info("1.参数校验")
	var input notify.InsertNotifyInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	if err := u.Lib.Add(&input); err != nil {
		return err
	}
	return "success"
}
