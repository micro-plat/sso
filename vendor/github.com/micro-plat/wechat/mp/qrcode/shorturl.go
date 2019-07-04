package qrcode

import (
	"github.com/micro-plat/wechat/mp"
	"github.com/micro-plat/wechat/mp/base"
)

// ShortURL 将一条长链接转成短链接.
func ShortURL(clt *mp.Context, longURL string) (shortURL string, err error) {
	return base.ShortURL(clt, longURL)
}
