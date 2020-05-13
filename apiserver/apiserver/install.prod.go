// +build prod

package main

import (
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/sso/common/module/model"
)

func (s *SSO) install() {
	s.Conf.SetInput(`#mysql_db_string`, `mysql数据库连接串`, `username:password@tcp(host)/sso?charset=utf8`)
	s.Conf.SetInput(`#redis_string`, `redis连接串`, ``)

	s.Conf.API.SetMain(conf.NewAPIServerConf(":6689"))
	s.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	s.Conf.Plat.SetDB(conf.NewMysqlConfForProd("#mysql_db_string"))
	s.Conf.Plat.SetCache(conf.NewRedisCacheConfForProd(1, "#redis_string"))

	s.Conf.API.SetApp(model.Conf{
		UserLoginFailCount:    5,
		UserLockTime:          24 * 60 * 60,
		AddUserUseDefaultRole: `{"mer17sup":1}`,
	})
}
