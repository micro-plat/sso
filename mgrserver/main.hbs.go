// +build hbs

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

var App = hydra.NewApp(
	hydra.WithPlatName("hbs_sso", "hbs单点登录系统"),
	hydra.WithSystemName("mgrserver", "hbs单点登录管理系统"),
	hydra.WithUsage("hbs单点登录管理系统"),
	hydra.WithServerTypes(http.Web),
	hydra.WithClusterName("prod"))

func main() {
	App.Start()
}
