package main

func (s *lgweb) install() {
	s.IsDebug = false
	s.Conf.WEB.SetMainConf(`{"address":":6686"}`)
	s.Conf.WEB.SetSubConf("static", `{
		"dir":"./static",
		"rewriters":["*"],
		"exts":[".ttf",".woff",".woff2"]			
	}`)
}
