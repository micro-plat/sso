package users

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/modules/logic"

	"github.com/micro-plat/sso/apiserver/modules/util"
)

//UserInfoHandler is
type UserInfoHandler struct {
	container component.IContainer
	sys       logic.ISystemLogic
	m         logic.IMemberLogic
}

//NewUserInfoHandler is
func NewUserInfoHandler(container component.IContainer) (u *UserInfoHandler) {
	return &UserInfoHandler{
		container: container,
		sys:       logic.NewSystemLogic(container),
		m:         logic.NewMemberLogic(container),
	}
}

/*
* Handle: 根据用户名查询用户的相关信息
* ident:子系统标识
* username:用户名称
* timestamp:时间戳
* sign:签名
 */
func (u *UserInfoHandler) InfoHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统用户远程登录---------")

	if err := ctx.Request.Check("username", "ident", "timestamp", "sign"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	d := map[string]interface{}{
		"username":  ctx.Request.GetString("username"),
		"ident":     ctx.Request.GetString("ident"),
		"timestamp": ctx.Request.GetString("timestamp"),
	}

	ctx.Log.Info("请求原数据", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	member, err := u.m.QueryUserInfo(ctx.Request.GetString("username"),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	return member
}

func (u *UserInfoHandler) getSecret(ident string) (string, error) {
	if ident == "" {
		return "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return "", err
	}

	return data.GetString("secret"), nil
}

/*
//GetHandle 子系统获取系统信息在info中有这个接口 ****这个好像没有用到,先放在这****
func (u *UserInfoHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("-------子系统获取系统信息------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("ident", "timestamp", "sign"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	//获取secret
	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//校验签名
	d := map[string]interface{}{}
	d["ident"] = ctx.Request.GetString("ident")
	d["timestamp"] = ctx.Request.GetString("timestamp")
	ctx.Log.Info("请求用户数据：", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	data, err := u.sys.Get(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	return data
}

//PostHandle 子系统获取用户 ****这个好像没有用到,先放在这****
func (u *UserInfoHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("-------子系统获取用户------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("ident", "timestamp", "sign"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	//获取secret
	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//校验签名
	d := map[string]interface{}{}
	d["ident"] = ctx.Request.GetString("ident")
	d["timestamp"] = ctx.Request.GetString("timestamp")
	ctx.Log.Info("请求用户数据：", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	ctx.Log.Info("2. 执行操作")
	data, datas, err := u.sys.GetUsers(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"users":   data,
		"alluser": datas,
	}
}
*/
