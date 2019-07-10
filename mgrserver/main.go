package main

import (
	"github.com/micro-plat/hydra/hydra"
)

type ums struct {
	*hydra.MicroApp
}

func main() {

	app := &ums{
		hydra.NewApp(
			hydra.WithPlatName("sso"),
			hydra.WithSystemName("mgrweb"),
			hydra.WithServerTypes("web"),
		),
	}
	app.install()
	app.Start()
}
