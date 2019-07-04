package system

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
	"github.com/micro-plat/sso/apiserver/modules/util"
)

//InfoHandler 系统信息
type InfoHandler struct {
	c   component.IContainer
	sys logic.ISystemLogic
}

//NewInfoHandler new
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		c:   container,
		sys: logic.NewSystemLogic(container),
	}
}

/*
* Handle: 获取子系统的相关信息
* ident:子系统标识
* timestamp:时间戳
* sign:签名字符
 */
func (u *InfoHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统调用，获取系统信息------")
	if err := ctx.Request.Check("ident", "timestamp", "sign"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	d := map[string]interface{}{
		"ident":     ctx.Request.GetString("ident"),
		"timestamp": ctx.Request.GetString("timestamp"),
	}
	ctx.Log.Info("请求请求系统信息数据：", d)

	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	data, err := u.sys.Get(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	return data
}

func (u *InfoHandler) getSecret(ident string) (string, error) {
	if ident == "" {
		return "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return "", err
	}

	return data.GetString("secret"), nil
}
