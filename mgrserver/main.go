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
			hydra.WithPlatName("ums"),
			hydra.WithSystemName("ums"),
			hydra.WithServerTypes("web"),
		),
	}
	app.install()
	app.Start()
}
