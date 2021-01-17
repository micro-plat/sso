package login

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/logic"
)

//CheckSignHandler is
type CheckSignHandler struct {
	m logic.ISystemLogic
}

//NewCheckSignHandler is
func NewCheckSignHandler() (u *CheckSignHandler) {
	return &CheckSignHandler{
		m: logic.NewSystemLogic(),
	}
}

func (u *CheckSignHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------检查签名---------")

	if err := ctx.Request().Check("ident", "timestamp", "sign"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	data, err := u.m.Get(ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}

	secret := data.GetString("secret")
	if err :=u.checkSign(ctx, secret); err != nil {
		return errs.NewError(http.StatusPaymentRequired, err)
	}
	return nil
}

// // getSecret 获取系统的secrect
// func getSecret(ident string) (string, error) {
// 	data, err := logic.NewSystemLogic().Get(ident)
// 	if err != nil {
// 		return "", err
// 	}
// 	return data.GetString("secret"), nil
// }

func (u *CheckSignHandler)  checkSign(ctx hydra.IContext, secret string) error {
	keys := ctx.Request().Keys()
	values := net.NewValues()
	var sign string
	for _, key := range keys {
		switch key {
		case "sign":
			sign = ctx.Request().GetString(key)
		default:
			values.Set(key, ctx.Request().GetString(key))
		}
	}
	values.Sort()
	raw := values.Join("", "")
	expect := md5.Encrypt(raw + secret)
	if strings.EqualFold(expect, sign) {
		return nil
	}
	return fmt.Errorf("签名验证失败,expect:%s,sign:%s", expect, sign)
}
