package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("sso_v4", "sso v4"),
	hydra.WithSystemName("mgrserver", "SSO管理系统"),
	hydra.WithUsage("单点登录管理系统"),
	hydra.WithServerTypes(http.Web),
	hydra.WithClusterName("prod"))

func main() {
	App.Start()
}
