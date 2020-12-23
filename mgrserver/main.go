// +build !hbs
// +build !17supv2

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

var App = hydra.NewApp(
	hydra.WithPlatName("sso_new", "新版sso"),
	hydra.WithSystemName("mgrserver", "sso单点登录管理系统"),
	hydra.WithUsage("单点登录管理系统"),
	hydra.WithServerTypes(http.Web),
	hydra.WithClusterName("prod"))

func main() {
	App.Start()
}
