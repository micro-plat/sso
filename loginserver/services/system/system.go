package system

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/loginserver/modules/logic"
)

type SystemHandler struct {
	sys logic.ISystemLogic
}

//NewSystemHandler 系统信息
func NewSystemHandler() (u *SystemHandler) {
	return &SystemHandler{
		sys: logic.NewSystemLogic(),
	}
}

//Handle 取系统信息配置
func (u *SystemHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取系统信息配置---------")

	ctx.Log().Info("1: 获取系统信息配置参数")
	data, err := u.sys.GetSystemConfig(ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log().Info("2: 返回数据")
	return data
}
