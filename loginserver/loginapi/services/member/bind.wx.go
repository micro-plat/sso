package member

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/logic"
)

//BindWxHandler 绑定微信
type BindWxHandler struct {
	mem logic.IMemberLogic
}

//NewBindWxHandler 绑定微信
func NewBindWxHandler() (u *BindWxHandler) {
	return &BindWxHandler{
		mem: logic.NewMemberLogic(),
	}
}

//InfoHandle 获取要绑定的用户信息
func (u *BindWxHandler) InfoHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取要绑定的用户信息---------")

	ctx.Log().Info("1: 获取用户信息")
	data, err := u.mem.QueryUserInfoByID(ctx.Request().GetInt64("user_id"))
	if err != nil {
		return err
	}

	ctx.Log().Info("2: 返回数据")
	return data
}

//CheckHandle 验证用户信息
func (u *BindWxHandler) CheckHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------验证要绑定的用户信息---------")

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("user_id", "sign", "timestamp"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2: 验证用户信息及签名")
	if err := u.mem.CheckUerInfo(ctx.Request().GetInt64("user_id"), ctx.Request().GetString("sign"), ctx.Request().GetString("timestamp")); err != nil {
		return err
	}

	ctx.Log().Info("3: 生成微信需要的statecode")
	stateCode, err := u.mem.GenerateWxStateCode(ctx.Request().GetInt64("user_id"))
	if err != nil {
		return err
	}

	ctx.Log().Info("4: 返回数据")
	config := model.GetConf()
	return map[string]interface{}{
		"wxlogin_url": config.WxPhoneLoginURL,
		"appid":       config.WxAppID,
		"state":       stateCode,
	}
}

//SaveHandle 保存绑定信息
func (u *BindWxHandler) SaveHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------保存用户的绑定信息---------")
	ctx.Log().Infof("参数为：code:%s, state:%s", ctx.Request().GetString("code"), ctx.Request().GetString("state"))

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("code", "state"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2: 验证state信息并获取openid")
	data, err := u.mem.ValidStateAndGetOpenID(ctx.Request().GetString("state"), ctx.Request().GetString("code"))
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 保存用户的openid")
	ctx.Log().Infof("data: %v+", data)
	if err := u.mem.UpdateUserOpenID(data); err != nil {
		return err
	}
	return "success"
}
