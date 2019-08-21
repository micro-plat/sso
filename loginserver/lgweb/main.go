package main

import (
	"github.com/micro-plat/hydra/hydra"
)

type lgweb struct {
	*hydra.MicroApp
}

func main() {

	app := &lgweb{
		hydra.NewApp(
			hydra.WithPlatName("sso_v3"),
			hydra.WithSystemName("lgweb"),
			hydra.WithServerTypes("web"),
		),
	}
	app.Conf.WEB.SetStatic(conf.NewWebServerStaticConf().WithArchive("./static.zip"))
	app.install()
	app.Start()
}