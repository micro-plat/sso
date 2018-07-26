package main

import "github.com/micro-plat/hydra/hydra"

//SSO 单点登录系统
type SSO struct {
	*hydra.MicroApp
}

func main() {
	app := &SSO{
		hydra.NewApp(
			hydra.WithPlatName("sso"),
			hydra.WithSystemName("sso"),
			hydra.WithServerTypes("api-ws"),
			hydra.WithDebug()),
	}

	app.init()
	app.install()
	app.handing()

	app.Start()
}
