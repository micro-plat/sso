package sso

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/lib4go/errs"
)

var onceLock sync.Once

func init() {
	onceLock.Do(func() {
		hydra.OnReady(func() {
			app := hydra.S
			app.Micro("/sso/login/verify", loginVerify)
			app.Micro("/sso/member/menus/get", userMenus)
			app.Micro("/sso/member/systems/get", userSystems)
			app.Micro("/sso/member/all/get", getAllUser)
			app.Micro("/sso/system/info/get", systemInfo)
			app.Micro("/sso/member/tag/display", getTags)

			app.Micro("/sso/member/changepwd", changePwd)
			app.Micro("/sso/member/forgetpwd", forgetPwd)
			fmt.Println("sso-sdk")
			app.Micro("/sso/system/config", GetSystemConfig)
		})
	})

}

//Config 自动生成相关的api接口(登录回调验证、获取菜单、获取系统信息)
func Config(ssoApiHost, ident, secret string) error {
	if err := saveSSOClient(ssoApiHost, ident, secret); err != nil {
		return err
	}
	return nil
}

//loginVerify 登录验证，如果成功了写子系统jwt
func loginVerify(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------sso登录后去取登录用户---------")

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("code"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("code不能为空"))
	}

	ctx.Log().Info("2: 调用sso api 用code取用户信息")
	data, err := GetSSOClient().CheckCodeLogin(ctx.Request().GetString("code"))
	if err != nil {
		return err
	}

	ctx.Log().Infof("data: %v", data)

	ctx.User().Auth().Response(data)

	ctx.Log().Info("3: 返回用户数据")
	return map[string]interface{}{
		"user_name": data.UserName,
		"role_name": data.RoleName,
	}
}

//userMenus 用户菜单信息
func userMenus(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------去sso获取菜单数据---------")

	ctx.Log().Info("1: 获取登录用户信息")
	mem := GetMember(ctx)

	ctx.Log().Info("2: 远程获取菜单信息")
	menus, err := GetSSOClient().GetUserMenu(int(mem.UserID))
	if err != nil {
		return err
	}

	ctx.Log().Info("3: 远程获取菜单信息")
	return menus
}

//userSystems 用户有权限的其他系统信息
func userSystems(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------获取用户可用的其他系统--------")

	ctx.Log().Info("1.获取用户信息")
	mem := GetMember(ctx)

	ctx.Log().Info("2.获取数据")
	data, err := GetSSOClient().GetUserOtherSystems(int(mem.UserID))
	if err != nil {
		return err
	}

	ctx.Log().Info("3.返回结果")
	return data
}

//SystemInfo 当前的系统信息
func systemInfo(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------去sso获取系统信息----------")

	ctx.Log().Info("1. 执行操作")
	data, err := GetSSOClient().GetSystemInfo()
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 返回数据")
	return data

}

//getAllUser 获取所有用户信息
func getAllUser(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------获取所有用户信息----------")

	ctx.Log().Info("1. 获取登录信息")
	mem := GetMember(ctx)

	ctx.Log().Info("1. 执行操作")
	data, err := GetSSOClient().GetAllUser(mem.Source, mem.SourceID)
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 返回数据")
	return data
}

//getTags 按钮是否显示
func getTags(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------获取页面的按钮是否显示----------")

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("tags"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("tags不能为空,如:user_new,user_delete"))
	}
	mem := GetMember(ctx)

	ctx.Log().Info("2. 执行操作")
	data, err := GetSSOClient().GetUserDisplayTags(int(mem.UserID), ctx.Request().GetString("tags"))
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 返回数据")
	return data
}

//forgetPwd 忘记密码并修改密码
func forgetPwd(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------忘记密码并修改密码----------")

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("source", "source_id", "possword"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2: 调用sdk忘记密码并修改密码")
	return GetSSOClient().ForgetPwd(ctx.Request().GetString("source"), ctx.Request().GetString("source_id"), ctx.Request().GetString("possword"))
}

//ChangePwd 修改密码
func changePwd(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("--------修改密码----------")

	ctx.Log().Info("1: 验证参数")
	if err := ctx.Request().Check("expassword", "newpassword"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("2: 调用sdk修改密码")
	mem := GetMember(ctx)
	return GetSSOClient().ChangePwd(mem.UserID, ctx.Request().GetString("expassword"), ctx.Request().GetString("newpassword"))
}

//GetSystemConfig VueConfig
func GetSystemConfig(ctx hydra.IContext) interface{} {
	configData := map[string]interface{}{}
	if _, err := ctx.APPConf().GetServerConf().GetSubObject("webconf", &configData); err != nil {
		return err
	}
	jwtConf, err := jwt.GetConf(ctx.APPConf().GetServerConf())
	if err != nil {
		return err
	}

	configData["jwt_name"] = jwtConf.Name
	configData["jwt_source"] = jwtConf.Source
	configData["jwt_authurl"] = jwtConf.AuthURL

	ctx.Response().ContentType("text/plain")
	bytes, _ := json.Marshal(configData)
	return fmt.Sprintf("window.globalConfig=%s", string(bytes))
}

/* getUserDataPermission 获取 [数据权限] 生成相应的sql语句
 *
 */
func GetDataPermission(userID int64, tableName string, opt ...PermissionOption) (string, error) {
	return GetSSOClient().getUserDataPermission(userID, tableName, opt...)
}

/*AddUser 增加用户
*userName 用户名没有就传手机号
*mobile 手机号
*fullName 中文名
*targetIdent 要给那个系统增加用户
*source 来源, 加油站、公司、下游渠道等
*sourceID 来源编号
 */
func AddUser(userName, mobile, fullName, targetIdent, source, sourceSecrect string, sourceID string, roleID int) error {
	return GetSSOClient().AddUser(userName, mobile, fullName, targetIdent, source, sourceSecrect, sourceID, roleID)
}

//Login 用户密码登录, 密码请用md5加密
func Login(userName, password string) (LoginState, error) {
	return GetSSOClient().Login(userName, password)
}

//GetSystemInfo 获取系统信息
func GetSystemInfo() (data *System, err error) {
	return GetSSOClient().GetSystemInfo()
}

//GetCurrentUserInfo 从服务端实时获取当前登录用户信息(主要是为禁用用户什么的)
func GetCurrentUserInfo(userName string) (info *User, err error) {
	return GetSSOClient().GetUserInfoByName(userName)
}
