// +build prod

package main

import (
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/sso/common/module/model"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {

	s.IsDebug = false
	s.Conf.SetInput(`#redis_string`, `redis连接串`, ``)
	s.Conf.SetInput(`#mysql_db_string`, `mysql数据库连接串`, `username:password@tcp(host)/sso?charset=utf8`)

	s.Conf.API.SetMain(conf.NewAPIServerConf(":6687"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain().WithAllowHeaders("X-Requested-With", "Content-Type", "__sso_jwt__"))
	s.Conf.Plat.SetDB(conf.NewMysqlConfForProd("#mysql_db_string"))
	s.Conf.Plat.SetCache(conf.NewRedisCacheConfForProd(1, "#redis_string"))

	s.Conf.API.SetAuthes(conf.NewAuthes().WithJWT(
		conf.NewJWT("__sso_jwt__", "HS512", "f0abd74b09bcc61449d66ae5d8128c18", 36000, "/system/config/get", "/member/login", "/member/bind/check", "/member/bind/save", "/member/sendcode").WithHeaderStore()))

	s.Conf.API.SetApp(model.Conf{
		UserLoginFailCount: 5,
		UserLockTime:       24 * 60 * 60,
	})
}
