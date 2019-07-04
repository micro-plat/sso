package notify

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/notify"
)

//NotifySendHandler is
type NotifySendHandler struct {
	container component.IContainer
	Lib       notify.INotify
}

//NewNotifySendHandler is
func NewNotifySendHandler(container component.IContainer) (u *NotifySendHandler) {
	return &NotifySendHandler{
		container: container,
		Lib:       notify.NewNotify(container),
	}
}

//Handle 发送消息
func (n *NotifySendHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-----执行定时任务发送消息------")
	err := n.Lib.SendMsg()
	if e := err.(context.IError); e.GetCode() == 204 {
		ctx.Response.SetStatus(204)
		return
	}
	return err
}
