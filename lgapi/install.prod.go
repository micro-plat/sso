// +build prod

package main

import "github.com/micro-plat/hydra/conf"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = false
	s.Conf.SetInput(`#redis_string`, `redis连接串`, ``)
	s.Conf.SetInput(`#mysql_db_string`, `mysql数据库连接串`, `username:password@tcp(host)/sso?charset=utf8`)

	s.Conf.SetInput(`#wx_appid`, `微信公众号appid`, ``)
	s.Conf.SetInput(`#wx_secret`, `微信公众号secret`, ``)
	s.Conf.SetInput(`#sendcode_key`, `调用其他部门发微信验证码的key`, ``)

	s.Conf.API.SetMain(conf.NewAPIServerConf(":6687"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	s.Conf.Plat.SetDB(conf.NewMysqlConfForProd("#mysql_db_string"))
	s.Conf.Plat.SetCache(conf.NewRedisCacheConfForProd(1, "#redis_string"))

	s.Conf.API.SetSubConf("app", `
			{
				"wxphonelogin_url":"https://open.weixin.qq.com/connect/oauth2/authorize",
				"wxlogin_url": "https://open.weixin.qq.com/connect/qrconnect",
				"wxgettoken_url":"https://api.weixin.qq.com/sns/oauth2/access_token",
				"appid":"#wx_appid",
				"secret":"#wx_secret",
				"sendcode_key":"#sendcode_key",
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
