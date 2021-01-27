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
	"github.com/micro-plat/sso/loginserver/loginapi/modules/const/enum"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/model"

	"github.com/micro-plat/hydra/conf"
)

//Archive 压缩包名称
var Archive = "login.static.zip"
var staticOpts = []static.Option{
	static.WithArchive(Archive),
	static.WithRewriters("/", "/index.htm", "/default.html", "/default.htm", "/choose", "/refresh", "/errpage", "/bindnotice", "/wxcallback", "/bindwx", "/*/changepwd", "/*/jump", "/*/login"),
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

	hydra.Conf.Vars().Custom("loginconf", "app", model.LoginConf{
		RequireValidCode:   true,
		UserLoginFailLimit: 5,
		UserLockTime:       24 * 60 * 60,
		ValidCodeType:      enum.ValidCodeTypeWechat,
		SMSTemplateID:      "10",
		SmsSendURL:         "http://smsv1.100bm0.com:8081/sms/msg/apply",
	})

	hydra.Conf.API("6689", api.WithDNS("ssov4.100bm0.com")).
		APIKEY("ivk:///check_sign", apikey.WithInvoker("ivk:///check_sign"))

	//登录的界面配置
	hydra.Conf.Web("6687", api.WithTimeout(300, 300), api.WithDNS("ssov4.100bm0.com")).
		Static(staticOpts...).
		Header(header.WithCrossDomain(), header.WithAllowHeaders("X-Requested-With", "Content-Type")).
		Jwt(jwt.WithMode("HS512"),
			jwt.WithSecret("bf8f3171946d8d5a13cca23aa6080c8e"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/dds/dictionary/get", "/loginweb/system/config/get", "/loginweb/member/login", "/loginweb/member/bind/*", "/loginweb/member/sendcode"),
		)

}

//生产环境配置
func prodConf() {
	//配置共有配置
	hydra.Conf.Vars().Cache().GoCache("gocache")
	hydra.Conf.Vars().HTTP("http")
	hydra.Conf.Vars().DB().MySQLByConnStr("db", conf.ByInstall, db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("redis", conf.ByInstall, cacheredis.WithDbIndex(1))

	hydra.Conf.Vars().Custom("loginconf", "app", &model.LoginConf{
		RequireValidCode:   true,
		UserLoginFailLimit: 5,
		UserLockTime:       24 * 60 * 60,
		ValidCodeType:      enum.ValidCodeTypeWechat,
		SMSTemplateID:      conf.ByInstall,
		SmsSendURL:         conf.ByInstall,
	})

	hydra.Conf.API(conf.ByInstall, api.WithDNS("login.sso.18jiayou.com")).
		APIKEY("ivk:///check_sign", apikey.WithInvoker("ivk:///check_sign"))

	hydra.Conf.Web(conf.ByInstall, api.WithTimeout(300, 300), api.WithDNS("login.sso.18jiayou.com")).
		Static(staticOpts...).
		Jwt(jwt.WithMode("HS512"),
			jwt.WithSecret("f0abd74b09bcc61449d66ae5d8128c18"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/dds/dictionary/get", "/loginweb/system/config/get", "/loginweb/member/login", "/loginweb/member/bind/*", "/loginweb/member/sendcode"),
		)

}
