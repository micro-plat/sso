//消息设置
package notify

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/notify"
)

type NotifySetHandler struct {
	container component.IContainer
	Lib    notify.INotify
}

func NewNotifySetHandler(container component.IContainer) (u *NotifySetHandler) {
	return &NotifySetHandler{
		container: container,
		Lib:    notify.NewNotify(container),
	}
}
//GetHandle 查询报警消息设置信息
func (u *NotifySetHandler) GetHandle(ctx *context.Context) (r interface{}){
	ctx.Log.Info("--------查询报警消息设置信息------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("user_id","sys_id","pi","ps"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	data,count,err :=u.Lib.Get(
		ctx.Request.GetString("user_id"),
		ctx.Request.GetString("sys_id"),
		ctx.Request.GetInt("pi"),
		ctx.Request.GetInt("ps"),
	)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")

	return map[string]interface{}{
		"list": data,
		"count": count,
	}
}
//PutHandle 添加报警消息设置信息
func (u *NotifySetHandler) PutHandle(ctx *context.Context) (r interface{}){
	ctx.Log.Info("--------添加报警消息设置信息------")
	ctx.Log.Info("1.参数校验")
	var input notify.SettingsInput
	if err := ctx.Request.Bind(&input);err != nil{
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	err := u.Lib.Add(&input)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}

//DeleteHandle 删除消息配置
func (u *NotifySetHandler) DeleteHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------删除报警消息设置信息------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("id");err != nil{
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	err := u.Lib.DeleteSettingsByID(ctx.Request.GetString("id"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}

//PostHandle 编辑报警消息设置信息
func (u *NotifySetHandler) PostHandle(ctx *context.Context) (r interface{}){
	ctx.Log.Info("--------编辑报警消息设置信息------")
	ctx.Log.Info("1.参数校验")
	var input notify.EditSettingsInput
	if err := ctx.Request.Bind(&input);err != nil{
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	err := u.Lib.Edit(&input)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}