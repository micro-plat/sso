module github.com/micro-plat/sso

go 1.12

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

replace github.com/micro-plat/lib4go => ../../../github.com/micro-plat/lib4go

replace gitlab.100bm.cn/micro-plat/dds/dds => ../../../gitlab.100bm.cn/micro-plat/dds/dds

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmz331/gpinyin v0.0.0-20150408151959-c4a0503fb352
	github.com/micro-plat/hydra v0.12.2
	github.com/micro-plat/lib4go v0.2.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/profile v1.3.0
	github.com/urfave/cli v1.22.1
	github.com/zkfy/log v0.0.0-20180312054228-b2704c3ef896
	gitlab.100bm.cn/micro-plat/dds/dds v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.2
)
