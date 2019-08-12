package main

func (s *mgrweb) install() {
	s.IsDebug = false
	s.Conf.WEB.SetMainConf(`{"address":":8081"}`)
	s.Conf.WEB.SetSubConf("static", `{
		"dir":"./static",
		"rewriters":["*"],
		"exts":[".ttf",".woff",".woff2"]			
		}`)
}
