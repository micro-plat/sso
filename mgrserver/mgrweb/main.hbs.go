// +build hbs

package main

import (
	"github.com/micro-plat/hydra/hydra"
)

type mgrweb struct {
	*hydra.MicroApp
}

func main() {

	app := &mgrweb{
		hydra.NewApp(
			hydra.WithPlatName("hbs_sso"),
			hydra.WithSystemName("mgrweb"),
			hydra.WithServerTypes("web"),
		),
	}
	app.install()
	app.Start()
}
