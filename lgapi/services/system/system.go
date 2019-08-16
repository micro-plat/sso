package system

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/lgapi/modules/logic"
)

type SystemHandler struct {
	c   component.IContainer
	sys logic.ISystemLogic
}

//NewSystemHandler 系统信息
func NewSystemHandler(container component.IContainer) (u *SystemHandler) {
	return &SystemHandler{
		c:   container,
		sys: logic.NewSystemLogic(container),
	}
}

//Handle 取系统信息配置
func (u *SystemHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------取系统信息配置---------")

	if err := ctx.Request.Check("ident"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	data, err := u.sys.QuerySysInfoByIdent(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	return map[string]string{
		"ident":    data.GetString("ident"),
		"sys_name": data.GetString("name"),
	}
}
