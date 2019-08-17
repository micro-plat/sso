package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/mgrapi/modules/model"
	"github.com/micro-plat/sso/mgrapi/services/base"
	"github.com/micro-plat/sso/mgrapi/services/function"
	"github.com/micro-plat/sso/mgrapi/services/image"
	"github.com/micro-plat/sso/mgrapi/services/member"
	"github.com/micro-plat/sso/mgrapi/services/menu"
	"github.com/micro-plat/sso/mgrapi/services/role"
	"github.com/micro-plat/sso/mgrapi/services/system"
	"github.com/micro-plat/sso/mgrapi/services/user"
	ssoSdk "github.com/micro-plat/sso/sso"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *SSO) init() {
	r.install()
	r.handling()

	r.Initializing(func(c component.IContainer) error {
		//检查配置信息
		var conf model.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		if err := conf.Valid(); err != nil {
			return err
		}
		model.SaveConf(c, &conf)

		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}

		//检查缓存配置是否正确
		if _, err := c.GetCache(); err != nil {
			return err
		}

		ssoCleint, err := ssoSdk.New(conf.SsoApiHost, "sso", conf.Secret)
		if err != nil {
			return err
		}
		model.SaveSSOClient(c, ssoCleint)
		return nil
	})

	r.Micro("/login", member.NewLoginHandler)       	//调用sso登录
	r.Micro("/menu", menu.NewMenuHandler)           	//菜单相关接口
	r.Micro("/base", base.NewBaseUserHandler)       	//基础数据
	r.Micro("/user", user.NewUserHandler)               //用户相关接口
	r.Micro("/auth", role.NewRoleAuthHandler)           //权限管理
	r.Micro("/role", role.NewRoleHandler)               //角色管理相关接口
	r.Micro("/sys", system.NewSystemHandler)            //系统管理相关接口
	r.Micro("/sys/func", function.NewSystemFuncHandler) //系统功能相关接口
	r.Micro("/img/upload", image.NewImageHandler("static/img"))  //图片上传
}
