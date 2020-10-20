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
			// hydra.WithPlatName("17sup_v2_sso"),
			hydra.WithPlatName("sso_v3"),
			hydra.WithSystemName("lgweb"),
			hydra.WithServerTypes("web"),
		),
	}
	app.install()
	app.Start()
}
