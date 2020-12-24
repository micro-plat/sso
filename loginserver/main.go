package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/sso/loginserver/loginapp"
)

func main() {
	loginapp.App.Start()
}
