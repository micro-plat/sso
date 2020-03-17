module github.com/micro-plat/sso

go 1.12

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

replace github.com/micro-plat/lib4go => ../../../github.com/micro-plat/lib4go

replace gitlab.100bm.cn/micro-plat/dds/dds => ../../../gitlab.100bm.cn/micro-plat/dds/dds

require (
	github.com/Owen-Zhang/base64Captcha v0.0.0-20200225080800-fd8d3d1462c2
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmz331/gpinyin v0.0.0-20150408151959-c4a0503fb352
	github.com/micro-plat/hydra v0.12.2
	github.com/micro-plat/lib4go v0.3.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	gitlab.100bm.cn/micro-plat/dds/dds v0.0.0-00010101000000-000000000000
)
