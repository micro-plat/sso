package user

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"
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

	err := u.userLib.ChangePwd(ctx.Request.GetInt("user_id"), ctx.Request.GetString("password_old"), ctx.Request.GetString("password"))
	if err != nil {
		return err
	}

	return "success"
}
