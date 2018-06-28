package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/app"
)

//QRCodeHandler 获取二维码信息
type QRCodeHandler struct {
	c component.IContainer
}

//NewQRCodeHandler 获取二维码信息
func NewQRCodeHandler(container component.IContainer) (u *QRCodeHandler) {
	return &QRCodeHandler{
		c: container,
	}
}

//Handle 返回当前用户的
func (u *QRCodeHandler) Handle(ctx *context.Context) (r interface{}) {
	if err := ctx.Request.Check("sysid"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	url := app.GetConf(u.c).CheckURL
	uuid := ctx.Request.GetUUID()
	sysid := ctx.Request.GetString("sysid")
	return map[string]interface{}{
		"url": fmt.Sprintf("%s?code=%s&sysid=%s", url, uuid, sysid),
	}
}
