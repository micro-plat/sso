package users

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
	"github.com/micro-plat/sso/apiserver/modules/util"
)

//PwdHandler is
type PwdHandler struct {
	container component.IContainer
	sys       logic.ISystemLogic
	userLib   logic.IUserLogic
}

//NewPwdHandler is
func NewPwdHandler(container component.IContainer) (u *PwdHandler) {
	return &PwdHandler{
		container: container,
		sys:       logic.NewSystemLogic(container),
		userLib:   logic.NewUserLogic(container),
	}
}

/*
* Handle 子系统修改密码
* ident:子系统标识
* user_id:用户标识
* password:新密码
* password_old:老密码
* timestamp:时间戳
* sign:签名
 */
func (u *PwdHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("-------子系统修改密码------")

	if err := ctx.Request.Check("ident", "user_id", "password", "password_old", "timestamp", "sign"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	d := map[string]interface{}{
		"ident":        ctx.Request.GetString("ident"),
		"user_id":      ctx.Request.GetString("user_id"),
		"password":     ctx.Request.GetString("password"),
		"password_old": ctx.Request.GetString("password_old"),
		"timestamp":    ctx.Request.GetString("timestamp"),
	}

	ctx.Log.Info("请求用户数据：", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	err = u.userLib.ChangePwd(ctx.Request.GetInt("user_id"), ctx.Request.GetString("password_old"), ctx.Request.GetString("password"))
	if err != nil {
		return err
	}

	return "success"
}

func (u *PwdHandler) getSecret(ident string) (string, error) {
	if ident == "" {
		return "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return "", err
	}

	return data.GetString("secret"), nil
}
