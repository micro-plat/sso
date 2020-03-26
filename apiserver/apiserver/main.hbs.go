// +build hbs

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/hydra"
)

//SSO 单点登录系统
type SSO struct {
	*hydra.MicroApp
}

func main() {
	app := &SSO{
		hydra.NewApp(
			hydra.WithPlatName("hbs_sso"),
			hydra.WithSystemName("apiserver"),
			hydra.WithServerTypes("api")),
	}
	app.init()
	app.Start()
}