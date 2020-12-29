package login

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/logic"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/model"
)

//LoginCheckHandler 验证用户是否已登录
type getVueConfHandler struct {
	m logic.IMemberLogic
}

//NewGetVueConfHandler 创建登录对象
func NewGetVueConfHandler() (u *getVueConfHandler) {
	return &getVueConfHandler{
		m: logic.NewMemberLogic(),
	}
}

//Handle 验证用户是否已经登录
func (u *getVueConfHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取vue前端页面配置信息---------")
	return model.GetConf()
}
