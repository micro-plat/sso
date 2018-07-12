package main

import "github.com/micro-plat/hydra/hydra"

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("ums-wl"),
		hydra.WithSystemName("sso"),
		hydra.WithServerTypes("api-ws"),
		hydra.WithDebug())
	bind(app)
	app.Start()
}
