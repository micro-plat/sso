// +build prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = false

	s.Conf.API.SetMainConf(`{"address":":6688"}`)
	s.Conf.API.SetSubConf("app", `
			{
				"pic_host": "#pic_host",
				"secret":"#secret",
				"sso_api_host":"http://192.168.106.226:6689",
				"sso_jump_host":"#sso_jump_host",
				"ident":"sso"
			}			
			`)
	s.Conf.API.SetSubConf("header", `
				{
					"Access-Control-Allow-Origin": "*",
					"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS",
					"Access-Control-Allow-Headers": "X-Requested-With,Content-Type",
					"Access-Control-Allow-Credentials": "true"
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
				"mode": "HS512",
				"name": "__jwt__",
				"secret": "12345678"
			}
		}
		`)

	s.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"mysql",
			"connString":"#db_string",
			"max":8,
			"maxOpen":20,
			"maxIdle":10,
			"lifeTime":600	
	}`)
	//sso/123456@orcl136
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
}
