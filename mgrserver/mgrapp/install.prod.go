// +build prod
// +build !hbs
// +build !17supv2

package mgrapp

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/conf/vars/cache/cacheredis"
	"github.com/micro-plat/hydra/conf/vars/db"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/model"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func install() {
	hydra.Conf.Web("#api_port").Static(static.WithArchive("static.zip"),
		static.WithRewriters("/", "/index.htm", "/default.html", "/default.htm", "/external/other", "/user/index", "/sys/index", "/sys/func/*", "/sys/data/permission/*", "/user/role", "/role/auth/*", "/role/dataauth/*", "/ssocallback")).
		Header(header.WithCrossDomain()).
		Jwt(jwt.WithName("__sso_jwt__"),
			jwt.WithMode("HS512"),
			jwt.WithSecret("bf8f3171946d8d5a13cca23aa6080c8e"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/sso/login/verify", "/image/upload")).
		Sub("app", model.Conf{
			PicHost:    "http://bj.images.cdqykj.cn",
			Secret:     "B128F779D5741E701923346F7FA9F95C",
			SsoApiHost: "http://api.sso.18jiayou.com",
			Ident:      "sso",
		})
	hydra.Conf.Vars().DB().MySQLByConnStr("db", "#mysql_db_string", db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("cache", "#redis_string", cacheredis.WithDbIndex(1))
	hydra.Conf.Vars().Cache().GoCache("gocache")
	hydra.Conf.Vars().HTTP("http")
}
