// +build !prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.Conf.API.SetMainConf(`{"address":":6688"}`)
	s.Conf.API.SetSubConf("app", `
			{
				"web_host_name": "http://sso2.100bm.cn"
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
				"/sso/sys/func/enable",
				"/sso/sys/manage/edit",
				"/sso/login/code",
				"/subsys/info",
				"/subsys/pwd",
				"/subsys/user",
				"/subsys/login",
				"/subsys/menu",
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
				"/sso/user/changepwd"],
				"expireAt": 36000,
				"mode": "HS512",
				"name": "__jwt__",
				"secret": "12345678"
			}
		}
		`)

	s.Conf.Plat.SetVarConf("db", "db", `{			
		"provider":"mysql",
		"connString":"root:rTo0CesHi2018Qx$@tcp(192.168.0.36:3306)/sso?charset=utf8",
		"max":8,
		"maxOpen":20,
	 	"maxIdle":10,
	 	"lifeTime":600
	}`)

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

}
