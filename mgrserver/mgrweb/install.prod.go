package main

import "github.com/micro-plat/hydra/conf"

func (s *mgrweb) install() {
	s.IsDebug = false
	s.Conf.WEB.SetMainConf(`{"address":":6678"}`)
	s.Conf.WEB.SetStatic(conf.NewWebServerStaticConf().WithArchive("./static.zip"))
}
