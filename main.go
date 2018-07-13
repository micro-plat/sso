package main

import "github.com/micro-plat/hydra/hydra"

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("ums-wl"),
		hydra.WithSystemName("sso"),
		hydra.WithServerTypes("api-ws"),
		hydra.WithDebug())
	bind(app)
	app.Start()
}

// func main() {

// 	tk := mp.NewDefaultAccessTokenByURL("wx9e02ddcc88e13fd4", "6acb2b999177524beba3d97d54df2de5", "http://59.151.30.153:9999/wx9e02ddcc88e13fd4/wechat/token/get")
// 	wectx := mp.NewContext(tk)
// 	m := &menu.Menu{
// 		Buttons: []menu.Button{
// 			menu.Button{Type: menu.ButtonTypeView, Name: "用户系统", URL: xurl},
// 		},
// 	}
// 	err := menu.Create(wectx, m)
// 	fmt.Println("err：", err)

// }
