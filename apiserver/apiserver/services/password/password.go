package password

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/password"
)

//PasswordHandler 角色相关功能
type PasswordHandler struct {
	c   component.IContainer
	sys password.IForgetPassword
}

//NewPasswordHandler new
func NewPasswordHandler(container component.IContainer) (u *PasswordHandler) {
	return &PasswordHandler{
		c:   container,
		sys: password.NewForgetPassword(container),
	}
}

//ChangeHandle: 忘记密码再修改密码
func (u *PasswordHandler) ChangeHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------忘记密码再修改密码------")

	ctx.Log.Info("-------验证参数------")
	if err := ctx.Request.Check("source", "source_id", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("-------获取数据------")
	err := u.sys.ForgetPassword(ctx.Request.GetString("source"), ctx.Request.GetString("source_id"), ctx.Request.GetString("password"))
	if err != nil {
		return err
	}

	ctx.Log.Info("------返回结果------")
	return "success"
}
