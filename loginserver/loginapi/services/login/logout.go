package login

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
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
	ctx.Log().Info("-------sso退出用户登录---------")
	ctx.Log().Info("1. 获取配置")
	srvConf := ctx.APPConf().GetServerConf()
	jwtConf, err := jwt.GetConf(srvConf)
	if err != nil || jwtConf.Disable {
		return err
	}
	ctx.User().Auth().Request(nil)
	return "success"
}
