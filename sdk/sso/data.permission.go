package sso

import (
	"fmt"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
)

//DataPermissionLogic 数据权限
type DataPermissionLogic struct {
	conf *Config
}

// NewDataPermissionLogic xx
func newDataPermissionLogic(conf *Config) *DataPermissionLogic {
	return &DataPermissionLogic{
		conf: conf,
	}
}

//Sync 同步
func (s *DataPermissionLogic) Sync(req SyncReq) (err error) {
	values := net.NewValues()
	values.Set("name", req.Name)
	values.Set("type", req.Type)
	values.Set("value", req.Value)
	values.Set("remark", req.Remark)
	values.Set("ident", s.conf.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + s.conf.secret
	values.Set("sign", md5.Encrypt(raw))

	result := make(map[string]string)
	a, err := remoteRequest(s.conf.host, syncDataPermission, values.Join("=", "&"), &result)
	if err != nil {
		return err
	}
	fmt.Println(a)
	return nil
}
