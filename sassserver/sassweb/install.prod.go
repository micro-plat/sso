package main

func (s *mgrweb) install() {
	s.IsDebug = false
	//s.Conf.WEB.SetMainConf(`{"address":":8083", "host":"web.sso.18jiayou1.com"}`)
	s.Conf.WEB.SetMainConf(`{"address":":8083"}`)
	//s.Conf.WEB.SetStatic(conf.NewWebServerStaticConf().WithArchive("./static.zip"))
	s.Conf.WEB.SetSubConf("static", `{
		"dir":"./static",
		"rewriters":["*"],
		"exts":[".ttf",".woff",".woff2"]			
	}`)
}
