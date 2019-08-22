package main

import "github.com/micro-plat/hydra/conf"

func (s *lgweb) install() {
	s.IsDebug = false
	s.Conf.WEB.SetMainConf(`{"address":":80"}`)
	s.Conf.WEB.SetStatic(conf.NewWebServerStaticConf().WithArchive("./static.zip"))
}
