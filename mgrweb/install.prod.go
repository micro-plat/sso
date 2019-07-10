package main

func (s *mgrweb) install() {
	s.Conf.WEB.SetMainConf(`{"address":":8090"}`)
	s.Conf.WEB.SetSubConf("static", `{
		"dir":"./static",
		"rewriters":["*"],
		"exts":[".ttf",".woff",".woff2"]			
		}`)
}
