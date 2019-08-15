// +build !prod

package main

import "github.com/micro-plat/hydra/conf"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = true

	s.Conf.API.SetMain(conf.NewAPIServerConf(":6687"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	s.Conf.Plat.SetDB(conf.NewMysqlConf("root", "rTo0CesHi2018Qx", "192.168.0.36:3306", "sso").WithConnect(20, 10, 600))
	s.Conf.Plat.SetCache(conf.NewRedisCacheConf(1, "192.168.0.111:6379",
		"192.168.0.112:6379", "192.168.0.113:6379", "192.168.0.114:6379",
		"192.168.0.115:6379", "192.168.0.116:6379"))

	s.Conf.API.SetSubConf("app", `
		{
			"wxphonelogin_url":"https://open.weixin.qq.com/connect/oauth2/authorize",
			"wxlogin_url": "https://open.weixin.qq.com/connect/qrconnect",
			"wxgettoken_url":"https://api.weixin.qq.com/sns/oauth2/access_token",
			"appid":"wx1234566",
			"secret":"123456",
			"sendcode_key":"123456",
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
					"/lg/login/typeconf",
					"/lg/login/getwxstate",
					"/lg/user/check",
					"/lg/user/wxbind"
				],
				"expireAt": 36000,
				"mode": "HS512",
				"name": "__sso_jwt__",
				"secret": "12345678"
			}
		}
		`)

}
