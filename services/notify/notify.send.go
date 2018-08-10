package notify


import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/notify"
)

type NotifySendHandler struct {
	container component.IContainer
	Lib    notify.INotify
}

func NewNotifySendHandler(container component.IContainer) (u *NotifySendHandler) {
	return &NotifySendHandler{
		container: container,
		Lib:    notify.NewNotify(container),
	}
}
//发送消息
func (n *NotifySendHandler) Handle(ctx *context.Context) (r interface{}){
	ctx.Log.Info("-----执行定时任务发送消息------")
	return n.Lib.SendMsg()
}