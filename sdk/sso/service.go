package sso

import (
	"fmt"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
)

//Bind 自动生成相关的api接口(登录回调验证、获取菜单、获取系统信息)
func Bind(app *hydra.MicroApp, ssoClient *Client) {
	saveSSOClient(ssoClient)

	app.Micro("/sso/login/verify", loginVerify)
	app.Micro("/sso/member/menus/get", userMenus)
	app.Micro("/sso/member/systems/get", userSystems)
	app.Micro("/sso/system/info/get", systemInfo)
}

//loginVerify 登录验证，如果成功了写子系统jwt
func loginVerify(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------sso登录后去取登录用户---------")

	ctx.Log.Info("1: 验证参数")
	if err := ctx.Request.Check("code"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("code不能为空"))
	}

	ctx.Log.Info("2: 调用sso api 用code取用户信息")
	data, err := getSSOClient().CheckCodeLogin(ctx.Request.GetString("code"))
	if err != nil {
		return err
	}

	ctx.Log.Infof("data: %v", data)
	ctx.Response.SetJWT(data)

	ctx.Log.Info("3: 返回用户数据")
	return map[string]interface{}{
		"user_name": data.UserName,
		"role_name": data.RoleName,
	}
}

//userMenus 用户菜单信息
func userMenus(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------去sso获取菜单数据---------")

	ctx.Log.Info("1: 获取登录用户信息")
	mem := GetMember(ctx)

	ctx.Log.Info("2: 远程获取菜单信息")
	menus, err := getSSOClient().GetUserMenu(int(mem.UserID))
	if err != nil {
		return err
	}

	ctx.Log.Info("3: 远程获取菜单信息")
	return menus
}

//userSystems 用户有权限的其他系统信息
func userSystems(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------获取用户可用的其他系统--------")

	ctx.Log.Info("1.获取用户信息")
	mem := GetMember(ctx)

	ctx.Log.Info("2.获取数据")
	data, err := getSSOClient().GetUserOtherSystems(int(mem.UserID))
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//SystemInfo 当前的系统信息
func systemInfo(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------去sso获取系统信息----------")

	ctx.Log.Info("1. 执行操作")
	data, err := getSSOClient().GetSystemInfo()
	if err != nil {
		return err
	}

	ctx.Log.Info("2. 返回数据")
	return data

}
