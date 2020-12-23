// +build !prod

package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/vars/cache/cacheredis"
	"github.com/micro-plat/hydra/conf/vars/db"
	"github.com/micro-plat/sso/common/module/model"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func install() {
	hydra.Conf.Web(":6687", api.WithTimeout(300, 300), api.WithDNS("loginapi.sso.18jiayou1.com")).
		Static().
		Header(header.WithCrossDomain(), header.WithAllowHeaders("X-Requested-With", "Content-Type", "__sso_jwt__")).
		Jwt(jwt.WithName("__sso_jwt__"),
			jwt.WithMode("HS512"),
			jwt.WithSecret("bf8f3171946d8d5a13cca23aa6080c8e"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/system/config/get", "/member/login", "/member/bind/check", "/member/bind/save", "/member/sendcode")).
		Sub("app", model.Conf{
			UserLoginFailCount: 5,
			UserLockTime:       24 * 60 * 60,
		})

	hydra.Conf.Vars().DB().MySQL("db", "root", "rTo0CesHi2018Qx", "192.168.0.36:3306", "sso", db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("redis", `192.168.0.111:6379,192.168.0.112:6379,192.168.0.113:6379,192.168.0.114:6379,192.168.0.115:6379,192.168.0.116:6379`, cacheredis.WithDbIndex(1))
	hydra.Conf.Vars().Cache().GoCache("gocache")
	hydra.Conf.Vars().HTTP("http")
}
