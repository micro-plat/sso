package main

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/dds/dds"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
	"github.com/micro-plat/sso/sassserver/sassapi/services/sso"
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
		dds.Bind(r.MicroApp, "db")
		sso.BindSass(r.MicroApp, "oil", "db")
		return nil
	})
}
