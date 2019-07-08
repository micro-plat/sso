// +build prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = false
	//s.PlatName = "sso2"
	// s.Conf.SetInput("email", "邮箱地址", "接收账户确认邮件时使用", func(v string) (string, error) {
	// 	if !strings.Contains(v, "@") {
	// 		return "", fmt.Errorf("请输入正确的邮箱地址")
	// 	}
	// 	return strings.Replace(v, "@", "\\@", -1), nil
	// })
	// s.Conf.SetInput("#wx_host_name", "服务器域名", "以http开头")

	s.Conf.API.SetMainConf(`{"address":":6688"}`)
	s.Conf.API.SetSubConf("app", `
			{
				"web_host_name": "#web_host_name"
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
			"provider":"ora",
			"connString":"#db_string",
			"maxOpen":10,
			"maxIdle":1,
			"lifeTime":10		
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
