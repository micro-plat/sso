package sdk

import (
	"fmt"

	"github.com/micro-plat/sso/sdk/model"
	"github.com/micro-plat/sso/sdk/service"
)

func TestMenu() {
	SetConfig(&model.Config{
		ApiHost: "http://192.168.106.226:6689",
		Ident:   "sso",
		Secret:  "B128F779D5741E701923346F7FA9F95C",
	})

	data, err := service.GetUserMenu(10123, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
