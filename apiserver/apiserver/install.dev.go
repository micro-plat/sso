// +build !prod

package main

import "github.com/micro-plat/hydra/conf"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = true

	s.Conf.API.SetMain(conf.NewAPIServerConf(":6689").WithDNS("api.sso.18jiayou1.com"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	s.Conf.Plat.SetDB(conf.NewMysqlConf("root", "rTo0CesHi2018Qx", "192.168.0.36:3306", "sso").WithConnect(20, 10, 600))
	s.Conf.Plat.SetCache(conf.NewRedisCacheConf(1, "192.168.0.111:6379",
		"192.168.0.112:6379", "192.168.0.113:6379", "192.168.0.114:6379",
		"192.168.0.115:6379", "192.168.0.116:6379"))
}
