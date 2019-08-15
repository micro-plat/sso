package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
)

//handling 验证api的参数
func (r *SSO) handling() {
	r.MicroApp.Handling(func(ctx *context.Context) (rt interface{}) {
		if err := ctx.Request.Check("ident", "sign", "timestamp"); err != nil {
			return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
		}

		secret, err := getSecret(ctx.GetContainer(), ctx.Request.GetString("ident"))
		if err != nil {
			return err
		}

		if ok := ctx.Request.CheckSign(secret); !ok {
			return context.NewErrorf(context.ERR_PAYMENT_REQUIRED, "sign签名错误")
		}

		// data := ctx.Request.GetRequestMap("utf8")
		// ctx.Log.Info("请求原数据", data)
		// if _, flag := data["sign"]; !flag {
		// 	return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign is empty")
		// }

		// delete(data, "sign")
		// if ok := util.VerifySign(ctx, data, secret, ctx.Request.GetString("sign")); ok != true {
		// 	return context.NewError(context.ERR_PAYMENT_REQUIRED, "sign签名错误(402)")
		// }
		return nil
	})
}

// getSecret 获取系统的secrect
func getSecret(container context.IContainer, ident string) (string, error) {
	if ident == "" {
		return "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident is empty")
	}
	data, err := logic.NewSystemLogic(container.(component.IContainer)).Get(ident)
	if err != nil {
		return "", err
	}
	return data.GetString("secret"), nil
}
