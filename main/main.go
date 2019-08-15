package main

import (
	"fmt"

	"github.com/micro-plat/sso/sso"
)

func main() {
	client, err := sso.New("http://192.168.5.78:6689", "sso", "B128F779D5741E701923346F7FA9F95C")
	if err != nil {
		fmt.Println(err)
	}

	client.GetSystemInfo()
}
