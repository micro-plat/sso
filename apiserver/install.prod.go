// +build prod

package main

import "github.com/micro-plat/hydra/conf"

func (s *SSO) install() {
	s.IsDebug = false

	s.Conf.SetInput(`#mysql_db_string`, `mysql数据库连接串`, `username:password@tcp(host)/sso?charset=utf8`)
	s.Conf.SetInput(`#redis_string`, `redis连接串`, ``)

	s.Conf.API.SetMain(conf.NewAPIServerConf(":6689"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	s.Conf.Plat.SetDB(conf.NewMysqlConfForProd("#mysql_db_string"))

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
