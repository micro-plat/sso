//消息设置
package notify

import (
	"fmt"
	"strconv"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/member"
	"github.com/micro-plat/sso/mgrapi/modules/notify"
)

var keywords = []string{"数据库", "网络", "参数"}

type NotifySetHandler struct {
	container component.IContainer
	Lib       notify.INotify
}

func NewNotifySetHandler(container component.IContainer) (u *NotifySetHandler) {
	return &NotifySetHandler{
		container: container,
		Lib:       notify.NewNotify(container),
	}
}

func isKeywords(f string) bool {
	for _, i := range keywords {
		if f == i {
			return true
		}
	}
	return false
}

//GetHandle 查询报警消息设置信息
func (u *NotifySetHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------查询报警消息设置信息------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("pi", "ps"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "没有权限访问")
	}
	userID := l.UserID
	sysID := int64(l.SystemID)
	ctx.Log.Info("2.执行操作")
	data, count, err := u.Lib.Get(
		userID,
		sysID,
		ctx.Request.GetInt64("pi"),
		ctx.Request.GetInt64("ps"),
	)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")

	return map[string]interface{}{
		"list":  data,
		"count": count,
	}
}

//PutHandle 添加报警消息设置信息
func (u *NotifySetHandler) PutHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------添加报警消息设置信息------")
	ctx.Log.Info("1.参数校验")
	var input notify.SettingsInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	l := member.Query(ctx, u.container)
	if l == nil {
		return context.NewError(context.ERR_FORBIDDEN, "没有权限访问")
	}
	input.UserID = strconv.Itoa(int(l.UserID))
	input.SysID = strconv.Itoa(l.SystemID)
	if !isKeywords(input.Keywords) {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("不是有效的关键字：%v", input.Keywords))
	}
	ctx.Log.Info("2.执行操作")
	err := u.Lib.AddSettings(&input)
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
	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	err := u.Lib.DeleteSettings(ctx.Request.GetInt64("id"), member.Get(ctx).UserID)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}

//PostHandle 编辑报警消息设置信息
func (u *NotifySetHandler) PostHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------编辑报警消息设置信息------")
	ctx.Log.Info("1.参数校验")
	var input notify.EditSettingsInput
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	err := u.Lib.EditSettings(&input)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回数据")
	return "success"
}
