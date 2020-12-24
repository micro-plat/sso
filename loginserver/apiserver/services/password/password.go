package password

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/password"
)

//PasswordHandler 角色相关功能
type PasswordHandler struct {
	sys password.IForgetPassword
}

//NewPasswordHandler new
func NewPasswordHandler() (u *PasswordHandler) {
	return &PasswordHandler{
		sys: password.NewForgetPassword(),
	}
}

//ChangeHandle: 忘记密码再修改密码
func (u *PasswordHandler) ChangeHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------忘记密码再修改密码------")

	ctx.Log().Info("-------验证参数------")
	if err := ctx.Request().Check("source", "user_name", "password"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("-------获取数据------")
	err := u.sys.ForgetPassword(ctx.Request().GetString("source"), ctx.Request().GetString("user_name"), ctx.Request().GetString("password"))
	if err != nil {
		return err
	}

	ctx.Log().Info("------返回结果------")
	return "success"
}
