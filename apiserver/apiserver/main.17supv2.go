// +build 17supv2

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

var App = hydra.NewApp(
	hydra.WithPlatName("17sup_v2_sso", "新版sso"),
	hydra.WithSystemName("apiserver", "sso单点登录接口"),
	hydra.WithUsage("单点登录系统接口"),
	hydra.WithServerTypes(http.API),
	hydra.WithClusterName("prod"))

func main() {
	App.Start()
}
