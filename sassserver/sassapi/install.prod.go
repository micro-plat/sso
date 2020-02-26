// +build prod

package main

import (
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/sso/sassserver/sassapi/modules/model"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = false

	s.Conf.SetInput(`#mysql_db_string`, `mysql数据库连接串`, `username:password@tcp(host)/sso?charset=utf8`)
	s.Conf.SetInput(`#redis_string`, `redis连接串`, ``)
	s.Conf.API.SetMain(conf.NewAPIServerConf(":6678"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	s.Conf.Plat.SetDB(conf.NewMysqlConfForProd("#mysql_db_string"))
	s.Conf.Plat.SetCache(conf.NewRedisCacheConfForProd(1, "#redis_string"))

	s.Conf.API.SetAuthes(
		conf.NewAuthes().WithJWT(
			conf.NewJWT("__jwt__", "HS512", "bf8f3171946d8d5a13cca23aa6080c8e", 36000, "/sso/login/verify", "/image/upload").WithHeaderStore()))

	s.Conf.API.SetApp(model.Conf{
		PicHost:    "http://bj.images.cdqykj.cn",
		Secret:     "#secret",
		SsoApiHost: "http://api.sso.18jiayou.com:80",
		Ident:      "sso",
	})
}
