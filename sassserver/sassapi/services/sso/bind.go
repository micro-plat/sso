package sso

import (
	"github.com/micro-plat/hydra/hydra"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model/config"
	"github.com/micro-plat/sso/sassserver/sassapi/services/login"
	"github.com/micro-plat/sso/sassserver/sassapi/services/manage/base"
	"github.com/micro-plat/sso/sassserver/sassapi/services/manage/role"
	"github.com/micro-plat/sso/sassserver/sassapi/services/manage/user"
)

/*BindSass　绑定面向客户系统的api,只会有部分功能
ident 系统标识
dbName dbName
*/
func BindSass(app *hydra.MicroApp, ident, dbName string) {
	config.SetConfig(dbName, ident)

	//用户登录获取信息相关接口
	app.Micro("/sso/login/verify", login.NewLoginHandler)         //用户登录(这个要验证加油站信息与公司信息是否禁用,用户这边没有相关状态信息)
	app.Micro("/sso/verifycode/get", login.NewVerifyCodeHandler)  //获取图片验证码
	app.Micro("/sso/member/menus/get", login.NewMenuHandler)      //获取用户菜单
	app.Micro("/sso/member/changepwd", login.NewChangePwdHandler) //修改密码
	app.Micro("/sso/system/info/get", login.NewSystemHandler)     //这个可能要业务系统那边去实现(用户这边没有加油站和公司信息)

	//管理员管理用户及相应的角色管理
	app.Micro("/user", user.NewUserHandler)     //用户相关接口(第一版会放出去这里面的接口)
	app.Micro("/base", base.NewBaseUserHandler) //基础数据 获取角色信息(第一版会放出去这里面的接口)
	app.Micro("/auth", role.NewRoleAuthHandler) //菜单权限管理 下个版本会开放给sass用户
	app.Micro("/role", role.NewRoleHandler)     //角色管理相关接口 下个版本会开放给sass用户
}
