package user

import (
	"encoding/json"
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	commodel "github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/common/service"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/logic"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/model"
)

//UserHandler 用户信息
type UserHandler struct {
	user logic.IUserLogic
}

//NewUserHandler new
func NewUserHandler() (u *UserHandler) {
	return &UserHandler{
		user: logic.NewUserLogic(),
	}
}

//AddHandle 增加用户
func (u *UserHandler) AddHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------增加用户------")

	var req model.UserInputNew
	if err := ctx.Request().Bind(&req); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	if err := u.user.AddUser(req); err != nil {
		return err
	}
	return "success"
}

//LoginHandle 用户名密码登录
func (u *UserHandler) LoginHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------用户名密码登录------")

	ctx.Log().Info("验证参数")
	if err := ctx.Request().Check("user_name", "password"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("登录及相关验证")
	mem, err := service.Login(ctx.Log(),
		commodel.LoginReq{
			UserName: ctx.Request().GetString("user_name"),
			Password: ctx.Request().GetString("password"),
			Ident:    ctx.Request().GetString("ident"),
		})
	if err != nil {
		return err
	}

	val, _ := json.Marshal(mem)

	ctx.Log().Info("返回数据", string(val))
	return mem
}

//ChangePwdHandle 修改密码
func (u *UserHandler) ChangePwdHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------用户修改密码------")

	ctx.Log().Info("验证参数")
	if err := ctx.Request().Check("user_id", "expassword", "newpassword"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("修改密码")
	err := service.ChangePwd(ctx.Request().GetInt("user_id"),
		ctx.Request().GetString("expassword"),
		ctx.Request().GetString("newpassword"))
	if err != nil {
		return err
	}

	ctx.Log().Info("修改密码完成")
	return "success"
}
