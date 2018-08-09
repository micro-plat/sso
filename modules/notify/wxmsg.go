package notify

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/modules/app"
	"github.com/micro-plat/wechat/mp/message/template"
	"github.com/micro-plat/sso/modules/system"
)

type IWxmsg interface {
	Send(msg *TpMsg) error
}

type Wxmsg struct {
	c         component.IContainer
	db        IDbNotify
	sys       system.ISystem
}

func NewWxmsg(c component.IContainer) *Wxmsg {
	return &Wxmsg{
		c:         c,
		db:        NewDbNotify(c),
		sys:       system.NewSystem(c),
	}
}

//Send 发送微信模板消息
func (l *Wxmsg) Send(msg *TpMsg) error {

	ctx := app.GetWeChatContext(l.c)
	if _, err := template.Send(ctx, &template.TemplateMessage2{
		ToUser:    msg.Openid,
		TemplateId: "_DL41WrU7r6uNYyjD45c5B11ECkOAhwdDG8qqQxbvGs",
		Data: map[string]interface{}{
			"first": map[string]string{
				"value": fmt.Sprintf("[%s]监控告警通知", msg.Name),
			},
			"keyword1": map[string]string{
				"value": msg.Content,
				"color": "#43CD80",
			},
			"keyword2": map[string]string{
				"value": msg.Time,
			},
			"remark": map[string]string{
				"value": "若非本人操作请注意账户安全",
			},
		},
	}); err != nil {
		return fmt.Errorf("发送通知失败:%v(openid: %s)", err, msg.Openid)
	}
	return nil
}
