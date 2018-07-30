// +build !prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.Conf.API.SetMainConf(`{"address":":9091"}`)
	s.Conf.API.SetSubConf("app", `
			{
				"qrlogin-check-url":"http://sso.100bm.cn/member/wxlogin",
				"wx-login-url":"http://sso.100bm.cn/member/wxlogin",
				"appid":"wx9e02ddcc88e13fd4",
				"secret":"45d25cb71f3bee254c2bc6fc0dc0caf1",
				"wechat-url":"http://59.151.30.153:9999/wx9e02ddcc88e13fd4/wechat/token/get",
				"hostname":"http://sso.100bm.cn"
			}			
			`)
	s.Conf.API.SetSubConf("header", `
				{
					"Access-Control-Allow-Origin": "*", 
					"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS", 
					"Access-Control-Allow-Headers": "__jwt__", 
					"Access-Control-Allow-Credentials": "true"
				}
			`)

	s.Conf.API.SetSubConf("auth", `
		{
			"jwt": {
				"exclude": ["/sso/login","/sso/login/code","/sso/wxcode/get","/sso/sys/get","/qrcode/login","/qrcode/login/put","/sso/user/bind"],
				"expireAt": 36000,
				"mode": "HS512",
				"name": "__jwt__",
				"secret": "12345678"
			}
		}
		`)

	s.Conf.WS.SetSubConf("app", `
			{
				"qrlogin-check-url":"http://sso.100bm.cn/member/wxlogin",
				"wx-login-url":"http://sso.100bm.cn/member/wxlogin",
				"appid":"wx9e02ddcc88e13fd4",
				"secret":"45d25cb71f3bee254c2bc6fc0dc0caf1",
				"wechat-url":"http://59.151.30.153:9999/wx9e02ddcc88e13fd4/wechat/token/get"
			}			
			`)
	s.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"ora",
			"connString":"sso/123456@orcl136",
			"maxOpen":10,
			"maxIdle":1,
			"lifeTime":10		
	}`)

	s.Conf.Plat.SetVarConf("cache", "cache", `
		{
			"proto":"redis",
			"addrs":[
					"192.168.0.111:6379",
					"192.168.0.112:6379",
					"192.168.0.113:6379",
					"192.168.0.114:6379",
					"192.168.0.115:6379",
					"192.168.0.116:6379"
			],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}
		
		`)
}
