package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/modules/app"
	xmenu "github.com/micro-plat/sso/modules/menu"
	"github.com/micro-plat/sso/services/base"
	"github.com/micro-plat/sso/services/function"
	"github.com/micro-plat/sso/services/image"
	"github.com/micro-plat/sso/services/member"
	"github.com/micro-plat/sso/services/menu"
	"github.com/micro-plat/sso/services/qrcode"
	"github.com/micro-plat/sso/services/role"
	"github.com/micro-plat/sso/services/system"
	"github.com/micro-plat/sso/services/user"
	"github.com/micro-plat/sso/services/wx"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) init() {
	//初始化
	r.Initializing(func(c component.IContainer) error {
		var conf app.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		app.SaveConf(c, &conf)
		if err := conf.Valid(); err != nil {
			return err
		}

		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}

		//检查缓存配置是否正确
		if _, err := c.GetCache(); err != nil {
			return err
		}
		r.Micro("/sso/wxcode/get", member.NewWxcodeHandler(conf.AppID, conf.Secret, conf.WechatTSAddr)) //发送微信验证码
		xmenu.Set(c)                                                                                    //保存全局菜单变量
		return nil
	})
	r.Micro("/sso/login", member.NewLoginHandler, "*")     //用户名密码登录
	r.Micro("/sso/login/code", member.NewCodeHandler, "*") //根据用户登录code设置jwt信息
	r.WS("/qrcode/login", qrcode.NewLoginHandler, "*")     //二维码登录（获取二维码登录地址,接收用户扫码后的消息推送）
	r.Micro("/qrcode/login", qrcode.NewLoginHandler, "*")  //二维码登录(调用二维码登录接口地址，推送到PC端登录消息)
	r.Micro("/wx/login", wx.NewLoginHandler, "*")          //微信端登录

	r.Micro("/sso/sys/get", system.NewSystemIdentHandler, "*") //根据系统编号获取系统信息
	r.Micro("/sso/menu/get", menu.NewMenuHandler, "*")         //获取用户所在系统的菜单信息
	r.Micro("/sso/popular", menu.NewPopularHandler, "*")       //获取用户所在系统的常用菜单

	//r.Micro("/sso/login/check", member.NewCheckHandler)  //用户登录状态检查，检查用户jwt是否有效
	r.Micro("/sso/member/query", member.NewQueryHandler, "*") //查询登录用户信息
	r.Micro("/sso/menu/verify", menu.NewVerifyHandler, "*")   //检查用户菜单权限

	r.Micro("/sso/user/query", user.NewUserHandler, "/user/index")
	r.Micro("/sso/user/change", user.NewUserChangeHandler, "/user/index")
	r.Micro("/sso/user/delete", user.NewUserDelHandler, "/user/index")
	r.Micro("/sso/user/info", user.NewUserInfoHandler, "/user/index")
	r.Micro("/sso/user/edit", user.NewUserEditHandler, "/user/index")
	r.Micro("/sso/user/save", user.NewUserSaveHandler, "/user/index")
	r.Micro("/sso/user/changepwd", user.NewUserPasswordHandler, "*")
	r.Micro("/sso/base/userrole", base.NewBaseUserHandler, "*")
	r.Micro("/sso/base/sys", base.NewBaseSysHandler, "*")

	r.Micro("/sso/role/query", role.NewRoleHandler, "/user/role")
	r.Micro("/sso/role/change", role.NewRoleChangeHandler, "/user/role")
	r.Micro("/sso/role/delete", role.NewRoleDelHandler, "/user/role")
	r.Micro("/sso/role/save", role.NewRoleSaveHandler, "/user/role")
	r.Micro("/sso/role/auth", role.NewRoleAuthHandler, "/user/role")
	r.Micro("/sso/role/authmenu", role.NewAuthMenuHandler, "/user/role")

	r.Micro("/sso/sys/manage", system.NewSystemHandler, "/sys/index")   //系统管理
	r.Micro("/sso/sys/edit", system.NewSystemEditHandler, "/sys/index") //系统编辑

	r.Micro("/sso/sys/func/query", function.NewSystemFuncQueryHandler, "/sys/index")   //获取功能列表
	r.Micro("/sso/sys/func/enable", function.NewSystemFuncEnableHandler, "/sys/index") //功能禁用/启用
	r.Micro("/sso/sys/func/delete", function.NewSystemFuncDeleteHandler, "/sys/index") //功能删除
	r.Micro("/sso/sys/func/edit", function.NewSystemFuncEditHandler, "/sys/index")     //功能编辑
	r.Micro("/sso/sys/func/add", function.NewSystemFuncAddHandler, "/sys/index")       //功能添加

	r.Micro("/sso/img/upload", image.NewImageHandler("./static/img", "http://192.168.7.188"), "/sys/index") //图片上传

}
