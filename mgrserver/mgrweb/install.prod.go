package main

import "github.com/micro-plat/hydra/conf"

func (s *mgrweb) install() {
	s.IsDebug = false
	//s.Conf.WEB.SetMainConf(`{"address":":8083", "host":"web.sso.18jiayou1.com"}`)
	s.Conf.WEB.SetMainConf(`{"address":":80", "host":"web.sso.100bm.cn"}`)
	s.Conf.WEB.SetStatic(conf.NewWebServerStaticConf().WithArchive("./static.zip"))
}
