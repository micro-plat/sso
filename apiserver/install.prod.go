// +build prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = false
	s.Conf.API.SetMainConf(`{"address":":6689"}`)
	s.Conf.SetInput(`#mysql_db_string`, `mysql数据库连接串`, `username:password@tcp(host)/sso?charset=utf8`)
	s.Conf.SetInput(`#redis_string`, `redis连接串`, ``)

	s.Conf.API.SetSubConf("header", `
				{
					"Access-Control-Allow-Origin": "*",
					"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS",
					"Access-Control-Allow-Headers": "X-Requested-With,Content-Type",
					"Access-Control-Allow-Credentials": "true"
				}
			`)

	s.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"mysql",
			"connString":"#mysql_db_string",
			"max":8,
			"maxOpen":20,
			"maxIdle":10,
			"lifeTime":600			
	}`)
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
