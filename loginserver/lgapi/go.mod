module github.com/micro-plat/sso/loginserver/lgapi

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/micro-plat/hydra v0.13.3
	github.com/micro-plat/lib4go v0.4.0
	github.com/micro-plat/sso v1.3.1
	gitlab.100bm.cn/micro-plat/vcs/vcs v0.0.1
)

replace github.com/micro-plat/sso => ../../../../../github.com/micro-plat/sso
