# sso

## 1 运行步骤：
``` shell
go build -tags "oci"

./backend install -r zk://192.168.0.101 -c oss

./backend run -r zk://192.168.0.101 -c oss

```