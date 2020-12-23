package system

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
)

//InfoHandler 系统信息
type InfoHandler struct {
	sys logic.ISystemLogic
}

//NewInfoHandler new
func NewInfoHandler() (u *InfoHandler) {
	return &InfoHandler{
		sys: logic.NewSystemLogic(),
	}
}

/*
* Handle: 获取子系统的相关信息
* ident:子系统标识
* timestamp:时间戳
* sign:签名字符
 */
func (u *InfoHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取系统信息------")
	data, err := u.sys.Get(ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}
	return data
}
