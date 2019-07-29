// +build !prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = true

	s.Conf.API.SetMainConf(`{"address":":6687"}`)
	s.Conf.API.SetSubConf("header", `
		{
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS",
			"Access-Control-Allow-Headers": "X-Requested-With,Content-Type",
			"Access-Control-Allow-Credentials": "true"
		}
	`)
	s.Conf.API.SetSubConf("app", `
		{
			"wxlogin_url": "https://open.weixin.qq.com/connect/qrconnect",
			"wxgettoken_url":"https: //api.weixin.qq.com/sns/oauth2/access_token",
			"appid":"wxbdc5610cc59c1631",
			"secret":"123456",
			"sendcode_key":"qxnw123456",
			"sendcodereq_url":"http://user.18pingtai.cn:9002/SendVerifyCodeHandler.ashx",
			"sendcode_timeout":30,
			"requirewx_login":false,
			"require_code":true
		}			
	`)

	s.Conf.API.SetSubConf("auth", `
		{
			"jwt": {
				"exclude": [
					"/lg/login/post",
					"/lg/login/wxconf",
					"/lg/login/wxcheck",
					"/lg/login/wxvalidcode",
					"/lg/login/typeconf"
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
		"connString":"root:rTo0CesHi2018Qx@tcp(192.168.0.36:3306)/sso?charset=utf8",
		"max":8,
		"maxOpen":20,
	 	"maxIdle":10,
	 	"lifeTime":600
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
