package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/sso/mgrserver/mgrapp"
)

func main() {
	mgrapp.App.Start()
}
