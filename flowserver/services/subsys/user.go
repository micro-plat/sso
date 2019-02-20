package subsys

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/flowserver/modules/member"
	"github.com/micro-plat/sso/flowserver/modules/system"
	"github.com/micro-plat/sso/flowserver/modules/util"
)

//UserHandler is
type UserHandler struct {
	container component.IContainer
	sys       system.ISystem
	m         member.IMember
}

//NewUserHandler is
func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		container: container,
		sys:       system.NewSystem(container),
		m:         member.NewMember(container),
	}
}

//GetHandle 子系统获取系统信息
func (u *UserHandler) GetHandle(ctx *context.Context) (r interface{}) {

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

//PostHandle 子系统获取用户
func (u *UserHandler) PostHandle(ctx *context.Context) (r interface{}) {

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

//Handle 子系统远程登录
func (u *UserHandler) InfoHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------子系统用户远程登录---------")

	ctx.Log.Info("1.检查参数")
	if err := ctx.Request.Check("username", "ident", "timestamp", "sign"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	//获取secret
	secret, err := u.getSecret(ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}
	//校验签名
	d := map[string]interface{}{}
	d["username"] = ctx.Request.GetString("username")
	d["ident"] = ctx.Request.GetString("ident")
	d["timestamp"] = ctx.Request.GetString("timestamp")
	ctx.Log.Info("请求原数据", d)
	if ok := util.VerifySign(ctx, d, secret, ctx.Request.GetString("sign")); ok != true {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, "sign签名错误")
	}

	ctx.Log.Info("2.执行操作")
	//处理用户登录
	member, err := u.m.QueryUserInfo(ctx.Request.GetString("username"),
		ctx.Request.GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return member

}

func (u *UserHandler) getSecret(ident string) (string, error) {
	if ident == "" {
		return "", context.NewError(context.ERR_NOT_ACCEPTABLE, "ident not exists")
	}
	data, err := u.sys.Get(ident)
	if err != nil {
		return "", err
	}

	return data.GetString("secret"), nil
}
