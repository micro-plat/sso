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
		TemplateId: "_UC3m5HLVglHeGLkrzP0mU6d_zoaThPC7pn0WVoMEP0",
		Data: map[string]interface{}{
			"first": map[string]string{
				"value": fmt.Sprintf("[%s]监控告警通知", msg.Name),
			},
			"content": map[string]string{
				"value": msg.Content,
				"color": "#FF0000",
			},
			"occurtime": map[string]string{
				"value": msg.Time,
			},
			"remark": map[string]string{
				"value": "请尽快处理",
			},
		},
	}); err != nil {
		return fmt.Errorf("发送模板消息失败:%v(openid: %s)", err, msg.Openid)
	}
	return nil
}
