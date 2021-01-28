package dds

import (
	"github.com/micro-plat/sso/mgrserver/mgrapi/dds/internal/modules/const/conf"
)

// Config 配置数据库,请通过hydra.OnReadyByInsert修改配置参数
func Config(opts ...ConfOption) {
	for _, opt := range opts {
		opt()
	}
}

//GetDBName 获取已配置的db节点名
func GetDBName() string {
	return conf.DbName
}

//ConfOption 配置选项
type ConfOption func()

// WithDBName 数据库节点名称
func WithDBName(dbName string) ConfOption {
	return func() {
		conf.DbName = dbName
	}
}
