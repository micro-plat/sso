// +build prod

package main

import "github.com/micro-plat/hydra/conf"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = false

	s.Conf.SetInput(`#mysql_db_string`, `mysql数据库连接串`, `username:password@tcp(host)/sso?charset=utf8`)
	s.Conf.SetInput(`#redis_string`, `redis连接串`, ``)
	s.Conf.SetInput(`#secret`, `用户管理系统的secret`, ``)

	s.Conf.API.SetMain(conf.NewAPIServerConf(":6677"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	s.Conf.Plat.SetDB(conf.NewMysqlConfForProd("#mysql_db_string"))

	s.Conf.API.SetSubConf("app", `
			{
				"pic_host": "http://sso.sinopecscsy.com",
				"secret":"#secret",
				"sso_api_host":"http://api.sso.18jiayou.com:6689",
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
					#redis_string
			],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}		
		`)

	// s.Conf.API.SetSubConf("header", `
	// 	{
	// 		"Access-Control-Allow-Origin": "*",
	// 		"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS",
	// 		"Access-Control-Allow-Headers": "X-Requested-With,Content-Type,__jwt__",
	// 		"Access-Control-Allow-Credentials": "true",
	// 		"Access-Control-Expose-Headers":"__jwt__"
	// 	}
	// `)

	// s.Conf.Plat.SetVarConf("db", "db", `{
	// 		"provider":"mysql",
	// 		"connString":"#mysql_db_string",
	// 		"max":8,
	// 		"maxOpen":20,
	// 		"maxIdle":10,
	// 		"lifeTime":600
	// }`)

	// s.Conf.API.SetMainConf(`{"address":":6677"}`)
	// s.Conf.SetInput(`#mysql_db_string`, `mysql数据库连接串`, `username:password@tcp(host)/sso?charset=utf8`)
}
