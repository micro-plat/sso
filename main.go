package main

import "github.com/micro-plat/hydra/hydra"

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("sso"),
		hydra.WithSystemName("sso"),
		hydra.WithServerTypes("api"))
	bind(app)
	app.Start()
}
