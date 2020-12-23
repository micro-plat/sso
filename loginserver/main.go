package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("sso_new", "新版sso"),
	hydra.WithSystemName("loginserver", "sso单点登录服务"),
	hydra.WithUsage("单点登录服务"),
	hydra.WithServerTypes(http.Web),
	hydra.WithClusterName("prod"))

func main() {
	App.Start()
}
