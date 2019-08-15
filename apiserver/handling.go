package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
)

//handling 验证api的参数
func (r *SSO) handling() {
	r.MicroApp.Handling(func(ctx *context.Context) (rt interface{}) {
		if err := ctx.Request.Check("ident", "timestamp", "sign"); err != nil {
			return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
		}
		secret, err := getSecret(ctx.GetContainer(), ctx.Request.GetString("ident"))
		if err != nil {
			return err
		}

		if ok, err := ctx.Request.CheckSign(secret); !ok {
			return context.NewError(context.ERR_PAYMENT_REQUIRED, err)
		}
		return nil
	})
}

// getSecret 获取系统的secrect
func getSecret(container context.IContainer, ident string) (string, error) {
	data, err := logic.NewSystemLogic(container.(component.IContainer)).Get(ident)
	if err != nil {
		return "", err
	}
	return data.GetString("secret"), nil
}
