package main

import (
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/hydra"
)

type mgrweb struct {
	*hydra.MicroApp
}

func main() {

	app := &mgrweb{
		hydra.NewApp(
			hydra.WithPlatName("sso_v3"),
			hydra.WithSystemName("mgrweb"),
			hydra.WithServerTypes("web"),
		),
	}
	app.Conf.WEB.SetStatic(conf.NewWebServerStaticConf().WithArchive("./static.zip"))
	app.install()
	app.Start()
}
