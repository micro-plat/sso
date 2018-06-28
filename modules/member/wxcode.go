package member

import (
	"github.com/micro-plat/hydra/component"
)

type IWxcode interface {
	GetWXCode(u string, sysid string) (string, error)
}

type Wxcode struct {
	c component.IContainer
}

const (
	wxCodeCacheFormat = "sso:login:wx-valid-code:{@userName}-{@sysid}"
)

func NewWxcode(c component.IContainer) *Wxcode {
	return &Wxcode{
		c: c,
	}
}

//GetWXCode 发送微信验证码
func (l *Wxcode) GetWXCode(u string, sysid string) (string, error) {
	// cache := l.c.GetRegularCache()
	// key := transform.Translate(wxCodeCacheFormat, "userName", u, "sysid", sysid)
	id := "1000"
	return id, nil
}
