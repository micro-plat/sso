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
			hydra.WithPlatName("sso_other"),
			hydra.WithSystemName("lgweb"),
			hydra.WithServerTypes("web"),
		),
	}
	app.install()
	app.Start()
}
