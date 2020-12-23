// +build !prod

package main

import (
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/sso/mgrserver/modules/model"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func install() {
	hydra.Conf.Web(":6677").Static().Header(header.WithCrossDomain()).
		Jwt(jwt.WithName("__sso_jwt__"),
			jwt.WithMode("HS512"),
			jwt.WithSecret("bf8f3171946d8d5a13cca23aa6080c8e"),
			jwt.WithExpireAt(36000),
			jwt.WithHeader(),
			jwt.WithExcludes("/sso/login/verify", "/image/upload")).
		Sub("app", model.Conf{
			PicHost:    "http://sso2.100bm.cn",
			Secret:     "B128F779D5741E701923346F7FA9F95C",
			SsoApiHost: "http://api.sso.18jiayou1.com:6689",
			Ident:      "sso",
		})

	hydra.Conf.Vars().DB().MySQL("db", "root", "rTo0CesHi2018Qx", "192.168.0.36:3306", "sso", db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().Cache().Redis("redis", `192.168.0.111:6379,192.168.0.112:6379,192.168.0.113:6379,192.168.0.114:6379,192.168.0.115:6379,192.168.0.116:6379`, cacheredis.WithDbIndex(1))
}
