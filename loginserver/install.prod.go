// +build prod

package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/auth/apikey"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/conf/vars/cache/cacheredis"
	"github.com/micro-plat/hydra/conf/vars/db"
	"github.com/micro-plat/sso/common/module/model"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func install() {
	hydra.Conf.API("#api_port", api.WithDNS("api.sso.18jiayou.com")).Header(header.WithCrossDomain()).
		APIKEY("SVS:///check_sign", apikey.WithExcludes("/sso/login/verify", "/image/upload"))

	hydra.Conf.Web("#web_port", api.WithTimeout(300, 300), api.WithDNS("loginapi.sso.18jiayou.com")).
		Static(static.WithArchive("static.zip"),
			static.WithRewriters("/", "/index.htm", "/default.html", "/default.htm", "/choose", "/refresh", "/errpage", "/bindnotice", "/wxcallback/*", "/bindwx", "/*/changepwd", "/*/jump", "/*/login")).
		Header(header.WithCrossDomain(), header.WithAllowHeaders("X-Requested-With", "Content-Type", "__sso_jwt__")).
		Jwt(jwt.WithName("__sso_jwt__"),
			jwt.WithMode("HS512"),
			jwt.WithSecret("f0abd74b09bcc61449d66ae5d8128c18"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/mgrweb/system/config/get", "/mgrweb/member/login", "/mgrweb/member/bind/check", "/mgrweb/member/bind/save", "/mgrweb/member/sendcode"))

	hydra.Conf.Vars().Custom("conf", "app", model.Conf{
		UserLoginFailCount:    5,
		UserLockTime:          24 * 60 * 60,
		AddUserUseDefaultRole: `{"mer17sup":1}`,
	})
	hydra.Conf.Vars().DB().MySQLByConnStr("db", "#mysql_db_string", db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("cache", "#redis_string", cacheredis.WithDbIndex(1))
	hydra.Conf.Vars().Cache().GoCache("gocache")
	hydra.Conf.Vars().HTTP("http")
}
