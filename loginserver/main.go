package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("sso_v4", "sso v4"),
	hydra.WithSystemName("loginserver", "SSO登录网址"),
	hydra.WithUsage("单点登录服务"),
	hydra.WithServerTypes(http.Web, http.API),
	hydra.WithClusterName("prod"))

func main() {
	App.Start()
}
