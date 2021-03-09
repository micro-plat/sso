package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/processor"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/conf/vars/cache/cacheredis"
	"github.com/micro-plat/hydra/conf/vars/db"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
)

//Archive 压缩包名称
var Archive = "mgr.static.zip"
var staticOpts = []static.Option{
	static.WithAutoRewrite(),
}

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	hydra.OnReadying(func() {
		if hydra.G.IsDebug() {
			//测试环境配置
			devConf()
			return
		}

		//生产环境的配置
		prodConf()
		return
	})
}

//测试环境配置
func devConf() {
	//配置共有配置
	hydra.Conf.Vars().Cache().GoCache("gocache")
	hydra.Conf.Vars().HTTP("http")
	hydra.Conf.Vars().DB().MySQL("db", "root", "rTo0CesHi2018Qx", "192.168.0.36:3306", "sso_new", db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("redis", `192.168.0.111:6379,192.168.0.112:6379,192.168.0.113:6379,192.168.0.114:6379,192.168.0.115:6379,192.168.0.116:6379`, cacheredis.WithDbIndex(1))

	hydra.Conf.Web("6677", api.WithDNS("ssov4.100bm0.com")).
		Static(append(staticOpts, static.WithAssetsPath(Archive))...).
		Header(header.WithCrossDomain()).
		Jwt(jwt.WithMode("HS512"),
			jwt.WithSecret("bf8f3171946d8d5a13cca23aa6080c8e"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/ssoapi/sso/login/verify", "/ssoapi/image/upload", "/ssoapi/dds/dictionary/get")).
		Processor(processor.WithServicePrefix("/ssoapi")).
		Sub("app", &model.Conf{
			PicHost:     "https://static.100bm.cn",
			Secret:      "B128F779D5741E701923346F7FA9F95C",
			SSOApiHost:  "http://ssov4.100bm0.com:6689",
			AuthIgnores: []string{"/dds/**", "/base/**"},
			Ident:       "sso",
		})
}

//生产环境配置
func prodConf() {
	//配置共有配置
	hydra.Conf.Vars().Cache().GoCache("gocache")
	hydra.Conf.Vars().HTTP("http")
	hydra.Conf.Vars().DB().MySQLByConnStr("db", conf.ByInstall, db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("redis", conf.ByInstall, cacheredis.WithDbIndex(1))

	hydra.Conf.Web(conf.ByInstall, api.WithDNS(conf.ByInstall)).
		Static(append(staticOpts, static.WithAssetsPath(Archive))...).
		Jwt(jwt.WithMode("HS512"),
			jwt.WithSecret("bf8f3171946d8d5a13cca23aa6080c8e"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/ssoapi/sso/login/verify", "/ssoapi/image/upload", "/ssoapi/dds/dictionary/get")).
		Processor(processor.WithServicePrefix("/ssoapi")).
		Sub("app", &model.Conf{
			PicHost:     conf.ByInstall,
			Secret:      "B128F779D5741E701923346F7FA9F95C",
			SSOApiHost:  conf.ByInstall,
			AuthIgnores: []string{"/dds/**", "/base/**"},
			Ident:       "sso",
		})
}
