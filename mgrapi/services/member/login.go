package member

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c component.IContainer
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c: container,
	}
}

// UserHandle sso登录后验证用户信息(通过code取登录用户)
func (u *LoginHandler) UserHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------sso登录后去取登录用户---------")

	if err := ctx.Request.Check("code"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("code不能为空"))
	}
	code := ctx.Request.GetString("code")

	ctx.Log.Info("调用sso api 用code取用户信息")
	data, err := model.GetSSOClient(u.c).CheckCodeLogin(code)
	if err != nil {
		return err
	}

	ctx.Log.Infof("data: %v", data)
	ctx.Response.SetJWT(data)

	return data
}
