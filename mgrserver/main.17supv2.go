// +build 17supv2

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

var App = hydra.NewApp(
	hydra.WithPlatName("yxtx_17supv2_sso", "一项天下单点登录系统"),
	hydra.WithSystemName("mgrserver", "一项天下单点登录管理系统"),
	hydra.WithUsage("一项天下单点登录管理系统"),
	hydra.WithServerTypes(http.Web),
	hydra.WithClusterName("prod"))

func main() {
	App.Start()
}
