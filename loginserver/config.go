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
	"github.com/micro-plat/sso/loginserver/loginapi/modules/model"

	"github.com/micro-plat/hydra/conf"
)

//Archive 压缩包名称
var Archive = "login.static.zip"
var staticOpts = []static.Option{
	static.WithArchive(Archive),
	static.WithRewriters("/", "/index.htm", "/default.html", "/default.htm", "/choose", "/refresh", "/errpage", "/bindnotice", "/wxcallback/*", "/bindwx", "/*/changepwd", "/*/jump", "/*/login"),
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
	hydra.Conf.Vars().Custom("loginconf", "app", model.Conf{UserLoginFailCount: 5, UserLockTime: 24 * 60 * 60})
	hydra.Conf.Vars().DB().MySQL("db", "root", "rTo0CesHi2018Qx", "192.168.0.36:3306", "sso_new", db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("redis", `192.168.0.111:6379,192.168.0.112:6379,192.168.0.113:6379,192.168.0.114:6379,192.168.0.115:6379,192.168.0.116:6379`, cacheredis.WithDbIndex(1))

	hydra.Conf.API("6689", api.WithDNS("ssov4.100bm0.com")).Header(header.WithCrossDomain()).
		APIKEY("ivk:///check_sign", apikey.WithInvoker("ivk:///check_sign"), apikey.WithExcludes("/api/login/auth"))

	//登录的界面配置
	hydra.Conf.Web("6687", api.WithTimeout(300, 300), api.WithDNS("ssov4.100bm0.com")).
		Static(staticOpts...).
		Header(header.WithCrossDomain(), header.WithAllowHeaders("X-Requested-With", "Content-Type", "__sso_jwt__")).
		Jwt(jwt.WithName("__sso_jwt__"),
			jwt.WithMode("HS512"),
			jwt.WithSecret("bf8f3171946d8d5a13cca23aa6080c8e"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/system/webconfig", "/dds/dictionary/get", "/loginweb/system/config/get", "/loginweb/member/login", "/loginweb/member/bind/check", "/loginweb/member/bind/save", "/loginweb/member/sendcode")).
		Sub("webconf", &model.WebConf{
			Wxcallbackhost:   "http://ssov4.100bm0.com",
			Wxcallbackurl:    "/wxcallback",
			CodeLabel:        "短信验证码",
			CodeHolder:       "请输入短信验证码",
			SendBtnLable:     "获取短信验证码",
			ShowText:         "短信验证码发送成功",
			StaticImageUrl:   "",
			CompanyRight:     "四川千行你我科技有限公司",
			CompanyRightCode: "蜀ICP备20003360号",
		})

}

//生产环境配置
func prodConf() {
	//配置共有配置
	hydra.Conf.Vars().Cache().GoCache("gocache")
	hydra.Conf.Vars().HTTP("http")
	hydra.Conf.Vars().Custom("loginconf", "app", model.Conf{UserLoginFailCount: 5, UserLockTime: 24 * 60 * 60})
	hydra.Conf.Vars().DB().MySQLByConnStr("db", conf.ByInstall, db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("redis", conf.ByInstall, cacheredis.WithDbIndex(1))

	hydra.Conf.API(conf.ByInstall, api.WithDNS("api.sso.18jiayou.com")).Header(header.WithCrossDomain()).
		APIKEY("ivk:///check_sign", apikey.WithInvoker("ivk:///check_sign"), apikey.WithExcludes("/api/login/auth"))

	hydra.Conf.Web(conf.ByInstall, api.WithTimeout(300, 300), api.WithDNS("loginapi.sso.18jiayou.com")).
		Static(staticOpts...).
		Header(header.WithCrossDomain(), header.WithAllowHeaders("X-Requested-With", "Content-Type", "__sso_jwt__")).
		Jwt(jwt.WithName("__sso_jwt__"),
			jwt.WithMode("HS512"),
			jwt.WithSecret("f0abd74b09bcc61449d66ae5d8128c18"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/system/webconfig", "/dds/dictionary/get", "/loginweb/system/config/get", "/loginweb/member/login", "/loginweb/member/bind/check", "/loginweb/member/bind/save", "/loginweb/member/sendcode")).
		Sub("webconf", &model.WebConf{
			Wxcallbackhost:   "//web.sso.18jiayou.com",
			Wxcallbackurl:    "/wxcallback",
			CodeLabel:        "短信验证码",
			CodeHolder:       "请输入短信验证码",
			SendBtnLable:     "获取短信验证码",
			ShowText:         "短信验证码发送成功",
			StaticImageUrl:   "",
			CompanyRight:     "四川千行你我科技有限公司",
			CompanyRightCode: "蜀ICP备20003360号",
		})

}
