package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/logic"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/model"
)

//BindWxHandler 绑定微信
type BindWxHandler struct {
	c   component.IContainer
	mem logic.IMemberLogic
}

//NewBindWxHandler 绑定微信
func NewBindWxHandler(container component.IContainer) (u *BindWxHandler) {
	return &BindWxHandler{
		c:   container,
		mem: logic.NewMemberLogic(container),
	}
}

//CheckHandle 验证用户信息
func (u *BindWxHandler) CheckHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------验证要绑定的用户信息---------")

	ctx.Log.Info("1: 验证参数")
	if err := ctx.Request.Check("user_id", "sign", "timestamp"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2: 验证用户信息及签名")
	if err := u.mem.CheckUerInfo(ctx.Request.GetInt64("user_id"), ctx.Request.GetString("sign"), ctx.Request.GetString("timestamp")); err != nil {
		return err
	}

	ctx.Log.Info("3: 生成微信需要的statecode")
	stateCode, err := u.mem.GenerateWxStateCode(ctx.Request.GetInt64("user_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("4: 返回数据")
	config := model.GetConf(u.c)
	return map[string]interface{}{
		"wxlogin_url": config.WxPhoneLoginURL,
		"appid":       config.WxAppID,
		"state":       stateCode,
	}
}

//SaveHandle 保存绑定信息
func (u *BindWxHandler) SaveHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------保存用户的绑定信息---------")
	ctx.Log.Infof("参数为：code:%s, state:%s", ctx.Request.GetString("code"), ctx.Request.GetString("state"))

	ctx.Log.Info("1: 验证参数")
	if err := ctx.Request.Check("code", "state"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2: 验证state信息并获取openid")
	data, err := u.mem.ValidStateAndGetOpenID(ctx.Request.GetString("state"), ctx.Request.GetString("code"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3: 保存用户的openid")
	ctx.Log.Infof("data: %v+", data)
	if err := u.mem.SaveUserOpenID(data); err != nil {
		return err
	}
	return "success"
}
