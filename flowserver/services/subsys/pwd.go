package subsys

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/flowserver/modules/system"
	"github.com/micro-plat/sso/flowserver/modules/user"
	"github.com/micro-plat/sso/flowserver/modules/util"
)

//PwdHandler is
type PwdHandler struct {
	container component.IContainer
	sys       system.ISystem
	userLib   user.IUser
}

//NewPwdHandler is
func NewPwdHandler(container component.IContainer) (u *PwdHandler) {
	return &PwdHandler{
		container: container,
		sys:       system.NewSystem(container),
		userLib:   user.NewUser(container),
	}
}

//Handle 子系统修改密码
func (u *PwdHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("-------子系统修改密码------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("ident", "user_id", "password", "password_old", "timestamp", "sign"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}
	//获取secret
	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//校验签名
	d := map[string]interface{}{}
	d["ident"] = ctx.Request.GetString("ident")
	d["user_id"] = ctx.Request.GetString("user_id")
	d["password"] = ctx.Request.GetString("password")
	d["password_old"] = ctx.Request.GetString("password_old")
	d["timestamp"] = ctx.Request.GetString("timestamp")
	ctx.Log.Info("请求用户数据：", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}
	ctx.Log.Info("2.执行操作")
	err = u.userLib.ChangePwd(ctx.Request.GetInt("user_id"), ctx.Request.GetString("password_old"), ctx.Request.GetString("password"))
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回结果")
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
