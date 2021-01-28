module github.com/micro-plat/sso

go 1.15

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

replace github.com/micro-plat/lib4go => ../../../github.com/micro-plat/lib4go

replace github.com/lib4dev/vcs => ../../../github.com/lib4dev/vcs

require (
	github.com/Owen-Zhang/base64Captcha v0.0.0-20200225080800-fd8d3d1462c2
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/mock v1.4.3 // indirect
	github.com/jmz331/gpinyin v0.0.0-20150408151959-c4a0503fb352
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lib4dev/vcs v1.0.0
	github.com/micro-plat/hydra v1.0.5
	github.com/micro-plat/lib4go v1.0.10
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	google.golang.org/appengine v1.6.5 // indirect
)
