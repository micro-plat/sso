module github.com/micro-plat/sso

go 1.12

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/bradfitz/gomemcache v0.0.0-20190329173943-551aad21a668
	github.com/dsnet/compress v0.0.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ole/go-ole v1.2.4
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.1
	github.com/gorilla/websocket v1.4.0
	github.com/jordan-wright/email v0.0.0-20190218024454-3ea4d25e7cf8
	github.com/json-iterator/go v1.1.6
	github.com/mattn/go-isatty v0.0.7
	github.com/micro-plat/hydra v0.10.6
	github.com/micro-plat/lib4go v0.1.4
	github.com/micro-plat/wechat v0.0.0-20190124071457-da29af712665
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/nwaples/rardecode v1.0.0
	github.com/pierrec/lz4 v2.2.4+incompatible
	github.com/pierrec/xxHash v0.1.5
	github.com/pkg/profile v1.3.0
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	github.com/shirou/gopsutil v2.19.6+incompatible
	github.com/ugorji/go v1.1.4
	github.com/ulikunitz/xz v0.5.6
	github.com/urfave/cli v1.20.0
	github.com/yosssi/gmq v0.0.1
	github.com/zkfy/archiver v1.1.2
	github.com/zkfy/cron v0.0.0-20170309132418-df38d32658d8
	github.com/zkfy/go-metrics v0.0.0-20161128210544-1f30fe9094a5
	github.com/zkfy/go-oci8 v0.0.0-20180327092318-ad9f59dedff0
	github.com/zkfy/jwt-go v3.0.0+incompatible
	github.com/zkfy/log v0.0.0-20180312054228-b2704c3ef896
	github.com/zkfy/stompngo v0.0.0-20170803022748-9378e70ca481
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223
	golang.org/x/text v0.3.0
	google.golang.org/appengine v1.1.0
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.22.0
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra
