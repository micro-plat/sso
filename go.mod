module github.com/micro-plat/sso

go 1.12

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.1
	github.com/gorilla/websocket v1.4.0
	github.com/jordan-wright/email v0.0.0-20190218024454-3ea4d25e7cf8
	github.com/json-iterator/go v1.1.6
	github.com/mattn/go-isatty v0.0.7
	github.com/micro-plat/hydra v0.10.6
	github.com/micro-plat/lib4go v0.1.4
)
