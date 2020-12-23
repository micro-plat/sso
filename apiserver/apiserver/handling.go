package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
	"net/http"
)

//handling 验证api的参数
func (r *SSO) handling() {
	r.MicroApp.Handling(func(ctx hydra.IContext) (rt interface{}) {
		if err := ctx.Request().Check("ident", "timestamp", "sign"); err != nil {
			return errs.NewError(http.StatusNotAcceptable, err)
		}
		secret, err := getSecret(ctx.Request().GetString("ident"))
		if err != nil {
			return err
		}

		if ok, err := ctx.Request().CheckSign(secret); !ok {
			return errs.NewError(http.StatusPaymentRequired, err)
		}
		return nil
	})
}

// getSecret 获取系统的secrect
func getSecret(ident string) (string, error) {
	data, err := logic.NewSystemLogic().Get(ident)
	if err != nil {
		return "", err
	}
	return data.GetString("secret"), nil
}
