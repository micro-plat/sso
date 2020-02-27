package user

import (
	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
	commodel "github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/common/service"
)

//UserHandler 用户信息
type UserHandler struct {
	c    component.IContainer
	user logic.IUserLogic
}

//NewUserHandler new
func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		c:    container,
		user: logic.NewUserLogic(container),
	}
}

//AddHandle 增加用户
func (u *UserHandler) AddHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------增加用户------")

	var req model.UserInputNew
	if err := ctx.Request.Bind(&req); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	if err := u.user.AddUser(req); err != nil {
		return err
	}
	return "success"
}

//LoginHandle 用户名密码登录
func (u *UserHandler) LoginHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------用户名密码登录------")

	ctx.Log.Info("验证参数")
	if err := ctx.Request.Check("user_name", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("登录及相关验证")
	mem, err := service.Login(u.c, ctx.Log,
		commodel.LoginReq{
			UserName: ctx.Request.GetString("user_name"),
			Password: ctx.Request.GetString("password"),
			Ident:    ctx.Request.GetString("ident"),
		})
	if err != nil {
		return err
	}

	ctx.Log.Info("返回数据")
	return mem
}

//ChangePwdHandle 修改密码
func (u *UserHandler) ChangePwdHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------用户修改密码------")

	ctx.Log.Info("验证参数")
	if err := ctx.Request.Check("user_id", "expassword", "newpassword"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("修改密码")
	err := service.ChangePwd(u.c,
		ctx.Request.GetInt("user_id"),
		ctx.Request.GetString("expassword"),
		ctx.Request.GetString("newpassword"))
	if err != nil {
		return err
	}

	ctx.Log.Info("修改密码完成")
	return "success"
}