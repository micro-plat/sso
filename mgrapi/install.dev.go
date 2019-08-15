// +build !prod

package main

import "github.com/micro-plat/hydra/conf"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = true

	s.Conf.API.SetMain(conf.NewAPIServerConf(":6677"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	s.Conf.Plat.SetDB(conf.NewMysqlConf("root", "rTo0CesHi2018Qx", "192.168.0.36:3306", "sso").WithConnect(20, 10, 600))

	s.Conf.API.SetSubConf("app", `
			{
				"pic_host": "http://sso2.100bm.cn",
				"secret":"B128F779D5741E701923346F7FA9F95C",
				"sso_api_host":"http://192.168.106.226:6689",
				"ident":"sso"
			}			
			`)

	s.Conf.API.SetSubConf("auth", `
		{
			"jwt": {
				"exclude": [
					"/sso/login",
					"/sso/login/user",
					"/sso/sys/func/enable",
					"/sso/sys/manage/edit",
					"/sso/login/code",
					"/sso/sys/get",
					"/sso/ident",
					"/sso/user/bind",
					"/sso/notify/send",
					"/sso/img/upload",
					"/sso/user/getall",
					"/sso/user/info",
					"/sso/user/save",
					"/sso/user/edit",
					"/sso/user/delete",
					"/sso/role/query",
					"/sso/menu/get",
					"/sso/sys/func/query",
					"/sso/sys/manage/up",
					"/sso/sys/manage/down",
					"/sso/user/changepwd"
				],
				"expireAt": 36000,
				"source":"H",
				"mode": "HS512",
				"name": "__jwt__",
				"secret": "12345678"
			}
		}
		`)

	s.Conf.Plat.SetVarConf("cache", "cache", `
		{
			"proto":"redis",
			"addrs":[
					"192.168.0.111:6379","192.168.0.112:6379","192.168.0.113:6379","192.168.0.114:6379","192.168.0.115:6379","192.168.0.116:6379"
			],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}		
		`)

	// s.Conf.Plat.SetVarConf("db", "db", `{
	// 	"provider":"mysql",
	// 	"connString":"root:rTo0CesHi2018Qx@tcp(192.168.0.36:3306)/sso?charset=utf8",
	// 	"max":8,
	// 	"maxOpen":20,
	//  	"maxIdle":10,
	//  	"lifeTime":600
	// }`)

	// s.Conf.API.SetSubConf("header", `
	// 		{
	// 			"Access-Control-Allow-Origin": "*",
	// 			"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS",
	// 			"Access-Control-Allow-Headers": "X-Requested-With,Content-Type,__jwt__",
	// 			"Access-Control-Allow-Credentials": "true",
	// 			"Access-Control-Expose-Headers":"__jwt__"
	// 		}
	// 	`)

}
