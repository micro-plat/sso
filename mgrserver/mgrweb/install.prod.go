package main

import "github.com/micro-plat/hydra/conf"

func (s *mgrweb) install() {
	s.IsDebug = false
	s.Conf.WEB.SetMainConf(`{"address":":8083"}`)
	s.Conf.WEB.SetStatic(conf.NewWebServerStaticConf().WithArchive("./static.zip"))
}
