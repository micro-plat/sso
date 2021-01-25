package login

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/lib4go/errs"
	"net/http"
)

//LogoutHandler 用户退出
type LogoutHandler struct {
}

//NewLogoutHandler 创建用户退出
func NewLogoutHandler() (u *LoginHandler) {
	return &LoginHandler{}
}

//Handle 账号登录
func (u *LogoutHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------用户账号退出---------")

	ctx.Log().Info("1. 获取配置")
	srvConf := ctx.APPConf().GetServerConf()
	jwtConf, err := jwt.GetConf(srvConf)
	if err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}
	ctx.Log().Info("2.构造 token")

	k, v := jwtConf.GetJWTForRspns("expired", true)
	ctx.Response().Header(k, v)
	ctx.Log().Info("3: 完成")
	return "success"
}
